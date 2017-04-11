package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	conf "github.com/ivancduran/edgecast/conf"
)

type resGkey struct {
	Id string `json="Id"`
}

func GlobalKey() string {
	url := conf.Url + conf.AccountNumber + "/fmsliveauth/globalkey"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", conf.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Host", "api.edgecast.com")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println("Response Body:", string(body))

	x := new(resGkey)
	err = json.Unmarshal(body, &x)
	if err != nil {
		panic(err)
	}

	return x.Id
}

func StreamKeys() {
	url := conf.Url + conf.AccountNumber + "/fmsliveauth/streamkeys"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", conf.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Host", "api.edgecast.com")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response Body:", string(body))
	//
	//	x := new(resGkey)
	//	err = json.Unmarshal(body, &x)
	//	if err != nil {
	// panic(err)
	//	}
	//
	//	return x.Id
}
