package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func main() {

	commandData, err := ioutil.ReadFile("command.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	hash := md5.New()
	hash.Write(commandData)
	fmt.Println(hex.EncodeToString(hash.Sum(nil)))
}
