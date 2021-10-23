package response

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"courier/pkg/utils"
)

func LogErr(path string, status int, message string, data interface{}) error {
	log_path, _ := filepath.Abs("logs/err")
	year, month, day, hour, min, sec := utils.GetTime()

	log := []byte(day + "-" + month + "-" + year + " " + hour + ":" + min + ":" + sec + " => " + path + "\n" + strconv.Itoa(status) + " => " + message + "\n" + fmt.Sprintf("%v", data) + "\n\n")

	filename := "err-" + year + "-" + month + "-" + day + ".log"

	f, err := os.OpenFile(log_path+"\\"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer f.Close()

	_, err = f.Write([]byte(log))

	return err
}
