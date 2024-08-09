package baseapp

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetApp(t *testing.T) {
	assert.Equal(t, GetApp(), GetApp(), "The app should be globally unique")
	assert.NotEqual(t, *app, BaseApp{}, "The app should be initialized before being imported")
}

func TestBaseapp_Defer(t *testing.T) {
	destructors := []destructFunc{
		func() error {
			return nil
		},
		func() error {
			return errors.New("someting wrong when destructing")
		},
	}
	app.Defer(destructors...)
	assert.Equal(t, len(destructors), len(app.destructFuncList), "The app should be able to add destructor")
}

func TestBaseapp_Exit(t *testing.T) {
	app.Exit()
	signal, ok := <-app.exit
	assert.Equal(t, signal, nil, "app return nil when exiting")
	assert.Equal(t, ok, false, "app's signal channel is closed when exiting")
}
