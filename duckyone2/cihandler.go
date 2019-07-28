package duckyone2

import (
	"log"
	"os"

	"github.com/yi-zhang/ducky-one-2-extreme-feedback/utils"

	// load color settings
	_ "github.com/joho/godotenv/autoload"
)

var (
	passingBackRgb      = os.Getenv("PASSING_BACK_RGB")
	passingActiveRgb    = os.Getenv("PASSING_ACTIVE_RGB")
	brokenBackRgb       = os.Getenv("BROKEN_BACK_RGB")
	brokenBlinkRgb      = os.Getenv("BROKEN_BLINK_RGB")
	brokenSpecialRgb    = os.Getenv("BROKEN_SPECIAL_RGB")
	buildingBackRgb     = os.Getenv("BUILDING_BACK_RGB")
	buildingProgressRgb = os.Getenv("BUILDING_PROGRESS_RGB")
	builtBackRgb        = os.Getenv("BUILT_BACK_RGB")
	builtWaveRgb        = os.Getenv("BUILT_WAVE_RGB")
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
		"BackRgb":   passingBackRgb,
		"ActiveRgb": passingActiveRgb,
		"Steps":     15,
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
		"BackRgb":     brokenBackRgb,
		"BlinkRgb":    brokenBlinkRgb,
		"SpecialRgb":  brokenSpecialRgb,
		"SpecialKeys": []string{digitToKey(total)},
		"Interval":    550,
	}
	c.setState(Broken, blinkModeAPI, data)
}

func (c *Controller) handleBuilding() {
	data := map[string]interface{}{
		"BackRgb":     buildingBackRgb,
		"ProgressRgb": buildingProgressRgb,
		"Speed":       70,
	}
	c.setState(Building, progressModeAPI, data)
}

func (c *Controller) handleBuilt() {
	data := map[string]interface{}{
		"BackRgb": builtBackRgb,
		"WaveRgb": builtWaveRgb,
	}
	c.setState(Built, waveModeAPI, data)
	c.lockStateChange(10000)
}
