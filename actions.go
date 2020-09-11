// All of the command actions
package main

// execute command
func executeCommand(cmd Command) {
	switch cmd.Action {
	case cmdActionNewWebProject:
		genWebProjectFolders()
	}
}

func genWebProjectFolders() {
	folders := []string{
		folderApplication,
		folderConfig,
		folderController,
		folderEntity,
		folderMiddleware,
		folderRepository,
		folderService,
	}

	for _, folder := range folders {
		genFolder(folder)
	}
}
