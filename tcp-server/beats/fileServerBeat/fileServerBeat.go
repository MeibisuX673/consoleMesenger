package fileServerBeat

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Log struct {
	Message string
	Time    time.Time
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

func CreateLog(text string) {

	logs := Log{
		Message: text,
		Time:    time.Now(),
	}

	js, err := json.Marshal(logs)
	if err != nil {
		log.Fatal(err)
	}

	writeLog(js)
}

func writeLog(logs []byte) {

	file, err := os.OpenFile("./../host_metrics_app/host_metrics_app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(file)

	_, err = file.Write(logs)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write([]byte("\n"))
	if err != nil {
		log.Fatal(err)
	}

}
