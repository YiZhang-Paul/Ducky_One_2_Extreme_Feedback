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
	case DeployFailed:
		c.handleDeployFailed()
	}
}

func (c Controller) handleDeploying() {

}

func (c Controller) handlePending() {

}

func (c Controller) handleDeployBroken() {

}

func (c Controller) handleDeployed() {

}

func (c Controller) handleDeployFailed() {

}
