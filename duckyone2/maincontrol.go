package duckyone2

import (
	"fmt"
	"os"
	"time"

	"github.com/yi-zhang/ducky-one-2-extreme-feedback/utils"

	// auto load .env variables
	_ "github.com/joho/godotenv/autoload"
)

var (
	engineHost      = os.Getenv("ENGINE_HOST")
	colorModeAPI    = fmt.Sprintf("%s/colorMode", engineHost)
	reactiveModeAPI = fmt.Sprintf("%s/reactive", colorModeAPI)
	shiftModeAPI    = fmt.Sprintf("%s/shift", colorModeAPI)
	sprintModeAPI   = fmt.Sprintf("%s/sprint", colorModeAPI)
	waveModeAPI     = fmt.Sprintf("%s/wave", colorModeAPI)
	progressModeAPI = fmt.Sprintf("%s/progress", colorModeAPI)
)

// NotificationMeta contains required information from ci/cd services
type NotificationMeta struct {
	Event string      `json:"event"`
	Mode  string      `json:"mode"`
	Data  interface{} `json:"data,omitempty"`
}

// Controller is the main control for device illumination
type Controller struct {
	state          string
	canChangeState bool
}

// NewController creates a new controller and sets default values
func NewController() *Controller {
	return &Controller{
		canChangeState: true,
	}
}

// Execute receives data from ci/cd services and controls devices to provide feedbacks on received data
func (c *Controller) Execute(meta NotificationMeta) {
	if meta.Event == "ci" {
		c.executeCi(meta.Mode)
	} else if meta.Event == "cd" {
		c.executeCd(meta.Mode)
	}
}

func (c *Controller) lockStateChange(milliseconds int) {
	c.canChangeState = false
	select {
	case <-time.After(time.Duration(milliseconds) * time.Millisecond):
		c.canChangeState = true
	}
}

func (c *Controller) setState(state, api string, data map[string]interface{}) {
	if c.state == state || !c.canChangeState {
		return
	}
	utils.PostJSON(api, data)
	c.state = state
}
