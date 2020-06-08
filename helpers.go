package main

import (
	"fmt"
	"log"
	"math"
	"os/user"
	"strconv"
	"strings"

	hst "github.com/shirou/gopsutil/host"
	ps "github.com/shirou/gopsutil/process"

	"github.com/gilliek/go-xterm256/xterm256"
	"github.com/jaypipes/ghw"
	"golang.org/x/sys/windows/registry"
)

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

// Find - Find String inside String
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// RoundUp - Round Float Up To New Value
func RoundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return
}

// ByteFormat - Format Bytes to Human Readable
func ByteFormat(inputNum float64, precision int) string {

	if precision <= 0 {
		precision = 1
	}

	var unit string
	var returnVal float64

	if inputNum >= 1000000000000000000000000 {
		returnVal = RoundUp((inputNum / 1208925819614629174706176), precision)
		unit = " YB" // yottabyte
	} else if inputNum >= 1000000000000000000000 {
		returnVal = RoundUp((inputNum / 1180591620717411303424), precision)
		unit = " ZB" // zettabyte
	} else if inputNum >= 10000000000000000000 {
		returnVal = RoundUp((inputNum / 1152921504606846976), precision)
		unit = " EB" // exabyte
	} else if inputNum >= 1000000000000000 {
		returnVal = RoundUp((inputNum / 1125899906842624), precision)
		unit = " PB" // petabyte
	} else if inputNum >= 1000000000000 {
		returnVal = RoundUp((inputNum / 1099511627776), precision)
		unit = " TB" // terrabyte
	} else if inputNum >= 1000000000 {
		returnVal = RoundUp((inputNum / 1073741824), precision)
		unit = " GB" // gigabyte
	} else if inputNum >= 1000000 {
		returnVal = RoundUp((inputNum / 1048576), precision)
		unit = " MB" // megabyte
	} else if inputNum >= 1000 {
		returnVal = RoundUp((inputNum / 1024), precision)
		unit = " KB" // kilobyte
	} else {
		returnVal = inputNum
		unit = " bytes" // byte
	}

	return strconv.FormatFloat(returnVal, 'f', precision, 64) + unit

}

func plural(count int, singular string) (result string) {
	if (count == 1) || (count == 0) {
		result = strconv.Itoa(count) + " " + singular + " "
	} else {
		result = strconv.Itoa(count) + " " + singular + "s "
	}
	return
}

func secondsToHuman(input int) (result string) {
	years := math.Floor(float64(input) / 60 / 60 / 24 / 7 / 30 / 12)
	seconds := input % (60 * 60 * 24 * 7 * 30 * 12)
	months := math.Floor(float64(seconds) / 60 / 60 / 24 / 7 / 30)
	seconds = input % (60 * 60 * 24 * 7 * 30)
	weeks := math.Floor(float64(seconds) / 60 / 60 / 24 / 7)
	seconds = input % (60 * 60 * 24 * 7)
	days := math.Floor(float64(seconds) / 60 / 60 / 24)
	seconds = input % (60 * 60 * 24)
	hours := math.Floor(float64(seconds) / 60 / 60)
	seconds = input % (60 * 60)
	minutes := math.Floor(float64(seconds) / 60)
	seconds = input % 60

	if years > 0 {
		result = plural(int(years), "year") + plural(int(months), "month") + plural(int(weeks), "week") + plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
	} else if months > 0 {
		result = plural(int(months), "month") + plural(int(weeks), "week") + plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
	} else if weeks > 0 {
		result = plural(int(weeks), "week") + plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
	} else if days > 0 {
		result = plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
	} else if hours > 0 {
		result = plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
	} else if minutes > 0 {
		result = plural(int(minutes), "minute") + plural(int(seconds), "second")
	} else {
		result = plural(int(seconds), "second")
	}

	return
}

