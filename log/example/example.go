package main

import (
	log2 "log"
	"marmoteduTest2/log"
	"os"
)

func main() {
	// log.New(log.WithLevel(log.DebugLevel), log.WithFormatter(&log.TextColorFormatter{false}))
	// log.Info("std log")

	log.SetOptions(log.WithLevel(log.DebugLevel), log.WithFormatter(&log.TextColorFormatter{false}))
	log.Info("std log")
	//log.Debug("change std log to debug level")
	//log.SetOptions(log.WithFormatter(&log.JsonFormatter{IgnoreBasicFields:false}))
	//log.Debug("log in json format")
	//log.Info("another log in json format")
	// 格式：%c[显示方式;前景色;背景色m
	// 其中0x1B是标记，%c是颜色标记，相当于固定格式[开始定义颜色，1代表高亮，40代表黑色背景，32代表绿色前景，0代表恢复默认颜色。显示效果为
	// fmt.Printf("%c[1;40;32m%s%c[0m",0x1B,"testPrintColor",0x1B)



	os.Exit(0)
	// 输出到文件
	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		// log.Fatalln("create file test.log failed")
		log2.Fatalln("create file test.log failed")
	}
	defer fd.Close()

	l := log.New(log.WithLevel(log.InfoLevel),
		log.WithOutput(fd),
		log.WithFormatter(&log.JsonFormatter{IgnoreBasicFields:false}),
	)
	l.Info("custom log with json formatter")
}

func PkgLog() {
	log2.Printf("%s\n", "test")
}

func SimpleLog() {
	log.Infof("%s\n","test")
}
