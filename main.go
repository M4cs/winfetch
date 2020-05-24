package main

import (
	"fmt"
	"log"
	"strings"
	"bufio"
	"os/user"
	"os"
	"encoding/json"
	"io/ioutil"
	"github.com/gilliek/go-xterm256/xterm256"
)

func getCustomColor(color string) xterm256.Color {
	var (
		Black       = xterm256.Color{ForegroundColor: 0, BackgroundColor: -1}
		DarkRed     = xterm256.Color{ForegroundColor: 1, BackgroundColor: -1}
		DarkGreen   = xterm256.Color{ForegroundColor: 2, BackgroundColor: -1}
		DarkYellow  = xterm256.Color{ForegroundColor: 3, BackgroundColor: -1}
		DarkBlue    = xterm256.Color{ForegroundColor: 4, BackgroundColor: -1}
		DarkMagenta = xterm256.Color{ForegroundColor: 5, BackgroundColor: -1}
		DarkCyan    = xterm256.Color{ForegroundColor: 6, BackgroundColor: -1}
		LightGray   = xterm256.Color{ForegroundColor: 7, BackgroundColor: -1}
		DarkGray    = xterm256.Color{ForegroundColor: 8, BackgroundColor: -1}
		Red         = xterm256.Color{ForegroundColor: 9, BackgroundColor: -1}
		Green       = xterm256.Color{ForegroundColor: 10, BackgroundColor: -1}
		Yellow      = xterm256.Color{ForegroundColor: 11, BackgroundColor: -1}
		Blue        = xterm256.Color{ForegroundColor: 12, BackgroundColor: -1}
		Magenta     = xterm256.Color{ForegroundColor: 13, BackgroundColor: -1}
		Cyan        = xterm256.Color{ForegroundColor: 14, BackgroundColor: -1}
		White       = xterm256.Color{ForegroundColor: 15, BackgroundColor: -1}
	
		Orange = xterm256.Color{ForegroundColor: 130, BackgroundColor: -1}
	)
	switch color {
	case "Black":
		return Black
	case "DarkRed":
		return DarkRed
	case "DarkGreen":
		return DarkGreen
	case "DarkYellow":
		return DarkYellow
	case "DarkBlue":
		return DarkBlue
	case "DarkMagenta":
		return DarkMagenta
	case "DarkCyan":
		return DarkCyan
	case "LightGray":
		return LightGray
	case "DarkGray":
		return DarkGray
	case "Red":
		return Red
	case "Green":
		return Green
	case "Yellow":
		return Yellow
	case "Blue":
		return Blue
	case "Magenta":
		return Magenta
	case "Cyan":
		return Cyan
	case "White":
		return White
	case "Orange":
		return Orange
	}
	return White
}

func main() {
	var s []string
	var winArt string = `                         ....::::
                 ....::::::::::::
        ....:::: ::::::::::::::::
....:::::::::::: ::::::::::::::::
:::::::::::::::: ::::::::::::::::
:::::::::::::::: ::::::::::::::::
:::::::::::::::: ::::::::::::::::
:::::::::::::::: ::::::::::::::::
................ ................
:::::::::::::::: ::::::::::::::::
:::::::::::::::: ::::::::::::::::
:::::::::::::::: ::::::::::::::::
'''':::::::::::: ::::::::::::::::
        '''':::: ::::::::::::::::
                 ''''::::::::::::
                         ''''::::`
	var winArtSmall string = `                       ...:::
               ....::::::::::
        ...::: ::::::::::::::
...::::::::::: ::::::::::::::
:::::::::::::: ::::::::::::::
:::::::::::::: ::::::::::::::
:::::::::::::: ::::::::::::::
.............. ..............
:::::::::::::: ::::::::::::::
:::::::::::::: ::::::::::::::
:::::::::::::: ::::::::::::::
:::::::::::::: ::::::::::::::
'''':::::::::: ::::::::::::::
        '''::: ::::::::::::::
                 ''::::::::::
                     ''''::::`
	user, err := user.Current()
	if _, err := os.Stat(user.HomeDir + "\\.winfetch.json"); os.IsNotExist(err) {
		config := newConfig()
		file, _ := json.MarshalIndent(config, "", " ")
		_ = ioutil.WriteFile(user.HomeDir + "\\.winfetch.json", file, 0644)
		fmt.Println("No Config File Found! This must be the first time running! Creating Config at: " + user.HomeDir + "\\.winfetch.json")
	}
	config := Config{}
	configFile, err := os.Open(user.HomeDir + "\\.winfetch.json")
	if err != nil {
		log.Fatal("Error Opening Config File", err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil{
		log.Fatal("Error Parsing Config File: ", err.Error())
	}
	var winArtResult []string
	if (config.UseSmallAscii){
		winArtResult = strings.Split(winArtSmall, "\n")
	} else {
		winArtResult = strings.Split(winArt, "\n")
	}
	if (config.UseCustomAscii){
		content, err := ioutil.ReadFile(config.CustomAsciiPath)
		if err != nil {
			log.Fatal(err)
		}
		winArtResult = strings.Split(string(content), "\n")
	}

	title := xterm256.Green
	ascii := xterm256.Blue
	sep := xterm256.Red
	userc := xterm256.Red
	info := xterm256.White

	if (!config.UseDefaultColors) {
		title = getCustomColor(config.TitleColor)
		ascii = getCustomColor(config.AsciiColor)
		sep = getCustomColor(config.SepColor)
		userc = getCustomColor(config.UserColor)
		info = getCustomColor(config.InfoColor)
	}
	s = generateInfo(config, title, info, userc, sep)
	if (config.ShowAscii){
		scanner := bufio.NewScanner(strings.NewReader(""))
		if (config.UseSmallAscii){
			scanner = bufio.NewScanner(strings.NewReader(winArtSmall))
		} else {
			scanner = bufio.NewScanner(strings.NewReader(winArt))
		}
		if (config.UseCustomAscii){
			content, err := ioutil.ReadFile(config.CustomAsciiPath)
			if (err != nil) {
				log.Fatal(err)
			}
			text := string(content)
			scanner = bufio.NewScanner(strings.NewReader(text))
		}
		index := 0
		for i, str := range s {
			fmt.Println(xterm256.Sprint(ascii, winArtResult[i]) + "    " + str)
		}
		for scanner.Scan() {
			if index >= len(s) {
				fmt.Println(xterm256.Sprint(ascii, scanner.Text()))
			}
			index++
		}
	} else {
		for _, str := range s {
			fmt.Println(str)
		}
	}
}
