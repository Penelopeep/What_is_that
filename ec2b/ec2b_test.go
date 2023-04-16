package ec2b

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestEC2B(t *testing.T) {
	d, err := base64.StdEncoding.DecodeString("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")
	if err != nil {
		panic(err)
	}
	res, err := Derive(d)
	if err != nil {
		panic(err)
	}
	fmt.Println(res[0], res[1], res[2], res[3])
}