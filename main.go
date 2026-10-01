package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const cpu_freq_path = "/sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq"

func main() {
	if runtime.GOOS == "linux" {
		for {
			data, err := os.ReadFile(cpu_freq_path)
			if err != nil {
				panic(err)
			}

			str_trim := strings.TrimSpace(string(data))

			freq, err := strconv.ParseFloat(str_trim, 64)
			if err != nil {
				panic(err)
			}

			fmt.Printf("cpu_freq: %.1f\n", freq/1000.0)

			time.Sleep(time.Second * 1)
		}
	} else {
		fmt.Println("This program runs only on Linux!")
		return
	}
}

