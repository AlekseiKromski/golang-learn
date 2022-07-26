package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

var OS string = runtime.GOOS
var cmd string

//small parser
func main() {
	result := checkAvailability()

	var answer string
	fmt.Print("Are you want continue? (y/n): ")
	fmt.Scan(&answer)

	if answer == "y" {
		startParse(result)
	} else {
		fmt.Printf("Bie, have a nice day")
	}

}

func checkAvailability() map[string]bool {
	links := []string{"vk.com", "usa.com"}
	result := make(map[string]bool)
	for _, link := range links {
		fmt.Printf("RUNNING PING FOR: %s\n", link)
		result[link] = callPing(link)
		fmt.Println("LOG - ", result)
	}

	return result
}

func callPing(link string) bool {

	//Exec command for win and mac
	commandBuild := []string{}
	var program string

	if OS == "windows" {
		program = "powershell.exe"
		commandBuild = append(commandBuild, "ping", link)
	} else {
		program = "ping"
		commandBuild = append(commandBuild, "-c 1", link)
	}

	_, err := exec.Command(program, commandBuild...).Output()

	if err != nil {
		return false
	}
	return true
}

func startParse(links map[string]bool) {
	fmt.Println("--------Program will run only avaliable links--------")
	for link, value := range links {
		if value {
			res, err := http.Get("https://" + link)

			if err != nil {
				log.Fatalf("Faild to send request", err)
			}

			fmt.Println("START SAVE PROCESS FOR: " + link)
			b, err := io.ReadAll(res.Body)
			os.WriteFile("./"+link+".txt", b, 0644)
		}
	}
}
