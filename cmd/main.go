package main

import "zhihu/boot"

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.DatabaseInit()
	boot.InitRouters()
}
