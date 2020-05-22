package main

import (
	"math"
	"strconv"
)

func RoundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return
}

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