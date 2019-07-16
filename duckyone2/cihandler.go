package duckyone2

func (c Controller) executeCi(mode string, data map[string]interface{}) {
	switch mode {
	case "passing":
		c.handlePassing(data)
	case "broken":
		c.handleBroken(data)
	case "building":
		c.handleBuilding(data)
	case "built":
		c.handleBuilt(data)
	case "build-failed":
		c.handleBuildFailed(data)
	}
}

func (c Controller) handlePassing(data map[string]interface{}) {

}

func (c Controller) handleBroken(data map[string]interface{}) {

}

func (c Controller) handleBuilding(data map[string]interface{}) {

}

func (c Controller) handleBuilt(data map[string]interface{}) {

}

func (c Controller) handleBuildFailed(data map[string]interface{}) {

}
