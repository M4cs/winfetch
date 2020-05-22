<p align="center">
    <b>
    <img src="https://raw.githubusercontent.com/M4cs/winfetch/master/gitimages/logo.png"/><br>
    <img src="https://img.shields.io/github/stars/M4cs/winfetch"> <img src="https://img.shields.io/github/issues/M4cs/winfetch"> <img src="https://travis-ci.com/M4cs/winfetch.svg?branch=master"> <a href="https://codeclimate.com/github/M4cs/winfetch/maintainability"><img src="https://api.codeclimate.com/v1/badges/5d5dcd97b51f9e01189c/maintainability" /></a>
    <p align="center">A command-line system info tool written in Go for Windows</p>
    <p align="center">Inspired by <a href="https://github.com/dylanaraps/neofetch">neofetch</a></p>
    </b>
</p>


<img src="https://raw.githubusercontent.com/M4cs/winfetch/master/gitimages/preview.png" align="right" height="250px">

### Overview
Winfetch is still extremely early in development. The configuration is minimal and there is little to no customization for titles until I add them! This is my first Go program and I have only known it as of last night (when creating this repository).

#### Planned Features:

- Scaling To Terminal Size
- Display Images (Does Windows Terminal even support this??)

<img src="https://raw.githubusercontent.com/M4cs/winfetch/master/gitimages/preview1.png" align="right" height="250px">

### Installation

#### Using `go get`

First install [Go](https://golang.org/)

Then use `go get` from the command line to install:

```
go get github.com/M4cs/winfetch
```

Now you can run `winfetch` to get your system information to display!

#### Building From Source

Make sure you have Go installed then run:

```
git clone https://github.com/M4cs/winfetch
cd winfetch
go build
```

You will now have a `winfetch.exe` file in your directory!

## Configuration

There isn't much to the config at the moment but you can find it at `.winfetch.json` in your Home folder.

The default config is as follows:

```json
{
 "showUser": true,
 "showSep": true,
 "showMem": true,
 "showTotalCPUCores": true,
 "showTotalCPUThreads": true,
 "showGPUS": true,
 "showTotalDiskSize": true,
 "showWindowsVersion": true,
 "showBios": true,
 "showBaseboard": true,
 "showAscii": true,
 "useDefaultColors": true,
 "useSmallAscii": false,
 "asciiColor": "Blue",
 "useCustomAscii": false,
 "customAsciiPath": "",
 "userColor": "Red",
 "sepColor": "Red",
 "titleColor": "Green",
 "infoColor": "White",
 "titles": {
  "memory": "Memory",
  "cpuCores": "CPU Cores",
  "cpuThreads": "CPU Threads",
  "gpus": "GPU #",
  "diskSize": "Disk Size",
  "windowsVersion": "Windows Ver.",
  "bios": "BIOS",
  "baseboard": "Baseboard"
 }
}
```


## Dependencies

[ghw](https://github.com/jaypipes/ghw) - Hardware/System Information

[xterm-256](https://github.com/gilliek/go-xterm256) - Color Printing To Terminal

[neofetch](https://github.com/dylanaraps/neofetch) - Inspiration and Ascii Art


