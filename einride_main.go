package main

import (
	"einride_test/internal"
)

func main() {
	mutexMap := internal.NewMutexMap()

	httpa := &internal.MutexMapHTTPAdapter{mutexMap: mutexMap}
	internal.NewApplication(*httpa)
}
