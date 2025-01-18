package main 

import (
	"fmt"
	"time"
	//"math/rand"
)

var (
	Debug  bool 
	LogLevel = "info"
	startUpTime = time.Now()
	seed int64 = 1234456789
)

func main() {
	
	fmt.Println(Debug, LogLevel, startUpTime)
	//fmt.Println(rand.NewSource(seed))
	Debug, LogLevel, startUpTime := getConfig()
	fmt.Println("Debug value is : ", Debug, "Loglevel is : ", LogLevel, "Start uptime: ", startUpTime)
}

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

//Debug, LogLevel, startUpTime := getConfig()