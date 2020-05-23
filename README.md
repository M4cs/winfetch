<p align="center">
    <b>
    <img src="https://raw.githubusercontent.com/M4cs/winfetch/master/gitimages/logo.png"/><br>
    <img src="https://img.shields.io/github/stars/M4cs/winfetch"> <img src="https://img.shields.io/github/issues/M4cs/winfetch"> <img src="https://travis-ci.com/M4cs/winfetch.svg?branch=master"> <a href="https://codeclimate.com/github/M4cs/winfetch/maintainability"><img src="https://api.codeclimate.com/v1/badges/5d5dcd97b51f9e01189c/maintainability" /></a>
    <p align="center">A command-line system info tool written in Go for Windows</p>
    <p align="center">Inspired by <a href="https://github.com/dylanaraps/neofetch">neofetch</a></p>
    </b>
</p>




### Overview
Winfetch is an alternative program for neofetch/screenfetch made for Windows! It allows you to display system information through your command line without needing to have any hacky bash fixes to run neofetch. It's also faster!

For More information and detailed instructions on configuration and installation read the [Wiki here](https://github.com/M4cs/winfetch/wiki)

<p align="center">
    <img src="https://raw.githubusercontent.com/M4cs/winfetch/master/gitimages/preview.png" align="center" height="270px">
    <img src="https://raw.githubusercontent.com/M4cs/winfetch/master/gitimages/preview1.png" align="center" height="270px">
</p>

## Installation

### Downloading Binary/Installer

You can find the Binary and Installer in the [Releases Section](https://github.com/M4cs/winfetch/releases). If you download the installer it will install `winfetch.exe` to whichever path you specify, the default is `C:\
Program Files\winfetch\winfetch.exe`. 

If you download the binary alone in `.zip` format you can unzip and extract it to somewhere inside your `$PATH`. 

**Make sure wherever you install `winfetch` to is inside your $PATH! For more info refer [here](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/)**

<hr>

### Using `go get`

To install from go simply run:

```
go get github.com/M4cs/winfetch
```

<hr>

### Building From Source

Clone the repository:

```
git clone https://github.com/M4cs/winfetch.git
```

Run the following commands:

```
cd winfetch
go build
./winfetch.exe # This will be built inside the winfetch/ directory
```

## Dependencies

[ghw](https://github.com/jaypipes/ghw) - Hardware/System Information

[xterm-256](https://github.com/gilliek/go-xterm256) - Color Printing To Terminal

[neofetch](https://github.com/dylanaraps/neofetch) - Inspiration and Ascii Art


## Contribution

Any contributions are welcome! I'm sure the code is ugly as hell since this is really my first Go program but anybody who would like to help out is greatly appreciated! Feel free to Open a PR and Issue with any feedback/suggestions for improvements or features in Winfetch.


