package main

import (
	"fmt"
	"os"
	"flag"
	"os/exec"
)


func main() {
	
	url := flag.String("url", "", "Download url")

	flag.Parse()

	fmt.Println(len(os.Args), os.Args)

	var downloadUrl string = *url
	fmt.Println("Download URL:", downloadUrl)

	// for _, e := range os.Environ() {
	// 	fmt.Println(e)
	// }

	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))
}
