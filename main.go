package main

import (
	"fmt"

	"github.com/ivancduran/edgecast/live"
)

func main() {

	stream := live.New("hls")
	id := stream.Create()
	fmt.Println(id)

	entry := stream.GetStream(id)
	fmt.Println(entry.PublishingPoints[0].Url)
	fmt.Println(entry.HLSPlaybackUrl)

}
