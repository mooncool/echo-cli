// All of the command actions
package main

import "os"

// execute command
func executeCommand(cmd Command) {
	switch cmd.Action {
	case cmdActionInitWebProject:
		initWebProjectFolders(cmd.ProjectDir)
	case cmdActionNewController:
	}
}

func initWebProjectFolders(projectRoot string) {
	for _, folder := range projectFolders {
		genFolder(projectRoot + string(os.PathSeparator) + folder)
	}
}
