package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/samifruit514/go-consul-client/src/client"
)

var filepath = flag.String("file", "", "the path to the json file")
var namespace = flag.String("namespace", "", "the consul namespace to use as a prefix")
var consulAddr = flag.String("consul", "", "the consul address to use as a prefix")

func main() {
	flag.Parse()
	if filepath == nil || *filepath == "" {
		log.Printf("Missing parameter -file")
		printHelp()
	}

	if consulAddr == nil || *consulAddr == "" {
		log.Printf("Missing parameter -consul")
		printHelp()
	}

	if _, err := os.Stat(*filepath); os.IsNotExist(err) {
		log.Fatalf("Given file does not exist: %s", *filepath)
	}
	data, err := ioutil.ReadFile(*filepath)
	if err != nil {
		log.Fatalf("Error reading file %s : %v", *filepath, err)
	}

	loader, err := client.NewCachedLoader(*namespace, *consulAddr)
	if err != nil {
		log.Fatalf("Error creating loader: %v", err)
	}

	err = loader.Import(data)
	if err != nil {
		log.Fatalf("Error importing data: %v", err)
	}
	log.Printf("Json from %s successfully loaded", *filepath)
}

func printHelp() {
	log.Println("Consul Client importer will import a json file into a consul KV store.")
	log.Println("Usage: ")
	log.Println("bin/importer -file /path/to/json/file -namespace dev/config")
	log.Println(" -file is the path to a json file to import")
	log.Println(" -namespace is a prefix to use in consul")
	log.Println(" -consul is the address for consul. e.g. 172.17.8.101:8500")
	os.Exit(1)
}
