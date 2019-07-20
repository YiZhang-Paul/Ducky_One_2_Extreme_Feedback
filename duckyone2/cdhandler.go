package duckyone2

func (c *Controller) executeCd(mode string) {
	switch mode {
	case Deploying:
		c.handleDeploying()
	case Pending:
		c.handlePending()
	case DeployBroken:
		c.handleDeployBroken()
	case Deployed:
		c.handleDeployed()
	}
}

func (c *Controller) handleDeploying() {
	data := map[string]interface{}{
		"BackRgb":   "149,0,149",
		"DropRgb":   "229,0,229",
		"SprintRgb": "0,0,255",
		"Speed":     45,
	}
	c.setState(Deploying, sprintModeAPI, data)
}

func (c *Controller) handlePending() {
	data := map[string]interface{}{
		"BackRgb":     "55,55,55",
		"BlinkRgb":    "138,43,226",
		"SpecialRgb":  "138,43,226",
		"SpecialKeys": make([]string, 0),
		"Interval":    850,
	}
	c.setState(Pending, blinkModeAPI, data)
}

func (c *Controller) handleDeployBroken() {
	data := map[string]interface{}{
		"BackRgb":     "55,55,55",
		"BlinkRgb":    "112,128,144",
		"SpecialRgb":  "112,128,144",
		"SpecialKeys": make([]string, 0),
		"Interval":    550,
	}
	c.setState(DeployBroken, blinkModeAPI, data)
}

func (c *Controller) handleDeployed() {
	data := map[string]interface{}{
		"BackRgb": "135,206,235",
		"WaveRgb": "0,0,255",
	}
	c.setState(Deployed, waveModeAPI, data)
	c.lockStateChange(10000)
}
