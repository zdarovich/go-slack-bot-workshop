package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PictureResponse struct {
	Results []Result `json:"results"`
}

type Result struct {
	Picture Picture `json:"picture"`
}

type Picture struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}

func getRandomPictureUrl() string {
	fmt.Println("get random picture")

	cli := &http.Client{}
	req, err := http.NewRequest("GET", "https://randomuser.me/api/", nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var respObj PictureResponse
	err = json.Unmarshal(body, &respObj)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return respObj.Results[0].Picture.Medium
}
