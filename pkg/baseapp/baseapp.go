package baseapp

import (
	"os"
	"os/signal"
	"syscall"
)

var app *BaseApp

func GetApp() *BaseApp {
	return app
}

type destructFunc func() error

type BaseApp struct {
	exit             chan os.Signal
	destructFuncList []destructFunc
}

func newApp() *BaseApp {
	return &BaseApp{
		exit:             make(chan os.Signal, 1),
		destructFuncList: []destructFunc{},
	}
}

func (app *BaseApp) onInitialize() {
	signal.Notify(app.exit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-app.exit
		app.destruct()
	}()
}

func (app *BaseApp) onDestruct() error {
	for _, f := range app.destructFuncList {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func (app *BaseApp) destruct() {
	if err := app.onDestruct(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func (app *BaseApp) Exit() {
	close(app.exit)
}

func (app *BaseApp) Defer(destructors ...destructFunc) {
	app.destructFuncList = append(app.destructFuncList, destructors...)
}

func init() {
	app = newApp()
	app.onInitialize()
}
