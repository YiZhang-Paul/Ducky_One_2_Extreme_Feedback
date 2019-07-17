package duckyone2

func (c Controller) executeCd(mode string) {
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

func (c Controller) handleDeploying() {
	data := map[string]interface{}{
		"BackRgb":   "185,0,185",
		"SprintRgb": "0,0,255",
		"Speed":     30,
	}
	setState(Deploying, sprintModeAPI, data)
}

func (c Controller) handlePending() {
	data := map[string]interface{}{
		"BackRgbs": "138,43,226|55,55,55",
		"Interval": 800,
	}
	setState(Pending, shiftModeAPI, data)
}

func (c Controller) handleDeployBroken() {
	data := map[string]interface{}{
		"BackRgbs": "112,128,144|55,55,55",
		"Interval": 550,
	}
	setState(DeployBroken, shiftModeAPI, data)
}

func (c Controller) handleDeployed() {
	data := map[string]interface{}{
		"BackRgb": "135,206,235",
		"WaveRgb": "0,0,255",
	}
	setState(Deployed, waveModeAPI, data)
	lockStateChange(5000)
}
