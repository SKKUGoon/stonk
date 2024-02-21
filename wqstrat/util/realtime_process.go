package util

import (
	"fmt"
	"strings"
)

func korExecMessage(msg string) error {
	strs := strings.Split(msg, "|")
	if len(strs) < 4 {
		return fmt.Errorf("wrong message string input %s", msg)
	}

	executeData := strings.Split(strs[3], "^")
	fmt.Println(executeData[0], executeData[1], executeData[2], executeData[18])

	return nil
}

func overseaExecMessage(msg string) error {
	strs := strings.Split(msg, "|")
	if len(strs) < 4 {
		return fmt.Errorf("wrong message string input %s", msg)
	}

	executeData := strings.Split(strs[3], "^")
	fmt.Println(executeData[3], executeData[5], executeData[1], executeData[11], executeData[24])
	return nil
}
