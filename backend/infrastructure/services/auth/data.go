package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Auth struct {
	Client *http.Client
}

var (
	headers      = map[string][]string{}
	authEndpoint = "auth/v1"
)

func (a *Auth) SignUp(credentials interface{}) (*Res, error) {
	API_URL := os.Getenv("API_URL")
	API_KEY := os.Getenv("API_KEY")

	headers["apikey"] = []string{API_KEY}
	reqBody, _ := json.Marshal(credentials)
	reqURL := fmt.Sprintf("%s/%s/signup", API_URL, authEndpoint)
	url, err := url.Parse(reqURL)
	if err != nil {
		return nil, err
	}

	request := &http.Request{
		Header: headers,
		Method: "POST",
		URL:    url,
		Body:   io.NopCloser(bytes.NewBuffer(reqBody)),
	}

	res, err := a.Client.Do(request)
	if err != nil {
		return nil, err
	}

	r, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var u *Res
	if err := json.Unmarshal(r, &u); err != nil {
		return nil, err
	}

	return u, nil

}

func (a *Auth) SignIn(credentials interface{}) (interface{}, error) {
	API_URL := os.Getenv("API_URL")
	API_KEY := os.Getenv("API_KEY")

	headers["apikey"] = []string{API_KEY}
	reqBody, _ := json.Marshal(credentials)
	reqURL := fmt.Sprintf("%s/%s/token?grant_type=password", API_URL, authEndpoint)
	url, err := url.Parse(reqURL)
	if err != nil {
		return nil, err
	}

	request := &http.Request{
		Header: headers,
		Method: "POST",
		URL:    url,
		Body:   io.NopCloser(bytes.NewBuffer(reqBody)),
	}

	res, err := a.Client.Do(request)
	if err != nil {
		return nil, err
	}

	r, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var u map[string]interface{}
	if err := json.Unmarshal(r, &u); err != nil {
		return nil, err
	}

	return u, nil
}

func (a *Auth) RefreshUser(userToken string, refreshToken string) (*interface{}, error) {
	return nil, nil
}

func (a *Auth) SigOut(userToken string) error {
	API_URL := os.Getenv("API_URL")
	API_KEY := os.Getenv("API_KEY")

	headers["apikey"] = []string{API_KEY}
	headers["Authorization"] = []string{userToken}

	reqURL := fmt.Sprintf("%s/%s/logout", API_URL, authEndpoint)
	url, err := url.Parse(reqURL)
	if err != nil {
		return err
	}

	request := &http.Request{
		Header: headers,
		Method: "POST",
		URL:    url,
		// Body:   io.NopCloser(bytes.NewBuffer(reqBody)),
	}

	res, err := a.Client.Do(request)
	if err != nil {
		return err
	}

	r, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var u map[string]interface{}
	if err := json.Unmarshal(r, &u); err != nil {
		return err
	}

	if u != nil {
		return fmt.Errorf("%v", u)
	}

	return nil
}
