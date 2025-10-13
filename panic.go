package main

import "os"

func main() {
	panic("a problem")

	_, err := os.Create("/etc/panic.txt")
	if err != nil {
		panic(err)
	}
}
