package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

const DefaultConfigPath = "config.json"

type Config struct {
	ExePath      string
	DelaySeconds float64
}

func main() {
	var customConfigPath string
	config := new(Config)
	flag.StringVar(&customConfigPath, "config", "", "config file path")
	flag.StringVar(&config.ExePath, "exe", "", "engine exe path")
	flag.Float64Var(&config.DelaySeconds, "delay", 0, "delay seconds")
	flag.Parse()
	if customConfigPath != "" {
		config = loadConfig(customConfigPath)
	} else if config.ExePath == "" {
		config = loadConfig(DefaultConfigPath)
	}

	cmd := exec.Command(config.ExePath)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	go func() {
		r := bufio.NewReader(stdout)
		hasPrefix := false
		for {
			line, prefix, err := r.ReadLine()
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			if !hasPrefix && strings.HasPrefix(string(line), "bestmove ") {
				time.Sleep(time.Duration(config.DelaySeconds * float64(time.Second)))
			}
			os.Stdout.Write(line)
			if !prefix {
				os.Stdout.Write([]byte{'\n'})
			}
			hasPrefix = prefix
		}
	}()
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			stdin.Write(buf[:n])
		}
	}()
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func loadConfig(filePath string) *Config {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		panic(err)
	}
	return &config
}
