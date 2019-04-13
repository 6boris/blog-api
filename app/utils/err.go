package utils

import (
	"log"
)

func SysErr(err error) {
	//log.Println("APP_RUN_ERROR:", err.Error())
	log.Fatalln(err.Error())
}
