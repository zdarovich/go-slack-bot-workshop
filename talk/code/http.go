package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

type NameResponse struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

func main() {
	cli := &http.Client{}
	req, err := http.NewRequest("GET", "https://uinames.com/api/", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var resp NameResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}