package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const Url = "https://adventofcode.com/2021/day/1/input"
const SessionTokenKey = "AOC_SESSION_TOKEN"

func main() {
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

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
}
