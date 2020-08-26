package main

import (
	"fmt"
	"encoding/hex"
)

func singleXor(c byte, s []byte)([]byte){
	Result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		Result[i] = c^s[i]
	}

	return Result
}

func  repeatXor(xkey []byte, s []byte)([]byte){
	Result := make([]byte, len(s))
	key_counter := 0
	for i := 0; i < len(s); i++ {
		Result[i] = xkey[key_counter] ^ s[i]
		key_counter = (key_counter + 1) % len(xkey)
	}

	return Result

}



func main() {
	test_string := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	xkey := []byte("ICE")
	final_string := hex.EncodeToString(repeatXor(xkey,test_string))
	fmt.Println(final_string)


}
