package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

//small parser
func main() {
	fmt.Println(checkAvailability())
}

func checkAvailability() map[string]bool {
	links := []string{"vk.com", "usa.com"}
	result := make(map[string]bool)
	go func() {
		for _, link := range links {
			fmt.Printf("RUNNING PING FOR: %s", link)
			result[link] = callPing(link)
		}
	}()

	return result
}

func callPing(link string) bool {
	if runtime.GOOS != "windows" {
		log.Fatal("The os is not supported")
	}
	_, err := exec.Command("powershell.exe", "ping", link).Output()

	if err != nil {
		fmt.Printf("Can get access to server - %s\n", link)
		return false
	}

	return true
}
