package messages

import (
	"strconv"
	"time"
)

var currentVersion = strconv.FormatInt(time.Now().UnixNano(), 10)

type WebMessage struct {
	Version string
}

func GenerateWebMessage() WebMessage {
	return WebMessage{
		Version: currentVersion,
	}
}
