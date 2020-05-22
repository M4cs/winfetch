package main

type Config struct {
	Format []string `json:"format"`
	ShowAscii bool `json:"showAscii"`
	UseDefaultColors bool `json:"useDefaultColors"`
	UseSmallAscii bool `json:"useSmallAscii"`
	AsciiColor string `json:"asciiColor"`
	UseCustomAscii bool `json:"useCustomAscii"`
	CustomAsciiPath string `json:"customAsciiPath"`
	UserColor string `json:"userColor"`
	SepColor string `json:"sepColor"`
	TitleColor string `json:"titleColor"`
	InfoColor string `json:"infoColor"`
	Titles TitleValues `json:"titles"`
}

type TitleValues struct {
	Memory string `json:"memory"`
	CPU string `json:"cpu"`
	CPUCores string `json:"cpuCores"`
	CPUThreads string `json:"cpuThreads"`
	GPUs string `json:"gpus"`
	DiskSize string `json:"diskSize"`
	WindowsVersion string `json:"windowsVersion"`		
	Bios string `json:"bios"`
	Baseboard string `json:"baseboard"`
	ProcessCount string `json:"processCount"`
	Uptime string `json:"uptime"`
}

func newConfig() Config {
	config := Config{}
	config.Format = []string{"user", "sep", "uptime", "mem", "cpu", "procs", "cpuCores", "cpuThreads", "disk", "wversion", "gpus", "bios", "baseboard"}
	config.ShowAscii = true
	config.UseSmallAscii = false
	config.UseCustomAscii = false
	config.CustomAsciiPath = ""
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
	config.Titles.Uptime = "Uptime"
	config.Titles.ProcessCount = "Processes"
	config.Titles.CPU = "CPU #"
	return config
}