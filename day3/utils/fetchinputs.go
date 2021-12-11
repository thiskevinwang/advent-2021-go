package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var Url = "https://adventofcode.com/2021/day/3/input"
var SessionTokenKey = "AOC_SESSION_TOKEN"

// fetch inputs from advent of code
func FetchInputs() []string {
	// session=...
	sessionToken, ok := os.LookupEnv(SessionTokenKey)
	if !ok {
		fmt.Println(SessionTokenKey, `is missing.
	You can grab one from the dev tools from the
	"session=..." cookie.
	- Then run export AOC_SESSION_TOKEN=<token>
	- Or :let $AOC_SESSION_TOKEN=<token>, from vim`)
	}

	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		panic(err)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: sessionToken})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.Trim(string(bodyBytes), "\n"), "\n")

}
