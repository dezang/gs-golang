package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.upbit.com/v1/ticker?markets=KRW-BTC")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(data))
}
