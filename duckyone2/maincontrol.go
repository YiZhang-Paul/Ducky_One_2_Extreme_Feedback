package duckyone2

import (
	"log"
	"os"

	// auto load .env variables
	_ "github.com/joho/godotenv/autoload"
)

var engineHost = os.Getenv("ENGINE_HOST")

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
	data, ok := meta.Data.(map[string]interface{})
	if !ok {
		log.Println("Invalid notification format.")
		return
	}
	if meta.Event == "ci" {
		c.executeCi(meta.Mode)
	} else if meta.Event == "cd" {
		c.executeCd(meta.Mode)
	}
}
