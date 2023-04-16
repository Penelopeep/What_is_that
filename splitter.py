import json
from datetime import datetime
import os

blacklist = []
if os.path.exists("blacklist.txt"):
    with open("blacklist.txt", "r") as f:
        blacklist = str(f.read()).split(",")

print("I don't recommend leaving name of logs like this: log_2023_04_16_12_51_33_1274.json because you have to type entire thing with .json")
filename = input("Enter the file name: ")
json_file = open(filename, "r")
json_data = json.load(json_file)
json_file.close()
for i in json_data:
    if i["protoName"] in blacklist:
        print(f"Skipped: {i['protoName']} at {datetime.fromtimestamp(round(i['time'], 0))}")
        continue
    print(f"Extracted: {i['protoName']} at {datetime.fromtimestamp(round(i['time'], 0))}")
    with open(f'./splitted-packets/{str(datetime.fromtimestamp(round(i["time"], 0))).replace(":",",")} - {i["protoName"]}.json', "w") as f:
        json.dump(i["object"], f, indent=4)