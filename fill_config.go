package main

import (
	"encoding/json"
	"log"
	"os"

	. "github.com/amar-jay/nat_wsl/pkg/config"
	"gopkg.in/yaml.v3"
)

func fill_random_config() {
	// filll config with random data to see structure
	file_path := "config.yaml"
	forwarded := Forwarding{}
	forwarded.Wsl.Listenport = 1234
	forwarded.Protocol = "tcp"
	forwarded.Type = "v4tov4"
	forwarded.Wsl.Listenhost = "localhost"
	forwarded.Remote.Connectport = 1234
	forwarded.Remote.Connectip = "localhost"

	conf := Config{
		"onestop": forwarded,
		"waaguan": forwarded,
	}

	// write to file
	file, err := os.Create(file_path)
	if err != nil {
		log.Fatalf("Can't create file: %v", err)
	}
	defer file.Close()
	d, err := yaml.Marshal(&conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	println(string(d))
	file.Write(d)
	println("----------------------------")
	// read from file
	content, err := os.ReadFile(file_path)
	yaml.Unmarshal(content, &conf)
	ee, _ := json.MarshalIndent(conf, "", "  ")
	println(string(ee))
}
