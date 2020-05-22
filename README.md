<p align="center">
    <b>
    <img src="https://raw.githubusercontent.com/M4cs/winfetch/master/gitimages/logo.png"/>
    <p align="center">winfetch</h1>
    <p align="center">A command-line system info tool written in Go for Windows</p>
    <p align="center">Inspired by [neofetch](https://github.com/dylanaraps/neofetch)</p></b>
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

## Dependencies

[ghw](https://github.com/jaypipes/ghw) - Hardware/System Information

[fatih/color](https://github.com/fatih/color) - Color Printing To Terminal

[neofetch](https://github.com/dylanaraps/neofetch) - Inspiration and Ascii Art


