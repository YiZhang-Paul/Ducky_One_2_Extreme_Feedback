package duckyone2

func (c Controller) executeCi(mode string) {
	switch mode {
	case "passing":
		c.handlePassing()
	case "broken":
		c.handleBroken()
	case "building":
		c.handleBuilding()
	case "built":
		c.handleBuilt()
	case "build-failed":
		c.handleBuildFailed()
	}
}

func (c Controller) handlePassing() {

}

func (c Controller) handleBroken() {

}

func (c Controller) handleBuilding() {

}

func (c Controller) handleBuilt() {

}

func (c Controller) handleBuildFailed() {

}
