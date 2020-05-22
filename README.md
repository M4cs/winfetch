<p align="center">
    <b>
    <img src="https://raw.githubusercontent.com/M4cs/winfetch/master/gitimages/logo.png"/>
    <p align="center">winfetch</h1>
    <p align="center">A command-line system info tool written in Go for Windows</p>
    <p align="center">Inspired by <a href="https://github.com/dylanaraps/neofetch">neofetch</a></p>
    <p align="center"><img src="https://raw.githubusercontent.com/M4cs/winfetch/master/gitimages/preview.gif"/></p></b>
</p>


### Overview
Winfetch is still extremely early in development. The configuration is minimal and there is little to no customization for titles until I add them! This is my first Go program and I have only known it as of last night (when creating this repository).

#### Planned Features:

- Customize Titles
- Customize Colors
- Scaling To Terminal Size
- Display Images (Does Windows Terminal even support this??)

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
 "asciiValue": "Blue",
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

Currently you can only decide what to display in the Winfetch prompt.


## Dependencies

[ghw](https://github.com/jaypipes/ghw) - Hardware/System Information

[fatih/color](https://github.com/fatih/color) - Color Printing To Terminal

[neofetch](https://github.com/dylanaraps/neofetch) - Inspiration and Ascii Art


