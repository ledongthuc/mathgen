package messages

import (
	"os"
	"strconv"
	"strings"
	"time"
)

var currentVersion = strconv.FormatInt(time.Now().UnixNano(), 10)

type WebMessage struct {
	Version           string
	ChromeOriginTrial string
	AppName           string
}

func GenerateWebMessage() WebMessage {
	return WebMessage{
		Version:           currentVersion,
		ChromeOriginTrial: strings.TrimSpace(os.Getenv("CHROME_ORIGIN_TRIAL")),
		AppName:           "Math Generator",
	}
}
