package main

import (
	"changeme/base/config"
	"changeme/base/listener"
	"changeme/base/mouse"
	"embed"
	"fmt"
	"sync"
	"time"

	hook "github.com/robotn/gohook"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// app := NewApp()

	// 1. load config
	config.Init()
	// 2. register listener event
	testKeyboardListener()
	testMouseMove()

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

	var wg sync.WaitGroup
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	wg.Add(1)
	// hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
	// 	// defer wg.Done()
	// 	// wg.Add(1)
	// 	fmt.Println("ctrl-shift-q")
	// 	hook.End()
	// })
	// fmt.Println("--- Please press w---")
	// hook.Register(hook.KeyDown, []string{"1"}, func(e hook.Event) {
	// 	locx, locy := robotgo.Location()
	// 	fmt.Println("before:" + fmt.Sprint(locx) + "," + fmt.Sprint(locy))
	// 	// robotgo.Move(locx, locy+10)
	// 	robotgo.MoveMouse(locx, locy+1)
	// 	locx, locy = robotgo.Location()
	// 	fmt.Println("after:" + fmt.Sprint(locx) + "," + fmt.Sprint(locy))

	// })

	// s := hook.Start()
	// <-hook.Process(s)
	wg.Wait()
	// robotgo.Move(1200, 20)
	// robotgo.DragSmooth(700, 200)

}

func testMouseMove() {
	// 获取当前鼠标位置
	currentX, currentY := mouse.GetMousePos()
	fmt.Printf("Current Mouse Position: (%d, %d)\n", currentX, currentY)

	mouse.SetMousePos(1475, 800)
	mouse.MoveMouse(-100, 0)

	// 等待一段时间，以便观察鼠标移动
	time.Sleep(3 * time.Second)

	// 获取移动后的鼠标位置
	newX, newY := mouse.GetMousePos()
	fmt.Printf("New Mouse Position: (%d, %d)\n", newX, newY)

	// 恢复鼠标原始位置
	mouse.SetMousePos(currentX, currentY)
}

func testKeyboardListener() {
	listener.RegisterFromConfig()
	listener.Start()
	// fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	// hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
	// 	fmt.Println("ctrl-shift-q")
	// 	hook.End()
	// })

	// fmt.Println("--- Please press
	// hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
	// 	fmt.Println("w")
	// })

	// s := hook.Start()
}

func low() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		fmt.Println("hook: ", ev)
	}
}

func testConfig() {
	config.Init()
}
