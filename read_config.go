package main

import (
	//"crypto/cipher"
	"fmt"
	"os"
	"log"
)

type Config struct {

}

func readConf(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
		//os.Exit()
	}
	defer file.Close()
	cfg := &Config{}
return cfg, nil
	
}
func setupLogging() {
	app_log, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("%v", err)
	}
	log.SetOutput(app_log)
}
func main() {
	setupLogging()
	cfg, err := readConf("path/to/config.toml")
	if err != nil {
		//fmt.Fprintf(os.Stderr, "error: %s\n", err)
		//log.Printf("error: %+v", err)
		os.Exit(1)
	}
	fmt.Println(cfg)
}