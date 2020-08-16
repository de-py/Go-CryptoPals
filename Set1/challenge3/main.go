package main

import (
    "encoding/hex"
    "fmt"
    "strings"
)

func singleXor(c byte, s []byte) ([]byte){
    Result := make([]byte, len(s))
    for i := 0; i < len(s); i++ {
        Result[i] = c^s[i]
    }

    return Result

}

func rateResult(s []byte) (int){
    scoring := map[string]int{
        "e": 13,
        "t": 12,
        "a": 11,
        "o": 10,
        "i": 9,
        "n": 8,
        " ": 7,
        "s": 6,
        "h": 5,
        "r": 4,
        "d": 3,
        "l": 2,
        "u": 1,
    }
    score := 0
    for i := 0; i < len(s); i++ {
        if val, ok := scoring[strings.ToLower(string(s[i]))]; ok{
            score += val
        }
    }

    return score
}

func main(){
    highScore := 0
    highLetter := ""
    highString := ""
    const TheString3 = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

    const allChars = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~`'"


    StrBytes, err := hex.DecodeString(TheString3)
    if err != nil{
        panic(err)
    }

    for i := 0; i < len(allChars); i++ {
    // test := []byte(allChars[i])
    // fmt.Println(test)
        Result := singleXor(allChars[i], StrBytes)
        Score := rateResult(Result)
    // fmt.Println(Score)
        if Score >= highScore {
            highScore = Score
            highLetter = string(allChars[i])
            highString = string(Result)

        //fmt.Println(string(Result))
        }



    }

    fmt.Println("Results:")
    fmt.Println("High Letter:", highLetter)
    fmt.Println("High String:", highString)
}
