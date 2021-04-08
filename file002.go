package main

import (
	"bufio"
	"fmt"
	"os/exec"
	//"golang.org/x/text/encoding/simplifiedchinese"
)

func getOutputDirectly(name string, args ...string) (output []byte) {
	//cmd := exec.Command(name, args...)
	cmd := exec.Command(name, args...)
	output, err := cmd.Output() // 等到命令执行完, 一次性获取输出
	if err != nil {
		panic(err)
	}
	//output, err = simplifiedchinese.GB18030.NewDecoder().Bytes(output)
	//if err != nil {
	//	panic(err)
	//}
	return
}

func getOutputContinually(name string, args ...string) <-chan struct{} {
	cmd := exec.Command(name, args...)
	closed := make(chan struct{})
	defer close(closed)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdoutPipe.Close()

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() { // 命令在执行的过程中, 实时地获取其输出
			//data, err := simplifiedchinese.GB18030.NewDecoder().Bytes(scanner.Bytes()) // 防止乱码
			data := scanner.Bytes() // 防止乱码
			//if err != nil {
			//	fmt.Println("transfer error with bytes:", scanner.Bytes())
			//	continue
			//}

			fmt.Printf("%s\n", string(data))
		}
	}()

	if err := cmd.Run(); err != nil {
		panic(err)
	}
	return closed
}

func main() {
	// 效果: 等一会儿, 打印出所有输出
	//output1 := getOutputDirectly("cmd","/C", "dir")
	getOutputDirectly("cmd", "/C", "dir")
	//fmt.Printf("%s\n", output1)

	// 不断输出, 直到结束
	<-getOutputContinually("cmd", "/C", "dir")
}
