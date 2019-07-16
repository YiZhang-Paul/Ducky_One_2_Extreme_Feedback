package duckyone2

import (
	"fmt"
	"os"

	// auto load .env variables
	_ "github.com/joho/godotenv/autoload"
)

var (
	engineHost   = os.Getenv("ENGINE_HOST")
	colorModeAPI = fmt.Sprintf("%s/colorMode", engineHost)
)

// NotificationMeta contains required information from ci/cd services
type NotificationMeta struct {
	Event string      `json:"event"`
	Mode  string      `json:"mode"`
	Data  interface{} `json:"data,omitempty"`
}

// Controller is the main control for device illumination
type Controller struct{}

// Execute receives data from ci/cd services and controls devices to provide feedbacks on received data
func (c Controller) Execute(meta NotificationMeta) {
	if meta.Event == "ci" {
		c.executeCi(meta.Mode)
	} else if meta.Event == "cd" {
		c.executeCd(meta.Mode)
	}
}