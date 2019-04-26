package main

import (
	// "bufio"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Version = "0.1.0-190424" // 当前版本号
	LogFile = "kjml-fmt.log" // Log文件名
)

var (
	logger          *log.Logger // 全局Logger
	argsWithProg    []string
	argsWithoutProg []string
)

func init() {
	os.Remove(LogFile) // 删除记录文件（如果有）
	// 设置记录文件
	logFile, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(err)
	}
	// defer logFile.Close()
	// 记录文件输出和控制台输出双通
	mw := io.MultiWriter(os.Stdout, logFile)
	logger = log.New(mw, "", log.LstdFlags)

	// 处理终端传递参数
	argsWithProg = os.Args
	argsWithoutProg = os.Args[1:]
}
func main() {
	f, err := os.Open("../../../sample.kjml")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(b)
}
func getHash() (hash string) {
	salt := []byte(strconv.Itoa(rand.Int()) + strconv.FormatInt(time.Now().UnixNano(), 10))
	h := strings.ToUpper(fmt.Sprintf("%x", md5.Sum(salt)))
	return h
}
