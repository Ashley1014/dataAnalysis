package main

import (
	_ "dataAnalysis/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func inc(i int) int {
	return i+1
}

func main() {
	beego.AddFuncMap("inc", inc)
	beego.Run()
}

