package duckyone2

func (c Controller) executeCd(mode string, data map[string]interface{}) {
	switch mode {
	case "deploying":
		c.handleDeploying(data)
	case "pending":
		c.handlePending(data)
	case "deploy-broken":
		c.handleDeployBroken(data)
	case "deployed":
		c.handleDeployed(data)
	case "deploy-failed":
		c.handleDeployFailed(data)
	}
}

func (c Controller) handleDeploying(data map[string]interface{}) {

}

func (c Controller) handlePending(data map[string]interface{}) {

}

func (c Controller) handleDeployBroken(data map[string]interface{}) {

}

func (c Controller) handleDeployed(data map[string]interface{}) {

}

func (c Controller) handleDeployFailed(data map[string]interface{}) {

}
