package duckyone2

func (c *Controller) executeCi(mode string) {
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

func (c *Controller) handlePassing() {
	data := map[string]interface{}{
		"BackRgb":   "1,28,73",
		"ActiveRgb": "255,255,255",
		"Steps":     60,
	}
	c.setState(Passing, reactiveModeAPI, data)
}

func (c *Controller) handleBroken() {
	data := map[string]interface{}{
		"BackRgbs": "255,0,0|55,55,55",
		"Interval": 550,
	}
	c.setState(Broken, shiftModeAPI, data)
}

func (c *Controller) handleBuilding() {
	data := map[string]interface{}{
		"BackRgb":    "255,0,255",
		"InnerRgb":   "0,255,225",
		"OuterRgb":   "255,99,71",
		"InnerSpeed": 25,
		"OuterSpeed": 55,
	}
	c.setState(Building, progressModeAPI, data)
}

func (c *Controller) handleBuilt() {
	data := map[string]interface{}{
		"BackRgb": "50,155,100",
		"WaveRgb": "0,255,0",
	}
	c.setState(Built, waveModeAPI, data)
	c.lockStateChange(5000)
}
