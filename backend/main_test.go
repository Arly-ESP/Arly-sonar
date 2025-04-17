package main

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/arly/arlyApi/config"
	"github.com/arly/arlyApi/database"
	"github.com/arly/arlyApi/index"
	"github.com/arly/arlyApi/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestMainFunction_Success(t *testing.T) {
	patchInitLogger := gomonkey.ApplyFunc(utilities.InitializeLogger, func() {})
	defer patchInitLogger.Reset()

	patchLogInfo := gomonkey.ApplyFunc(utilities.LogInfo, func(msg string) {})
	defer patchLogInfo.Reset()

	patchLogFatal := gomonkey.ApplyFunc(utilities.LogFatal, func(msg string, err error) {
		panic("LogFatal called: " + msg)
	})
	defer patchLogFatal.Reset()

	patchConnectDb := gomonkey.ApplyFunc(database.ConnectDb, func() {
		database.Database = database.DbInstance{Db: nil}
	})
	defer patchConnectDb.Reset()

	patchPrepareApp := gomonkey.ApplyFunc(index.PrepareApp, func(app *fiber.App) {})
	defer patchPrepareApp.Reset()

	patchServerUrl := gomonkey.ApplyFunc(config.ServerUrl, func() string {
		return ":3000"
	})
	defer patchServerUrl.Reset()

	patchListen := gomonkey.ApplyMethod(reflect.TypeOf(&fiber.App{}), "Listen", func(app *fiber.App, addr string) error {
		return nil
	})
	defer patchListen.Reset()

	assert.NotPanics(t, func() {
		main()
	}, "main() should run without panicking")
}
