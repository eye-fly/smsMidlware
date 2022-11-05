package out

import (
	"net/http"
	"net/url"
	"strings"
)

func (bc *BackEnd) requestInitialSessionId() (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet,
		"http://"+bc.hostIP+"/index.html", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Upgrade-Insecure-Requests", "1")
	return req, nil
}

func (bc *BackEnd) requestPostAuth() (*http.Request, error) {

	values := &url.Values{}
	values.Add("token", bc.token)
	values.Add("ok_redirect", "/index.html")
	values.Add("err_redirect", "/error.json")
	values.Add("session.password", "klop")
	reqBody := strings.NewReader(values.Encode())

	req, err := http.NewRequest(http.MethodPost,
		"http://"+bc.hostIP+"/Forms/config", reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	return req, nil
}
