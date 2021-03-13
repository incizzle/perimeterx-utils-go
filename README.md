# Perimeterx Helper Functions

## Install
```bash
$ go get github.com/incizzle/perimeterx-utils-go
```
```go
import pxutils "github.com/incizzle/perimeterx-utils-go"
```

## Usage

Create PC Variable
```go
pxutils.CreatePC(jsonstring, tag) // Create PC Variable pass in jsonstring payload and uuid:tag:ftag
```

Obfuscate String
```go
pxutils.ObfuscateString(text, factor) // Simple Function to Obfuscate string using a factor
```

SimpleTextEncode String
```go
pxutils.SimpleTextEncode(text, factor) // Simple Function to encode text using a factor
```

DecodePayload String
```go
pxutils.DecodePayload(text, factor) // DecodePayload decodes a payload using a factor
```

EncodePayload String
```go
pxutils.EncodePayload(text, factor) // EncodePayload encodes a string using a factor
```

Other Encrpyt Functions
```go
pxutils.H1(n, t)
```

```go
pxutils.H12(t)
```

By: iNcizzle#1337  
Twitter: https://twitter.com/iNcizzle
Decoder Website: https://px.incizzle.dev/  
Enjoy 💜  