package sms

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	version = "v1"
	host    = "api.kavenegar.com"
)

type SMS struct {
	apikey  string
	baseURL string
}

func New(apikey string) *SMS {
	s := new(SMS)
	s.apikey = apikey
	s.baseURL = "http://" + host + "/" + version + "/" + apikey + "/" +
		"%s/%s.json"
	return s
}

func (s SMS) request(action, method string, params url.Values) (int, string, error) {
	url := fmt.Sprintf(s.baseURL, action, method)
	resp, err := http.PostForm(url, params)
	if err != nil {
		return 400, "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(body), nil
}

func (s SMS) Send(receptor, message string) error {
	params := url.Values{"receptor": {receptor}, "message": {message}}
	_, _, err := s.request("sms", "send", params)
	return err
}
