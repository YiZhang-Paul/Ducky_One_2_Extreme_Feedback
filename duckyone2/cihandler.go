package duckyone2

import (
	"log"

	"github.com/yi-zhang/ducky-one-2-extreme-feedback/utils"
)

func (c *Controller) executeCi(mode string, data map[string]interface{}) {
	switch mode {
	case Passing:
		c.handlePassing()
	case Broken:
		c.handleBroken(data)
	case Building:
		c.handleBuilding()
	case Built:
		c.handleBuilt()
	}
}

func (c *Controller) handlePassing() {
	data := map[string]interface{}{
		"BackRgb":   "1,28,73",
		"ActiveRgb": "255,255,255",
		"Steps":     60,
	}
	c.setState(Passing, reactiveModeAPI, data)
}

func (c *Controller) handleBroken(data map[string]interface{}) {
	total, ok := utils.FloatFromMap(data, "total")
	if !ok {
		log.Print("Missing total number of broken builds.")
		return
	}
	data = map[string]interface{}{
		"BackRgb":     "55,55,55",
		"BlinkRgb":    "255,99,71",
		"SpecialRgb":  "255,99,71",
		"SpecialKeys": []string{digitToKey(total)},
		"Interval":    550,
	}
	c.setState(Broken, blinkModeAPI, data)
}

func (c *Controller) handleBuilding() {
	data := map[string]interface{}{
		"BackRgb":     "255,0,255",
		"ProgressRgb": "0,255,225",
		"Speed":       25,
	}
	c.setState(Building, progressModeAPI, data)
}

func (c *Controller) handleBuilt() {
	data := map[string]interface{}{
		"BackRgb": "50,155,100",
		"WaveRgb": "0,255,0",
	}
	c.setState(Built, waveModeAPI, data)
	c.lockStateChange(10000)
}
