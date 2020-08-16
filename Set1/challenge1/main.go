package main

import (
    "encoding/hex"
    "encoding/base64"
    "fmt"
)

func main() {
    const TheString = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

    TheBytes, err := hex.DecodeString(TheString)
    if err != nil {
        panic(err)
    }
    Result := base64.StdEncoding.EncodeToString(TheBytes)
    fmt.Print(Result)
}

