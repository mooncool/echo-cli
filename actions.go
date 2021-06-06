// All of the command actions
package main

// execute command
func executeCommand(cmd Command) {
	switch cmd.Action {
	case cmdActionInitWebProject:
		genWebProjectFolders()
	}
}

func genWebProjectFolders() {
	for _, folder := range projectFolders {
		genFolder(folder)
	}
}
