package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(path string) []byte {
	file_byte, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return file_byte
}

func ReadFile2(fileName string) {
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		contentByte, _ := ioutil.ReadAll(f)
		fmt.Println(string(contentByte))
	}

}
