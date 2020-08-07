package snakes

import (
	"bytes"
	"encoding/json"
	snek "github.com/JMTyler/battlesnake/_"
	"io/ioutil"
	"net/http"
)

type Proxy struct{}

const baseUrl = "http://battlesnake.jaredtyler.ca/rufio"

func (me *Proxy) GetName() string {
	return "proxy"
}

func (me *Proxy) GetInfo() SnakeInfo {
	res, err := http.Get(baseUrl)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// TODO: should close body?

	var info SnakeInfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		panic(err)
	}
	return info
}

func (me *Proxy) StartGame(context *snek.Context) {
	b, err := json.Marshal(context)
	if err != nil {
		panic(err)
	}
	res, err := http.Post(baseUrl+"/start", "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	res.Body.Close()

}

func (me *Proxy) Move(context *snek.Context) string {
	b, err := json.Marshal(context)
	if err != nil {
		panic(err)
	}
	res, err := http.Post(baseUrl+"/move", "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	res.Body.Close()

	var payload map[string]string
	err = json.Unmarshal(body, &payload)
	if err != nil {
		panic(err)
	}
	return payload["move"]
}

func (me *Proxy) EndGame(context *snek.Context) {
	b, err := json.Marshal(context)
	if err != nil {
		panic(err)
	}
	res, err := http.Post(baseUrl+"/end", "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	res.Body.Close()
}
