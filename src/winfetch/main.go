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


	"github.com/jaypipes/ghw"
	"golang.org/x/sys/windows/registry"
	"github.com/fatih/color"
)

func main() {
	var s []string
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed, color.Bold).SprintFunc()
	blue := color.New(color.FgBlue, color.Bold).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
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
	winArtResult := strings.Split(winArt, "\n")
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
	s = append(s, red(user.Username))
	s = append(s, red("----------------"))
	if (config.ShowPhysMem){
		s = append(s, green("Physical Memory: ") + ByteFormat(float64(memory.TotalPhysicalBytes), 1))
	}
	if (config.ShowUsableMem){
		s = append(s, green("Usable Phys Mem.: ") + ByteFormat(float64(memory.TotalUsableBytes), 1))
	}
	if (config.ShowTotalCPUCores || config.ShowTotalCPUThreads){
		cpu, err := ghw.CPU()
		if err != nil {
			fmt.Printf("Error getting CPU info: %v", err)
		}
		if (config.ShowTotalCPUCores){
			s = append(s, green("Total CPU Cores: ") +  fmt.Sprint(cpu.TotalCores))
		}
		if (config.ShowTotalCPUThreads){
			s = append(s, green("Total CPU Threads: ") + fmt.Sprint(cpu.TotalThreads))
		}
	}
	if (config.ShowWindowsVersion){
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	
		pn, _, err := k.GetStringValue("ProductName")
		if err != nil {
			log.Fatal(err)
		}
		s = append(s, green("Windows Version: ") + pn)
	}
	if (config.ShowTotalDiskSize){
		bi, err := ghw.Block()
		if err != nil {
			fmt.Printf("Error getting disk info: %v", err)
		}
		s = append(s, green("Total Disk Size: ") + ByteFormat(float64(bi.TotalPhysicalBytes), 1))
	}
	if (config.ShowGPUS){
		gpu, err := ghw.GPU()
		if err != nil {
			fmt.Printf("Error getting GPU info: %v", err)
		}
		gpuin := 0
		for _, c := range gpu.GraphicsCards {
			s = append(s, green("GPU #" + fmt.Sprint(gpuin)) +": " + c.DeviceInfo.Product.Name)
			gpuin++
		}
	}
	if (config.ShowWindowsVersion){
		
	}
	if (config.ShowBios){
		bios, err := ghw.BIOS()
		if err != nil {
			fmt.Printf("Error getting BIOS info: %v", err)
		}
		s = append(s, green("BIOS: ") + bios.Vendor)
	}
	if (config.ShowBaseboard){
		bb, err := ghw.Baseboard()
		if err != nil {
			fmt.Printf("Error getting BB info: %v", err)
		}
		s = append(s, green("Baseboard: ")  + bb.Vendor)
	}
	s = append(s, "")
	s = append(s, "     " + red("███") + green("███") + yellow("███") + blue("███") + magenta("███") + cyan("███"))
	scanner := bufio.NewScanner(strings.NewReader(winArt))
	index := 0
	for i, str := range s {
		fmt.Println(blue(winArtResult[i]) + "    " + str)
	}
	for scanner.Scan() {
		if index >= len(s) {
			fmt.Println(blue(scanner.Text()))
		}
		index++
	}
}
