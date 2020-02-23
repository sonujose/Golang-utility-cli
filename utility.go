package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func listtopic(a int, b int) int {
	cmdStr := "docker run --name listtopic listtopic:0.1.0"
	out, err := exec.Command("cmd.exe", "/C", cmdStr).Output()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(out))
	} else {
		fmt.Printf("%s", out)
	}
	res:= a+b
	fmt.Println("result is", res)
    return res
}

func mul(a int, b int) int {
	res:= a*b
	fmt.Println("result is", res)
    return res
}

func main() {
	a, err := strconv.Atoi(os.Args[2])
	b, err := strconv.Atoi(os.Args[3])

	if err != nil {
        fmt.Println(err)
    } else {
        switch os.Args[1] {
		case "listtopic":
			listtopic(a, b)
		
		case "mul":
			mul(a, b)
		}
	
	}
	cmdStr := "docker rm listtopic"
	out, err := exec.Command("cmd.exe", "/C", cmdStr).Output()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(out))
	} else {
		fmt.Printf("%s", out)
	}
}