package main

import (
"encoding/base64"
"fmt"
"io/ioutil"
"log"
"os"
)

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var coder = base64.NewEncoding(base64Table)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("主要功能：将图片转换成base64字符串\n")
		fmt.Printf("使用方法：./imgToBase64 [filePath]\n")
		fmt.Printf("   其中，filepath为图片的文件路径\n")
		return
	}
	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("图片转换后的base64为：\n")
	fmt.Printf("data:image/png;base64,%s\n", Base64Encode(data))
}

func Base64Encode(encode_byte []byte) []byte {
	return []byte(coder.EncodeToString(encode_byte))
}

func Base64Decode(decode_byte []byte) ([]byte, error) {
	return coder.DecodeString(string(decode_byte))
}