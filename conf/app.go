//go:build !prod

package conf

const (
	AppId            = "fdsim_main"
	WindowTitle      = "FDSim"
	WindowWidth      = 900
	WindowHeight     = 600
	WindowFixed      = true
	EnableConsoleLog = true
	Version          = "dev"
	DbFiles          = "fdsim.db"
)
