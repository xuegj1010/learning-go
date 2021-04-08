package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

var s1 []string

func main() {
	s1 = readFile("abc.txt")
	s2 := pingJie(s1)

	for _, c := range s2 {
		cmd := exec.Command("cmd", "/C", c)
		output, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(output))
	}

}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close() //关闭

	line := bufio.NewReader(file)
	for {
		content, _, err := line.ReadLine()
		if err == io.EOF {
			break
		}
		s1 = append(s1, string(content))
	}
	return s1

}

func pingJie(s1 []string) []string {
	var aaa = "echo "
	var s2 []string
	for _, s := range s1 {
		s2 = append(s2, fmt.Sprintf("%s%s", aaa, s))
	}

	return s2
}
