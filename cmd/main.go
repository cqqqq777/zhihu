package main

import (
	"zhihu/boot"
	"zhihu/cron"
)

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.DatabaseInit()
	cron.Cron()
	boot.InitRouters()
}
