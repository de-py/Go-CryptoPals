package main

import (
    "encoding/hex"
    "fmt"
)

func doXor(a, b []byte) ([]byte,error) {
    n := len(a)
    
    if len(a) != len(b) {
        return nil, fmt.Errorf("Bytes Slices not the same length")
    }
    Result := make([]byte, n)
    for i := 0; i < n; i++ {
        Result[i] = a[i] ^ b[i]
    }
    return Result, nil
}

func main() {
    const  TheString2 = "1c0111001f010100061a024b53535009181c"
    const Key2 = "686974207468652062756c6c277320657965"
    StrBytes , err := hex.DecodeString(TheString2)
    if err != nil{
        panic(err)
    }
    KeyBytes, err := hex.DecodeString(Key2)
    if err != nil{
        panic(err)
    }
    Result, err := doXor(KeyBytes,StrBytes)
    if err != nil{
        panic(err)
    }
    fmt.Println(hex.EncodeToString(Result))

}
