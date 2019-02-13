// zabbix alert
// GOARCH=amd64 GOOS=linux go build -o zabbixAlert main.go

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const dingUrl = "https://oapi.dingtalk.com/*********"


type dingMsg struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

func (msg dingMsg) sendMsg() {
	// marshal msg to json
	jsondata, _ := json.Marshal(msg)
	log.Println("msgData:", string(jsondata))

	// new reader
	msgstr := bytes.NewReader(jsondata)

	req, err := http.NewRequest("POST", dingUrl, msgstr)
	if err != nil {
		log.Panic(err)
	}
	// add content-type
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// start a request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	log.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
}

func main() {
	msg := dingMsg{
	}
	// init msg
	msg.Msgtype = "text"
	msg.Text.Content = os.Args[1]
	msg.At.AtMobiles = nil
	msg.At.IsAtAll = false

	msg.sendMsg()
}
