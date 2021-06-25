package main

type Config struct {
	ProjectDir string // project directory
}

// Command command parameters
type Command struct {
	Action     string // command action
	ProjectDir string // project root directory (command working directory)
}
