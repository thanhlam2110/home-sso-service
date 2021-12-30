package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"github.com/thanhlam/sso-service/config"
)

type RTokens struct {
	Accesstoken string `json:"access_token"`
	Tokentype   string `json:"token_type"`
	Expiresin   string `json:"expires_in"`
	Scope       string `json:"scope"`
}

func BasicAuth(token string) (s string, err error) {
	//config.ReadConfig()
	//link := viper.GetString(`sso.url`)
	link := "https://sso-hp.com.vn:8444/cas/oauth2.0"
	var username string = "exampleOauthClient"
	var passwd string = "exampleOauthClientSecret"
	client := &http.Client{}
	url := link + "/profile?access_token=" + token
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return "", err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s = string(bodyText)
	return s, nil
}

type Tokens struct {
	Accesstoken  string `json:"access_token"`
	Refreshtoken string `json:"refresh_token"`
	Tokentype    string `json:"token_type"`
	Expiresin    string `json:"expires_in"`
	Scope        string `json:"scope"`
}

func RequestTokenv2(client_id string, client_secret string, username string, password string) (string, string) {
	link := "https://sso-hp.com.vn:8444/cas/oauth2.0"
	url := link + "/token?grant_type=password&client_id=" + client_id + "&client_secret=" + client_secret + "&username=" + username + "&password=" + password
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		//Sys log
		fmt.Println(err)
		return "", ""
		//Sys log
	}
	body, _ := ioutil.ReadAll(res.Body)
	stringBody := string(body)
	if stringBody != "" {
		var tokens Tokens
		json.Unmarshal([]byte(stringBody), &tokens)
		return tokens.Accesstoken, tokens.Refreshtoken
	} else {
		return "", ""
	}
}
