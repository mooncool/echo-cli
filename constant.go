// Constants definitions
package main

const (
	Version = "0.1.0"
)

const (
	folderApplication = "app"
	folderBootstrap   = "bootstrap"
	folderBuild       = "build"
	folderConfig      = "config"
	folderController  = "controller"
	folderEntity      = "entity"
	folderMiddleware  = "middleware"
	folderRepository  = "repository"
	folderService     = "service"
)

const (
	defaultWorkingDir = "./" // default working directory

	cmdActionInitWebProject = "init" // init web project
	cmdActionNewWebProject  = "new"  // new web project
)

var (
	projectFolders = []string{
		folderApplication,
		folderBootstrap,
		folderBuild,
		folderConfig,
		folderController,
		folderEntity,
		folderMiddleware,
		folderRepository,
		folderService,
	}
)
