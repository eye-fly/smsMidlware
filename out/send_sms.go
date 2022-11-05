package out

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (bc *BackEnd) SendSms(number, msgText string) error {
	err := bc.auth()
	if err != nil {
		return fmt.Errorf("auth fail: %w", err)
	}

	values := &url.Values{}
	values.Add("sms.sendMsg.receiver", number)
	values.Add("sms.sendMsg.text", msgText)
	values.Add("sms.sendMsg.clientId", "test")
	values.Add("action", "send")
	values.Add("ok_redirect", "/index.html")
	values.Add("err_redirect", "/error.json")
	values.Add("token", bc.token)
	reqBody := strings.NewReader(values.Encode())

	fmt.Println(values.Encode())

	req, err := http.NewRequest(http.MethodPost,
		"http://"+bc.hostIP+"/Forms/smsSendMsg", reqBody)
	if err != nil {
		return fmt.Errorf("request prase err: %w", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := bc.client.Do(req)
	if err != nil {
		return fmt.Errorf("request do err: %w", err)
	}

	// bdy, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(bdy))

	//check if corrent site
	if res.Request.URL.Path != "/index.html" {
		return fmt.Errorf("got unnexpected redirect while sending sms")
	}

	return nil
}
