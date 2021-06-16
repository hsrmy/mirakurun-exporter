package main

import (
	"fmt"

	"mirakurun-exporter/collector"
)

func main() {
	status := collector.NewStatusCollector()
	version := collector.NewVersionCollector()
	fmt.Println(status)
	fmt.Println(version)
}
