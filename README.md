# ptd-response-decoder

A tool to decode PTD server responses captured with MITM http proxy tool like Burp suite.
Based on esterTion's work here. 

https://estertion.win/2019/03/project-tokyo-dolls-%E6%8B%86%E8%A7%A3%E7%9B%B8%E5%85%B3/

I rewrote the tool in go and added a feature to decode responses after logging in, which can be used with ptd-hook.
Currently only supports decoding responses. Requests cannot be decoded.

# Build

go 1.16 or higher is required.

```
go build
```

# Usage

You don't need to give shared key to decode GetNativeToken/Login response. Just throw in json captured with the tool.

```
./ptd-response-decoder decode Login.json
```

For other responses, you need to give `--shared-security-key`. The game uses a key unique to each user, to encrypt all requests and responses between your device and server. You can find it in one of responses from sqex-bridge server.
The key file must be exactly 32 bytes in length.

```
./ptd-response-decoder decode GetServerTime.json --shared-security-key shared-security-key.txt
```
