package main

import (
	"trackergo/tracker"
)

func main() {
	tracker.CreateDefaultCategories(make(map[string]*tracker.Category))

}
