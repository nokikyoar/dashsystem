package tools

import (
	"io"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// ExecCmd 执行命令,返回结果
func ExecCmd(cmd string) (string, error) {
	cmdExec := exec.Command("sh", "-c", cmd)
	stdin, _ := cmdExec.StdinPipe()
	stdout, _ := cmdExec.StdoutPipe()
	if err := cmdExec.Start(); err != nil {
		log.Fatalf("cmd failed with '%s'", err)
	}
	err := stdin.Close()
	if err != nil {
		log.Fatalf("cmd failed with '%s'", err)
	}
	bytes, err := io.ReadAll(stdout)
	if err != nil {
		log.Fatalf("cmd failed with '%s'", err)
	}
	err = cmdExec.Wait()
	if err != nil {
		log.Fatalf("cmd failed with '%s'", err)
	}
	return string(bytes), err

}

// TrimSpace 去除字符串中的空格，换行符，制表符
func TrimSpace(str string) string {
	// 去除字符串中的空格，换行符，制表符，%
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "%", "", -1)
	return str
}

// BytesToGB bytes 转换为 GB
func BytesToGB(str string) string {
	// bytes 转换为 MB
	bytes, _ := strconv.ParseFloat(str, 64)
	gb := bytes / 1024 / 1024 / 1024
	return strconv.FormatFloat(gb, 'f', 2, 64)
}
