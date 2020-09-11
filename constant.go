// Constants definitions
package main

const (
	Version = "0.0.1"
)

const (
	folderApplication = "app"
	folderConfig      = "config"
	folderController  = "controller"
	folderEntity      = "entity"
	folderMiddleware  = "middleware"
	folderRepository  = "repository"
	folderService     = "service"
)

const (
	cmdActionNewWebProject = "new"
)

var (
	projectFolders = []string{
		folderApplication,
		folderConfig,
		folderController,
		folderEntity,
		folderMiddleware,
		folderRepository,
		folderService,
	}
)
