package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type NameResponse struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

func getRandomName() string {
	fmt.Println("get random name")
	cli := &http.Client{}
	req, err := http.NewRequest("GET", "https://uinames.com/api/", nil)
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
	var respObj NameResponse
	err = json.Unmarshal(body, &respObj)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf("%s %s", respObj.Name, respObj.Surname)
}
