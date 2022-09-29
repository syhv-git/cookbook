package cmd

func Help(mode string) {
	switch mode {
	case "basic":
	default:
		Fatal("## Unknown keyword: %s\n\tExpected: cookbook file:enumerate help", mode)
	}
}
