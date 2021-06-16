package main

import (
	"fmt"

	"mirakurun-exporter/collector"
)

func main() {
	status := collector.NewStatusCollector()
	fmt.Println(status)
}
