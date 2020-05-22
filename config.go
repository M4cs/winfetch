package main

type Config struct {
	ShowPhysMem bool `json:"showPhysMem"`
	ShowUsableMem bool `json:"showUsableMem"` 
	ShowTotalCPUCores bool `json:"showTotalCPUCores"`
	ShowTotalCPUThreads bool `json:"showTotalCPUThreads"`
	ShowGPUS bool `json:"showGPUS"`
	ShowTotalDiskSize bool `json:"showTotalDiskSize"`
	ShowWindowsVersion bool `json:"showWindowsVersion"`
	ShowBios bool `json:"showBios"`
	ShowBaseboard bool `json:"showBaseboard"`
	ShowAscii bool `json:"showAscii"`
}

func newConfig() Config {
	config := Config{}
	config.ShowPhysMem = true
	config.ShowUsableMem = true
	config.ShowTotalCPUCores = true
	config.ShowTotalCPUThreads = true
	config.ShowGPUS = true
	config.ShowTotalDiskSize = true
	config.ShowWindowsVersion = true
	config.ShowBios = true
	config.ShowBaseboard = true
	config.ShowAscii = true
	return config
}