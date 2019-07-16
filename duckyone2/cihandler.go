package duckyone2

import (
	"time"

	"github.com/yi-zhang/ducky-one-2-extreme-feedback/utils"
)

func (c Controller) executeCi(mode string) {
	switch mode {
	case Passing:
		c.handlePassing()
	case Broken:
		c.handleBroken()
	case Building:
		c.handleBuilding()
	case Built:
		c.handleBuilt()
	}
}

func (c Controller) handlePassing() {
	if currentState == Passing || !canChangeState {
		return
	}
	data := map[string]interface{}{
		"BackRgb":   "1,28,73",
		"ActiveRgb": "255,255,255",
		"Steps":     60,
	}
	utils.PostJSON(reactiveModeAPI, data)
	currentState = Passing
}

func (c Controller) handleBroken() {
	if currentState == Broken || !canChangeState {
		return
	}
	data := map[string]interface{}{
		"BackRgb":  "255,0,0",
		"Interval": 550,
	}
	utils.PostJSON(blinkModeAPI, data)
	currentState = Broken
}

func (c Controller) handleBuilding() {
	if currentState == Building || !canChangeState {
		return
	}
	data := map[string]interface{}{
		"BackRgb":    "255,0,255",
		"InnerRgb":   "0,255,225",
		"OuterRgb":   "255,99,71",
		"InnerSpeed": 25,
		"OuterSpeed": 55,
	}
	utils.PostJSON(progressModeAPI, data)
	currentState = Building
}

func (c Controller) handleBuilt() {
	if currentState == Built || !canChangeState {
		return
	}
	data := map[string]interface{}{
		"BackRgb":  "0,255,0",
		"Interval": 350,
	}
	utils.PostJSON(blinkModeAPI, data)
	currentState = Built
	canChangeState = false
	select {
	case <-time.After(time.Millisecond * 3500):
		canChangeState = true
	}
}
