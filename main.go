package main

import (
	"log"
	"net/http"
	"os"
	"sms/input"
	"sms/out"
)

func main() {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Printf("Starting new sesion")

	outBC := out.NewBackEnd()
	inBc := input.NewBackEnd(outBC)

	http.HandleFunc("/sendSMS", inBc.HandleSendSmsRequest)
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Panicf("error in ListenAndServe: %s", err)
	}
}
