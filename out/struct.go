package out

import (
	"log"
	"net/http"
	"net/http/cookiejar"
)

type BackEnd struct {
	client http.Client
	hostIP string
	token  string
}

func NewBackEnd() *BackEnd {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Panicf("error in newJar: %s", err)
	}

	return &BackEnd{
		client: http.Client{
			Jar: jar,
		},
	}
}

func (bc *BackEnd) SetHostIp(ip string) {
	bc.hostIP = ip
}
