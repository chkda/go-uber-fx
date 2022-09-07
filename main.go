package main

import (
	"github.com/chkda/go-uber-fx-tutorial/bundlefx"
	"github.com/chkda/go-uber-fx-tutorial/httpHandler"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		bundlefx.Module,
		fx.Invoke(httpHandler.New),
	).Run()
}
