package main

import (
	"embed"
	"fmt"

	"github.com/go-vgo/robotgo"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	// app := NewApp()
	testrobotgo()

	// Create application with options
	// err := wails.Run(&options.App{
	// 	Title:  "MouseK",
	// 	Width:  1024,
	// 	Height: 768, AssetServer: &assetserver.Options{
	// 		Assets: assets,
	// 	},
	// 	BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
	// 	OnStartup:        app.startup,
	// 	Bind: []interface{}{
	// 		app,
	// 	},
	// })

	// if err != nil {
	// 	println("Error:", err.Error())
	// }
}

func testrobotgo() {
	fmt.Println("test robot go")
	robotgo.Move(500, 20)

}
