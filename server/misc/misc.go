package misc

import (
	"os"
)

func GetCmdDir() string {
	cmdDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return cmdDir
}
