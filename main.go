package main

import (
	"log"
	"os"
	"sample/search"
	_ "sample/matchers"
)

func init()  {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("a")
}