func generateInfo(config Config, title xterm256.Color, info xterm256.Color, userc xterm256.Color, sep xterm256.Color) []string {
	var s []string
	cpu, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}
	for x := range config.Format {
		switch config.Format[x] {
		case "user":
			user, _ := user.Current()
			s = append(s, xterm256.Sprint(userc, strings.ReplaceAll(user.Username, "\\", "@")))
		case "uptime":
			uptime, err := hst.Uptime()
			if err != nil {
				log.Fatal("Failed to Get Uptime!")
			}
			uptimes := secondsToHuman(int(uptime))
			s = append(s, xterm256.Sprint(title, config.Titles.Uptime+": ")+xterm256.Sprint(info, uptimes))
		case "sep":
			s = append(s, xterm256.Sprint(sep, "--------------------------------"))
		case "mem":
			memory, err := ghw.Memory()
			if err != nil {
				fmt.Printf("Error getting memory info: %v", err)
			}
			memorySplit := strings.Split(memory.String(), "(")
			mem := strings.Split(memorySplit[1], ",")
			usableMem := strings.Split(mem[1], "usable")
			physMem := strings.Split(mem[0], "physical")
			s = append(s, xterm256.Sprint(title, config.Titles.Memory+": ")+xterm256.Sprint(info, strings.ReplaceAll(usableMem[0], "MB ", "GB")+"/"+strings.ReplaceAll(physMem[0], "MB", "GB")))
		case "cpuThreads":
			s = append(s, xterm256.Sprint(title, config.Titles.CPUThreads+": ")+xterm256.Sprint(info, fmt.Sprint(cpu.TotalThreads)))
		case "cpuCores":
			s = append(s, xterm256.Sprint(title, config.Titles.CPUCores+": ")+xterm256.Sprint(info, fmt.Sprint(cpu.TotalCores)))
		case "cpu":
			in := 0
			for x := range cpu.Processors {
				s = append(s, xterm256.Sprint(title, "CPU #"+fmt.Sprint(in)+": ")+xterm256.Sprint(info, cpu.Processors[x].Model))
			}
		case "procs":
			pids, err := ps.Pids()

			if err != nil {
				log.Fatal("Couldn't get Processes!")
			}

			s = append(s, xterm256.Sprint(title, "Proccesses Running: ")+xterm256.Sprint(info, int64(len(pids))))
		case "wversion":
			k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)

			pn, _, err := k.GetStringValue("ProductName")
			if err != nil {
				log.Fatal(err)
			}
			s = append(s, xterm256.Sprint(title, config.Titles.WindowsVersion+": ")+xterm256.Sprint(info, pn))

		case "disk":
			bi, err := ghw.Block()
			if err != nil {
				fmt.Printf("Error getting disk info: %v", err)
			}
			s = append(s, xterm256.Sprint(title, config.Titles.DiskSize+": ")+xterm256.Sprint(info, ByteFormat(float64(bi.TotalPhysicalBytes), 1)))
		case "gpus":
			gpu, err := ghw.GPU()
			if err != nil {
				fmt.Printf("Error getting GPU info: %v", err)
			}
			gpuin := 0
			if len(gpu.GraphicsCards) > 1 {
				for _, c := range gpu.GraphicsCards {
					s = append(s, xterm256.Sprint(title, config.Titles.GPUs+" #"+fmt.Sprint(gpuin)+": ")+xterm256.Sprint(info, c.DeviceInfo.Product.Name))
					gpuin++
				}
			} else {
				s = append(s, xterm256.Sprint(title, config.Titles.GPUs+": ")+xterm256.Sprint(info, gpu.GraphicsCards[0].DeviceInfo.Product.Name))
			}
		case "bios":
			bios, err := ghw.BIOS()
			if err != nil {
				fmt.Printf("Error getting BIOS info: %v", err)
			}
			s = append(s, xterm256.Sprint(title, config.Titles.Bios+": ")+xterm256.Sprint(info, bios.Vendor))

		case "baseboard":
			bb, err := ghw.Baseboard()
			if err != nil {
				fmt.Printf("Error getting BB info: %v", err)
			}
			s = append(s, xterm256.Sprint(title, config.Titles.Baseboard+": ")+xterm256.Sprint(info, bb.Vendor))

		default:
			s = append(s, "\n")
		}

	}
	s = append(s, "")
	s = append(s, "    "+xterm256.Sprint(xterm256.LightGray, "███")+xterm256.Sprint(xterm256.Red, "███")+xterm256.Sprint(xterm256.Green, "███")+xterm256.Sprint(xterm256.Yellow, "███")+xterm256.Sprint(xterm256.Blue, "███")+xterm256.Sprint(xterm256.Magenta, "███")+xterm256.Sprint(xterm256.Cyan, "███"))
	s = append(s, "    "+xterm256.Sprint(xterm256.DarkGray, "███")+xterm256.Sprint(xterm256.DarkRed, "███")+xterm256.Sprint(xterm256.DarkGreen, "███")+xterm256.Sprint(xterm256.DarkYellow, "███")+xterm256.Sprint(xterm256.DarkBlue, "███")+xterm256.Sprint(xterm256.DarkMagenta, "███")+xterm256.Sprint(xterm256.DarkCyan, "███"))
	return s
}
