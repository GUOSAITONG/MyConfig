package main

import (
	"github.com/GUOSAITONG/MyConfig/inits"
)

func main() {
	inits.InitViper()
	inits.InitMysql()
	inits.InitRedis()
	inits.InitMongoDB()
	inits.InitEs()
}
