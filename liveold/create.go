package live

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/ivancduran/edgecast/conf"
	"github.com/ivancduran/edgecast/settings"
	"github.com/ivancduran/edgecast/utils"
)

type Stream interface {
	Create() int
	GetStream(s int) *Response
}

type Hls struct {
	EventName        string
	Expiration       string
	InstanceName     string
	KeyFrameInterval int
}

type Smooth struct {
	EventName        string
	Expiration       string
	InstanceName     string
	KeyFrameInterval int
}

type Response struct {
	Id               int
	EventName        string
	InstanceName     string
	KeyFrameInterval int
	PublishingPoints []Regions
	Expiration       string
	HLSPlaybackUrl   string
	HDSPlaybackUrl   string
	Encrypted        bool
	DvrDuration      string
	SegmentSize      int
}

type Regions struct {
	Region string
	Url    string
}

type resCreate struct {
	Id int `json="Id"`
}

func New(s string) Stream {

	var m = Hls{
		utils.Rands(15),
		"2100-01-01",
		"default",
		10,
	}

	if s == "smooth" {
		var m = Smooth{
			utils.Rands(15),
			"2100-01-01",
			"default",
			5,
		}
		return m
	}

	return m

}

func (this Hls) Create() int {
	url := conf.Url + conf.AccountNumber + "/httpstreaming/livehlshds"

	b := new(bytes.Buffer)

	json.NewEncoder(b).Encode(this)

	req, err := http.NewRequest("POST", url, b)
	req.Header.Set("Authorization", conf.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Host", "api.edgecast.com")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", string(b.Len()))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// fmt.Println("Response Status:", res.Status)
	// fmt.Println("Response Headers:", res.Header)
	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println("Response Body:", string(body))

	x := new(resCreate)
	err = json.Unmarshal(body, &x)
	if err != nil {
		panic(err)
	}

	return x.Id
}

func (this Hls) GetStream(s int) *Response {
	ss := strconv.Itoa(s)
	url := conf.Url + conf.AccountNumber + "/httpstreaming/livehlshds/" + ss

	fmt.Println(url)

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

	x := new(Response)
	err = json.Unmarshal(body, &x)
	if err != nil {
		panic(err)
	}

	hls := strings.Replace(x.HLSPlaybackUrl, "&lt;streamName&gt;", x.EventName, 1)
	hds := strings.Replace(x.HDSPlaybackUrl, "&lt;streamName&gt;", x.EventName, 1)

	x.HDSPlaybackUrl = hds
	x.HLSPlaybackUrl = hls

	evkey := x.EventName + "?" + settings.GlobalKey() + "&"

	for k, elem := range x.PublishingPoints {
		elem.Url = strings.Replace(elem.Url, "&lt;streamName&gt;?", evkey, 1)
		x.PublishingPoints[k].Url = elem.Url
	}

	return x

}

func (this Smooth) Create() int {
	return 0
}

func (this Smooth) GetStream(s int) *Response {
	ss := strconv.Itoa(s)
	url := conf.Url + conf.AccountNumber + "/httpstreaming/livehlshds/" + ss

	fmt.Println(url)

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

	x := new(Response)
	err = json.Unmarshal(body, &x)
	if err != nil {
		panic(err)
	}

	hls := strings.Replace(x.HLSPlaybackUrl, "&lt;streamName&gt;", x.EventName, 1)
	hds := strings.Replace(x.HDSPlaybackUrl, "&lt;streamName&gt;", x.EventName, 1)

	x.HDSPlaybackUrl = hds
	x.HLSPlaybackUrl = hls

	evkey := x.EventName + "?" + settings.GlobalKey() + "&"

	for k, elem := range x.PublishingPoints {
		elem.Url = strings.Replace(elem.Url, "&lt;streamName&gt;?", evkey, 1)
		x.PublishingPoints[k].Url = elem.Url
	}

	return x
}
