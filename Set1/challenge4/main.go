package main

import (
    "encoding/hex"
    "fmt"
    "strings"
    "os"
    "bufio"
)


type high struct {
    score int
    letter string
    result string

}

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

func highScore(s string) *high {

    const allChars = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~`'"

    sBytes, err := hex.DecodeString(s)
    if err != nil{
        panic(err)
    }


    h_temp  := high{score: 0, letter: "", result: ""}

    for i := 0; i < len(allChars); i++ {
        Result := singleXor(allChars[i], sBytes)
        Score := rateResult(Result)
        if Score >= h_temp.score {
            h_temp.score = Score
            h_temp.letter = string(allChars[i])
            h_temp.result = string(Result)

        }


    }

    return &h_temp
}


func readLines(path string)([]string, error){
    file, err := os.Open(path)
    if err != nil{
        panic(err)
    }

    // Defer executes after surrounding surrounding function returns
    defer file.Close()
    var lines []string
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    
    return lines,scanner.Err()
}

func main(){

    h := high{score: 0, letter: "", result: ""}
    fileName := "4.txt"

    data, err := readLines(fileName)
    if err != nil{
        panic(err)
    }
    
    for i := 0; i < len(data); i++ {
    
        best := highScore(data[i])

        if best.score > h.score {
            h.score = best.score
            h.letter = best.letter
            h.result = best.result
        }
    }
    
    fmt.Println("Results...")
    fmt.Println("High Letter:", h.letter)
    fmt.Println("High String:", h.result)

    
}
