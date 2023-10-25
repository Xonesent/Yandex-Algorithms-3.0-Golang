package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line1, _ := reader.ReadString('\n')
	hour, _ := strconv.Atoi(line1[0:2])
	min, _ := strconv.Atoi(line1[3:5])
	sec, _ := strconv.Atoi(line1[6:8])
	time1 := hour*3600 + min*60 + sec

	line2, _ := reader.ReadString('\n')
	hour, _ = strconv.Atoi(line2[0:2])
	min, _ = strconv.Atoi(line2[3:5])
	sec, _ = strconv.Atoi(line2[6:8])
	time2 := hour*3600 + min*60 + sec

	line3, _ := reader.ReadString('\n')
	hour, _ = strconv.Atoi(line3[0:2])
	min, _ = strconv.Atoi(line3[3:5])
	sec, _ = strconv.Atoi(line3[6:8])
	time3 := hour*3600 + min*60 + sec

	var plustime int
	if time3 >= time1 {
		plustime = time3 - time1
	} else {
		plustime = 86400 - time1 + time3
	}
	time2 += int(math.Round(float64(plustime) / float64(2)))
	if time2 > 86400 {
		time2 -= 86400
	}

	fmt.Printf("%02d:%02d:%02d", time2/3600, (time2%3600)/60, time2%60)
}
