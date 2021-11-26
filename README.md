# ptd-tool

A tool for decoding/encoding PTD related data files.
Currently it can decode

* MD (stands for Master Data?) files
* server responses captured with MITM http proxy tools like Burp suite, mitmproxy, etc.
* Some of save data files in `SaveData` directory. Files with name starting with numbers cannot currently be decoded. pa.ds is a plain BSON file that keeps truck of downloaded asset bundles1.

Encoding save data back to encrypted BSON is also supported.

Response/MD decryption code is based on esterTion's work here. 

https://estertion.win/2019/03/project-tokyo-dolls-%E6%8B%86%E8%A7%A3%E7%9B%B8%E5%85%B3/

I rewrote the tool in go and added many other features described above.
For response decoding, I added a feature to decode responses after logging in.
I also fixed issues with string fields of MD, which causes decoded strings and subsequent entries to be corrupted.

Currently requests cannot be decoded.

## Building

go 1.16 or higher is required.

```
go build
```

Or you could use `go run .` to compile and run with single command.

## Usage

Most of subcommands writes result to standard output. Format with jq or your favorite text editor to get formatted output.

### Decoding responses

You don't need to give shared key to decode GetNativeToken/Login response. Just throw in json captured with the tool.

```
./ptd-tool decode-response Login-input.json > Login.json
```

For other responses, you need to give `--shared-security-key`. The game uses a key unique to each user, to encrypt all requests and responses between your device and server. You can find it in one of responses from sqex-bridge server.
The key file must be exactly 32 bytes in length.

```
./ptd-tool decode-response GetServerTime-input.json --shared-security-key shared-security-key.txt > Login.json
```

### Decoding/Encoding local save data

Local save data always uses embedded key for encoding/decoding so you don't need to prepare one.

Encoding:

```
./ptd-tool decode-save-data ./OrigSaveData/oiou_0.ds > ./DecodedSaveData/oiou_0.json
```

Decoding:

```
./ptd-tool encode-save-data ./DecodedSaveData/oiou_0.json > ./SaveData/oiou_0.ds
```

### Decoding MD

You need to have copy of md directory created by the app on your computer.
Replace `/path/to/md-dir` with actual md directory path.

```
./ptd-tool decode-md /path/to/md-dir/MD_Avatar.json --config ./assets/md_loader_configs/MD_Avatar_config.json > MD_Avatar.json
```

`./assets/md_loader_configs/` contains configs needed for decoding some MD files.

### Generating Login response

First, decode MD files with the following script first to generate required JSON files.
Replace `/path/to/md-dir` with actual md directory path.

```
MD_DIR=/path/to/md-dir/md ./generate-md-assets.sh
```

Then generate response with this command. Don't send json formatted with jq to the device if you want to use it with ptd-hook. The app does not accept formatted JSON.

```
./ptd-tool generate-login-response > Login.json
```
