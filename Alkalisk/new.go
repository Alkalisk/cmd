package main

var newCmd = &Command{
	Execute:     new,
	Name:        "new",
	Usage:       "alkalisk new projectname",
	Description: "Creates a new project in your working directory.",
}

func new(args []string) {
	_ := mustCopyDir()
}
