package duckyone2

func (c Controller) executeCd(mode string) {
	switch mode {
	case "deploying":
		c.handleDeploying()
	case "pending":
		c.handlePending()
	case "deploy-broken":
		c.handleDeployBroken()
	case "deployed":
		c.handleDeployed()
	case "deploy-failed":
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
