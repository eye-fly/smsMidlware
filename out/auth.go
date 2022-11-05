package out

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

func findToken(res *http.Response) string {
	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)

	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if strings.Contains(line, "secToken") {
			line = strings.TrimSpace(line)
			line = strings.TrimPrefix(line, "\"secToken\": ")
			line = strings.Trim(line, "\"")
			return line
		}
	}
	return ""
}

func (bc *BackEnd) postAuth() error {
	req, err := bc.requestPostAuth()
	if err != nil {
		return err
	}
	_, err = bc.client.Do(req)
	if err != nil {
		return err
	}

	//check if correct site

	return err
}

func (bc *BackEnd) auth() error {
	req, err := bc.requestInitialSessionId()
	if err != nil {
		return err
	}

	res, err := bc.client.Do(req)
	if err != nil {
		return err
	}

	bc.token = findToken(res)
	if len(bc.token) == 0 {
		return fmt.Errorf("did'n find token")
	}

	bc.postAuth()
	// fmt.Println(client.Jar.Cookies(req.URL))
	return nil
}
