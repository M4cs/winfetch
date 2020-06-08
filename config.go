package main

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
)

// Config Struct for Configuration File
type Config struct {
	Format           []string    `json:"format"`           // Format Array for Modules
	ShowASCII        bool        `json:"showASCII"`        // True/False to Display ASCII
	UseDefaultColors bool        `json:"useDefaultColors"` // Use Default Colors
	UseSmallASCII    bool        `json:"useSmallASCII"`    // Use Small ASCII Art
	ASCIIColor       string      `json:"ASCIIColor"`       // Change ASCII Color
	UseCustomASCII   bool        `json:"useCustomASCII"`   // Use Custom ASCII TRUE/FALSE
	CustomASCIIPath  string      `json:"customASCIIPath"`  // Absolute Path to Custom ASCII
	UserColor        string      `json:"userColor"`        // Color for Username
	SepColor         string      `json:"sepColor"`         // Color for Separator
	TitleColor       string      `json:"titleColor"`       // Color for Title
	InfoColor        string      `json:"infoColor"`        // Color for Information
	Titles           TitleValues `json:"titles"`           // Title Struct
	AutoUpdate       bool        `json:"autoupdate"`       // Should Auto Update?
	Version          int64       `json:"configVersion"`    // Configuration Version
}

// TitleValues Struct for Title Strings
type TitleValues struct {
	Memory         string `json:"memory"`
	CPU            string `json:"cpu"`
	CPUCores       string `json:"cpuCores"`
	CPUThreads     string `json:"cpuThreads"`
	GPUs           string `json:"gpus"`
	DiskSize       string `json:"diskSize"`
	WindowsVersion string `json:"windowsVersion"`
	Bios           string `json:"bios"`
	Baseboard      string `json:"baseboard"`
	ProcessCount   string `json:"processCount"`
	Uptime         string `json:"uptime"`
}

func updateConfig(config Config) {
	user, _ := user.Current()
	if config.Version == 0 {
		config.AutoUpdate = true
	}
	config.Version = 1
	file, _ := json.MarshalIndent(config, "", " ")
	_ = ioutil.WriteFile(user.HomeDir+"\\.winfetch.json", file, 0644)
}

func newConfig() Config {
	config := Config{}
	config.Format = []string{"user", "sep", "uptime", "mem", "cpu", "procs", "cpuCores", "cpuThreads", "disk", "wversion", "gpus", "bios", "baseboard"}
	config.ShowASCII = true
	config.UseSmallASCII = false
	config.UseCustomASCII = false
	config.CustomASCIIPath = ""
	config.UseDefaultColors = true
	config.ASCIIColor = "Blue"
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
	config.AutoUpdate = true
	config.Version = 1
	return config
}
