package common

import (
	"io"
	"log"
	"os"
	"testing"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	Info = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}
func TestLog(t *testing.T) {
	Info.Println("飞雪无情的博客:", "http://www.flysnow.org")
	Warning.Printf("飞雪无情的微信公众号：%s\n", "flysnow_org")
	Error.Println("欢迎关注留言")
}
