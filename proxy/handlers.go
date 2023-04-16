package proxy

import (
	"encoding/base64"
	"sync"

	"github.com/MoonlightPS/Iridium-gidra/gover/gen"
	"github.com/MoonlightPS/Iridium-gidra/gover/utils"
	"github.com/yezihack/colorlog"
)

type Handler = func(*KCPConn, []byte, bool) ([]byte, error)

var handlersMap = map[int]Handler{}

var b64 = base64.StdEncoding

var prngHistory = struct {
	seeds []uint64
	mu    sync.Mutex
}{}

func saveSeed(seed uint64) {
	prngHistory.mu.Lock()
	defer prngHistory.mu.Unlock()
	prngHistory.seeds = append(prngHistory.seeds, seed)
}

func sniffSeed(mSeed, sentMs uint64) uint64 {
	find := func(ts uint64) uint64 {
		prng := utils.NewCompatPrng(int32(ts))
		for times := 0; times < 1e5; times++ {
			seed := prng.SafeUInt64()
			if seed == mSeed {
				colorlog.Debug("seed found, seed: %d with ts: %d by times: %d", seed, ts, times)
				return seed
			}
		}
		return 0
	}
	for _, ts := range prngHistory.seeds {
		if seed := find(ts); seed != 0 {
			colorlog.Debug("seed found from history")
			return seed
		}
	}
	for times := uint64(0); times < 1e5; times++ {
		if seed := find(sentMs + times); seed != 0 {
			saveSeed(sentMs + times)
			return seed
		}
		if seed := find(sentMs - times); seed != 0 {
			saveSeed(sentMs - times)
			return seed
		}
	}
	colorlog.Debug("seed not found")
	return 0
}

func HandleGetPlayerTokenReq(conn *KCPConn, data []byte, bypass bool) ([]byte, error) {
	msg, err := conn.parser.Parse(data)
	if err != nil {
		return nil, err
	}
	body := msg.Body.(*gen.GetPlayerTokenReq)
	if bypass {
		conn.clientSeed = msg.Header.GetSentMs()
		return conn.parser.Compose(msg)
	}

	seedEncrypted, err := b64.DecodeString(body.GetClientRandKey())
	if err != nil {
		return nil, err
	}

	seedBytes, err := utils.Decrypt(seedEncrypted, utils.SIGN_KEY)
	if err != nil {
		return nil, err
	}
	conn.seed = be.Uint64(seedBytes)
	sniffSeed(conn.seed, msg.Header.GetSentMs())

	keyID := int(body.GetKeyId()) + 1000
	seedEncrypted, err = utils.Encrypt(seedBytes, keyID)
	if err != nil {
		return nil, err
	}

	body.ClientRandKey = b64.EncodeToString(seedEncrypted)
	return conn.parser.Compose(msg)
}

func HandleGetPlayerTokenRsp(conn *KCPConn, data []byte, bypass bool) ([]byte, error) {
	msg, err := conn.parser.Parse(data)
	if err != nil {
		return nil, err
	}
	body := msg.Body.(*gen.GetPlayerTokenRsp)

	seedEncrypted, err := b64.DecodeString(body.GetServerRandKey())
	if err != nil {
		return nil, err
	}

	keyID := int(body.GetKeyId())
	seedBytes, err := utils.Decrypt(seedEncrypted, keyID)
	if err != nil {
		return nil, err
	}

	if bypass {
		conn.seed = be.Uint64(seedBytes)
		return conn.parser.Compose(msg)
	}

	conn.seed = be.Uint64(seedBytes) ^ conn.seed
	colorlog.Info("get server key: %d", conn.seed)

	signature, err := utils.Sign(seedBytes, utils.SIGN_KEY)
	if err != nil {
		return nil, err
	}

	body.Sign = b64.EncodeToString(signature)
	ret, err := conn.parser.Compose(msg)
	conn.key.GenKey(conn.seed)
	return ret, err
}

func HandlePlayerLoginReq(conn *KCPConn, data []byte, bypass bool) ([]byte, error) {
	if bypass {
		return data, nil
	}
	msg, err := conn.parser.Parse(data)
	if err != nil {
		return nil, err
	}
	body := msg.Body.(*gen.PlayerLoginReq)

	if conn.keyID == utils.CN_KEY || conn.keyID == utils.CN1_KEY {
		body.Checksum = "4fa709ab639fc791f8288975f0428f0c912b36de981c9553145ce0c7a35f088725"
	} else {
		body.Checksum = "c1fed3cda007abe60b0d17c6c7a5442aec6d3bc7770693949a7b1ab0483fc16225"
	}

	return conn.parser.Compose(msg)
}

func init() {
	handlersMap[utils.GetPlayerTokenReq] = HandleGetPlayerTokenReq
	handlersMap[utils.GetPlayerTokenRsp] = HandleGetPlayerTokenRsp
	handlersMap[utils.PlayerLoginReq] = HandlePlayerLoginReq
}
