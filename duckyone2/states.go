package duckyone2

// ci states
const (
	Passing  = "passing"
	Broken   = "broken"
	Building = "building"
	Built    = "built"
)

// cd states
const (
	Deploying    = "deploying"
	Pending      = "pending"
	DeployBroken = "deploy-broken"
	Deployed     = "deployed"
)
