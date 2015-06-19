package utils

import (
	"URLServer/logs"
	"fmt"
)

func Check(e error) {
	if e != nil {
		fmt.Println(e.Error())
		logger.Log.Error("Server err: %v", e)
	}
}
