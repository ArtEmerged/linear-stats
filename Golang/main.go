package main

import (
	"liner-stats/internal"
)

func main() {
	internal.ParseData()
	internal.Dataprocessing()
	internal.Output()
}
