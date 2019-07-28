package duckyone2

import (
	"os"

	// load color settings
	_ "github.com/joho/godotenv/autoload"
)

var (
	deployingBackRgb       = os.Getenv("DEPLOYING_BACK_RGB")
	deployingDropRgb       = os.Getenv("DEPLOYING_DROP_RGB")
	deployingSprintRgb     = os.Getenv("DEPLOYING_SPRINT_RGB")
	pendingBackRgb         = os.Getenv("PENDING_BACK_RGB")
	pendingBlinkRgb        = os.Getenv("PENDING_BLINK_RGB")
	pendingSpecialRgb      = os.Getenv("PENDING_SPECIAL_RGB")
	deployBrokenBackRgb    = os.Getenv("DEPLOY_BROKEN_BACK_RGB")
	deployBrokenBlinkRgb   = os.Getenv("DEPLOY_BROKEN_BLINK_RGB")
	deployBrokenSpecialRgb = os.Getenv("DEPLOY_BROKEN_SPECIAL_RGB")
	deployedBackRgb        = os.Getenv("DEPLOYED_BACK_RGB")
	deployedWaveRgb        = os.Getenv("DEPLOYED_WAVE_RGB")
)

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
		"BackRgb":   deployingBackRgb,
		"DropRgb":   deployingDropRgb,
		"SprintRgb": deployingSprintRgb,
		"Speed":     75,
	}
	c.setState(Deploying, sprintModeAPI, data)
}

func (c *Controller) handlePending() {
	data := map[string]interface{}{
		"BackRgb":     pendingBackRgb,
		"BlinkRgb":    pendingBlinkRgb,
		"SpecialRgb":  pendingSpecialRgb,
		"SpecialKeys": make([]string, 0),
		"Interval":    850,
	}
	c.setState(Pending, blinkModeAPI, data)
}

func (c *Controller) handleDeployBroken() {
	data := map[string]interface{}{
		"BackRgb":     deployBrokenBackRgb,
		"BlinkRgb":    deployBrokenBlinkRgb,
		"SpecialRgb":  deployBrokenSpecialRgb,
		"SpecialKeys": make([]string, 0),
		"Interval":    550,
	}
	c.setState(DeployBroken, blinkModeAPI, data)
}

func (c *Controller) handleDeployed() {
	data := map[string]interface{}{
		"BackRgb": deployedBackRgb,
		"WaveRgb": deployedWaveRgb,
	}
	c.setState(Deployed, waveModeAPI, data)
	c.lockStateChange(10000)
}
