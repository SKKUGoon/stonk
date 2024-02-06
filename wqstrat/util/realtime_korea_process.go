package util

import (
	"errors"
	"fmt"
	"strings"
)

func korExecMessage(msg string) error {
	strs := strings.Split(msg, "|")
	if len(strs) < 4 {
		return errors.New(fmt.Sprintf("wrong message string input %s", msg))
	}

	executeData := strings.Split(strs[3], "^")
	fmt.Println(executeData[0], executeData[1], executeData[2], executeData[18])

	return nil
}
