package token

import (
	"time"
	http2 "net/http"
	"os"
	"fmt"
	"io/ioutil"
)

const (
	apiUrl = "http://localhost:5000"
)

type OpsGenieToken struct {
	timestamp time.Time
	token     string
}

func IsValidClient(token OpsGenieToken) (OpsGenieToken, error) {
	if checkIsTime(token.timestamp) {
		return tokenGenerator()
	} else {
		return token, nil
	}
}

func (tok OpsGenieToken) Token() string {
	return tok.token
}

func InitialTokenGenerator() OpsGenieToken {
	tok, err := tokenGenerator()
	if err != nil {
		panic(err)
	}
	return OpsGenieToken{time.Now(), tok.token}
}

func tokenGenerator() (OpsGenieToken, error) {
	resp, err := http2.Get(apiUrl)
	contents := ""
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}
	return OpsGenieToken{time.Now(), string(contents)}, nil
}

func checkIsTime(latestCall time.Time) bool {
	diff := time.Now().Sub(latestCall)
	if diff.Seconds() >= 3555 {
		return true
	}
	return false
}
