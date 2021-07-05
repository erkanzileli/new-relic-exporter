package config

import "flag"

var (
	Addr           = flag.String("addr", ":8080", "metric server address, default is :8080")
	AppId          = flag.Int("app-id", 0, "your new-relic application id")
	PersonalApiKey = flag.Bool("personal-key", false, "specify if your api key is personal")
	ApiKey         = flag.String("api-key", "", "your new-relic api key")
	ScrapeInterval = flag.String("interval", "@every 0h0m10s", "scrape interval, default is every 10 seconds")
)
