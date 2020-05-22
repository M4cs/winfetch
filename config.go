package main

type Config struct {
	ShowUser bool `json:"showUser"`
	ShowSep bool `json:"showSep"`
	ShowMem bool `json:"showMem"`
	ShowTotalCPUCores bool `json:"showTotalCPUCores"`
	ShowTotalCPUThreads bool `json:"showTotalCPUThreads"`
	ShowGPUS bool `json:"showGPUS"`
	ShowTotalDiskSize bool `json:"showTotalDiskSize"`
	ShowWindowsVersion bool `json:"showWindowsVersion"`
	ShowBios bool `json:"showBios"`
	ShowBaseboard bool `json:"showBaseboard"`
	ShowAscii bool `json:"showAscii"`
	UseDefaultColors bool `json:"useDefaultColors"`
	AsciiColor string `json:"asciiColor"`
	UserColor string `json:"userColor"`
	SepColor string `json:"sepColor"`
	TitleColor string `json:"titleColor"`
	InfoColor string `json:"infoColor"`
	Titles TitleValues `json:"titles"`
}

type TitleValues struct {
	Memory string `json:"memory"`
	CPUCores string `json:"cpuCores"`
	CPUThreads string `json:"cpuThreads"`
	GPUs string `json:"gpus"`
	DiskSize string `json:"diskSize"`
	WindowsVersion string `json:"windowsVersion"`		
	Bios string `json:"bios"`
	Baseboard string `json:"baseboard"`
}

func newConfig() Config {
	config := Config{}
	config.ShowUser = true
	config.ShowSep = true
	config.ShowMem = true
	config.ShowTotalCPUCores = true
	config.ShowTotalCPUThreads = true
	config.ShowGPUS = true
	config.ShowTotalDiskSize = true
	config.ShowWindowsVersion = true
	config.ShowBios = true
	config.ShowBaseboard = true
	config.ShowAscii = true
	config.UseDefaultColors = true
	config.AsciiColor = "Blue"
	config.UserColor = "Red"
	config.SepColor = "Red"
	config.TitleColor = "Green"
	config.InfoColor = "White"
	config.Titles.Memory = "Memory"
	config.Titles.CPUCores = "CPU Cores"
	config.Titles.CPUThreads = "CPU Threads"
	config.Titles.GPUs = "GPU #"
	config.Titles.DiskSize = "Disk Size"
	config.Titles.WindowsVersion = "Windows Ver."
	config.Titles.Bios = "BIOS"
	config.Titles.Baseboard = "Baseboard"
	return config
}