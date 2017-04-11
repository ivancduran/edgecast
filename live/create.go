package live

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ivancduran/edgecast/conf"
	"github.com/ivancduran/edgecast/settings"
	"github.com/ivancduran/edgecast/utils"
)

type Stream interface {
	Create() *CreateResponse
}

type Hls struct {
	// Encrypted    bool
	InstanceName string
	// DvrDuration  int
	SegmentSize int
}

type Smooth struct {
	EventName        string
	Expiration       string
	InstanceName     string
	KeyFrameInterval int
}

type CreateResponse struct {
	Id           int
	CustomerId   int
	DvrDuration  int
	Encrypted    bool
	InstanceName string
	PlaybackUrl  string
	PublishUrl   string
	HLS          string
	DASH         string
	SegmentSize  int
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
		// false,
		utils.Rands(15),
		// 0,
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

func (this Hls) Create() *CreateResponse {
	url := conf.Url + conf.AccountNumber + "/httpstreaming/dcp/live"

	b := new(bytes.Buffer)

	json.NewEncoder(b).Encode(this)

	fmt.Println(string(b.String()))

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

	body, _ := ioutil.ReadAll(res.Body)

	x := new(CreateResponse)
	err = json.Unmarshal(body, &x)

	if err != nil {
		panic(err)
	}

	playback := strings.Replace(x.PlaybackUrl, "<streamName>", x.InstanceName, 1)

	x.PublishUrl = "rtmp://fso.dca.34C45.xicdn.net" + x.PublishUrl

	publishName := strings.Replace(x.PublishUrl, "<streamName>", x.InstanceName, 1)
	x.PublishUrl = publishName
	publishLive := strings.Replace(x.PublishUrl, "<Live Authentication Key>", settings.GlobalKey(), 1)
	x.PublishUrl = publishLive

	x.PlaybackUrl = playback
	x.HLS = playback + ".m3u8"
	x.DASH = playback + ".mpd"

	return x
}

func (this Smooth) Create() *CreateResponse {
	x := CreateResponse{}
	return &x
}
