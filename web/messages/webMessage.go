package messages

import (
	"os"
	"strconv"
	"time"
)

var currentVersion = strconv.FormatInt(time.Now().UnixNano(), 10)

type WebMessage struct {
	Version           string
	ChromeOriginTrial string
}

func GenerateWebMessage() WebMessage {
	return WebMessage{
		Version:           currentVersion,
		ChromeOriginTrial: os.Getenv("CHROME_ORIGIN_TRIAL"),
	}
}
