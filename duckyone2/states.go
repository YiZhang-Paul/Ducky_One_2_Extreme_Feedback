package duckyone2

// ci states
const (
	Passing     = "passing"
	Broken      = "broken"
	Building    = "building"
	Built       = "built"
	BuildFailed = "build-failed"
)

// cd states
const (
	Deploying    = "deploying"
	Pending      = "pending"
	DeployBroken = "deploy-broken"
	Deployed     = "deployed"
	DeployFailed = "deploy-failed"
)
