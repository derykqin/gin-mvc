package log

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogrus() {
	log.SetFormatter(&log.JSONFormatter{}) //text formatter or json formatter
	log.SetReportCaller(true)              //添加记录日志方法
	log.SetLevel(log.InfoLevel)            //info level
	writer := getFileWriter("./logs/log.txt")
	log.SetOutput(io.MultiWriter(os.Stdout, writer)) //输出多个通道
}
func getFileWriter(p string) *os.File {
	writer, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	return writer

}
