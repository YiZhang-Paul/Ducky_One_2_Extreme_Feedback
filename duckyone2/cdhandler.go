package duckyone2

import (
	"time"

	"github.com/yi-zhang/ducky-one-2-extreme-feedback/utils"
)

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
	if currentState == Deploying || !canChangeState {
		return
	}
	data := map[string]interface{}{
		"BackRgb":   "185,0,185",
		"SprintRgb": "0,0,255",
		"Speed":     30,
	}
	utils.PostJSON(sprintModeAPI, data)
	currentState = Deploying
}

func (c Controller) handlePending() {
	if currentState == Pending || !canChangeState {
		return
	}
	data := map[string]interface{}{
		"BackRgbs": "138,43,226|55,55,55",
		"Interval": 800,
	}
	utils.PostJSON(shiftModeAPI, data)
	currentState = Pending
}

func (c Controller) handleDeployBroken() {
	if currentState == DeployBroken || !canChangeState {
		return
	}
	data := map[string]interface{}{
		"BackRgbs": "112,128,144|55,55,55",
		"Interval": 550,
	}
	utils.PostJSON(shiftModeAPI, data)
	currentState = DeployBroken
}

func (c Controller) handleDeployed() {
	if currentState == Deployed || !canChangeState {
		return
	}
	data := map[string]interface{}{
		"BackRgb": "135,206,235",
		"WaveRgb": "0,0,255",
	}
	utils.PostJSON(waveModeAPI, data)
	currentState = Deployed
	canChangeState = false
	select {
	case <-time.After(time.Millisecond * 5000):
		canChangeState = true
	}
}
