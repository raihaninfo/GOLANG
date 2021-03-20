package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	resp := []byte("\x00" + "raihanmahmudi35@gmail.com" + "\x00" + "your password")

	fmt.Println(string(resp), resp)
	sEnc := base64.StdEncoding.EncodeToString([]byte(resp))
	fmt.Println(sEnc)
}
