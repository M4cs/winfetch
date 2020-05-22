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

	ps "github.com/shirou/gopsutil/process"
	hst "github.com/shirou/gopsutil/host"
	// cp "github.com/shirou/gopsutil/cpu"
	"github.com/jaypipes/ghw"
	"golang.org/x/sys/windows/registry"
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
	var winArt string = `                                ..,
                    ....,,:;+ccllll
      ...,,+:;  cllllllllllllllllll
,cclllllllllll  lllllllllllllllllll
llllllllllllll  lllllllllllllllllll
llllllllllllll  lllllllllllllllllll
llllllllllllll  lllllllllllllllllll
llllllllllllll  lllllllllllllllllll
llllllllllllll  lllllllllllllllllll
                                   
llllllllllllll  lllllllllllllllllll
llllllllllllll  lllllllllllllllllll
llllllllllllll  lllllllllllllllllll
llllllllllllll  lllllllllllllllllll
llllllllllllll  lllllllllllllllllll
 'ccllllllllll  lllllllllllllllllll
       ' \\*::  :ccllllllllllllllll
                   ''''''''''*::cll
                               ''''`
var winArtSmall string = `                              .
                  ....,,:;+ccll
      ..,+:;  cllllllllllllllll
,cclllllllll  lllllllllllllllll
llllllllllll  lllllllllllllllll
llllllllllll  lllllllllllllllll
llllllllllll  lllllllllllllllll
                               
llllllllllll  lllllllllllllllll
llllllllllll  lllllllllllllllll
llllllllllll  lllllllllllllllll
 'ccllllllll  lllllllllllllllll
     ' \\*::  :ccllllllllllllll
                 ''''''''''*::;`
	memory, err := ghw.Memory()
	if err != nil {
		fmt.Printf("Error getting memory info: %v", err)
	}
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

	s = append(s, xterm256.Sprint(userc, strings.ReplaceAll(user.Username, "\\", "@")))
	s = append(s, xterm256.Sprint(sep, "--------------------------------"))
	
	if (config.ShowUptime){
		uptime, err := hst.Uptime()
		if (err != nil) {
			log.Fatal("Failed to Get Uptime!")
		}
		uptimes := secondsToHuman(int(uptime))
		s = append(s, xterm256.Sprint(title, config.Titles.Uptime + ": ") + xterm256.Sprint(info, uptimes))
	}

	memorySplit := strings.Split(memory.String(), "(")
	mem := strings.Split(memorySplit[1], ",")
	usableMem := strings.Split(mem[1], "usable")
	physMem := strings.Split(mem[0], "physical")
	if (config.ShowMem) {
		s = append(s, xterm256.Sprint(title, config.Titles.Memory + ": ") + xterm256.Sprint(info, strings.ReplaceAll(usableMem[0], "MB ", "GB") + "/" + strings.ReplaceAll(physMem[0], "MB", "GB")))
	}
	if (config.ShowTotalCPUCores || config.ShowTotalCPUThreads || config.ShowCPU){
		cpu, err := ghw.CPU()
		if err != nil {
			fmt.Printf("Error getting CPU info: %v", err)
		}
		if (config.ShowCPU){
			in := 0
			for x := range cpu.Processors {
				s = append(s, xterm256.Sprint(title, "CPU #" + fmt.Sprint(in) + ": ") + xterm256.Sprint(info, cpu.Processors[x].Model))
			}
		}
		if (config.ShowTotalCPUCores){
			s = append(s, xterm256.Sprint(title, config.Titles.CPUCores + ": ") +   xterm256.Sprint(info, fmt.Sprint(cpu.TotalCores)))
		}
		if (config.ShowTotalCPUThreads){
			s = append(s, xterm256.Sprint(title, config.Titles.CPUThreads + ": ") +  xterm256.Sprint(info, fmt.Sprint(cpu.TotalThreads)))
		}
	}
	if (config.ShowProcessCount){
		pids, err := ps.Pids()

		if err != nil {
			log.Fatal("Couldn't get Processes!")
		}

		s = append(s, xterm256.Sprint(title, "Proccesses Running: ") + xterm256.Sprint(info, int64(len(pids))))
	}
	if (config.ShowWindowsVersion){
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	
		pn, _, err := k.GetStringValue("ProductName")
		if err != nil {
			log.Fatal(err)
		}
		s = append(s, xterm256.Sprint(title, config.Titles.WindowsVersion + ": ") + xterm256.Sprint(info, pn))
	}
	if (config.ShowTotalDiskSize){
		bi, err := ghw.Block()
		if err != nil {
			fmt.Printf("Error getting disk info: %v", err)
		}
		s = append(s, xterm256.Sprint(title, config.Titles.DiskSize + ": ") +  xterm256.Sprint(info, ByteFormat(float64(bi.TotalPhysicalBytes), 1)))
	}
	if (config.ShowGPUS){
		gpu, err := ghw.GPU()
		if err != nil {
			fmt.Printf("Error getting GPU info: %v", err)
		}
		gpuin := 0
		for _, c := range gpu.GraphicsCards {
			s = append(s, xterm256.Sprint(title, config.Titles.GPUs + fmt.Sprint(gpuin) + ": ") +  xterm256.Sprint(info, c.DeviceInfo.Product.Name))
			gpuin++
		}
	}
	if (config.ShowBios){
		bios, err := ghw.BIOS()
		if err != nil {
			fmt.Printf("Error getting BIOS info: %v", err)
		}
		s = append(s, xterm256.Sprint(title, config.Titles.Bios + ": ") +  xterm256.Sprint(info, bios.Vendor))
	}
	if (config.ShowBaseboard){
		bb, err := ghw.Baseboard()
		if err != nil {
			fmt.Printf("Error getting BB info: %v", err)
		}
		s = append(s, xterm256.Sprint(title, config.Titles.Baseboard + ": ")  + xterm256.Sprint(info, bb.Vendor))
	}
	s = append(s, "")
	s = append(s, "    " + xterm256.Sprint(xterm256.LightGray, "███") + xterm256.Sprint(xterm256.Red, "███") + xterm256.Sprint(xterm256.Green, "███") + xterm256.Sprint(xterm256.Yellow, "███") + xterm256.Sprint(xterm256.Blue, "███") + xterm256.Sprint(xterm256.Magenta, "███") + xterm256.Sprint(xterm256.Cyan, "███"))
	s = append(s, "    " + xterm256.Sprint(xterm256.DarkGray, "███") + xterm256.Sprint(xterm256.DarkRed, "███") + xterm256.Sprint(xterm256.DarkGreen, "███") + xterm256.Sprint(xterm256.DarkYellow, "███") + xterm256.Sprint(xterm256.DarkBlue, "███") + xterm256.Sprint(xterm256.DarkMagenta, "███") + xterm256.Sprint(xterm256.DarkCyan, "███"))
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
}
