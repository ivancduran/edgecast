package conf

import "os"

var (
	AccountNumber = os.Getenv("EDGECAST_NO")
	Url           = os.Getenv("EDGECAST_URL")
	Token         = os.Getenv("EDGECAST_TOKEN")
)
