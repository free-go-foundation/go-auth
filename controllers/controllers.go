package controllers

import "github.com/free-go-foundation/go-auth/core"

var DataBase *core.DataBase = nil

func init() {
	DataBase = core.InitDataBase()
}
