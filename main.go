package main

import (
	"changeme/base"
	"embed"
	"fmt"
	"sync"
	"time"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	// app := NewApp()
	// testrobotgo()
	testwin32()

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

func testwin32() {
	// 获取当前鼠标位置
	currentX, currentY := base.GetMousePos()
	fmt.Printf("Current Mouse Position: (%d, %d)\n", currentX, currentY)

	base.SetMousePos(1475, 800)
	base.MoveMouse(-100, 0)

	// 等待一段时间，以便观察鼠标移动
	time.Sleep(3 * time.Second)

	// 获取移动后的鼠标位置
	newX, newY := base.GetMousePos()
	fmt.Printf("New Mouse Position: (%d, %d)\n", newX, newY)

	// 恢复鼠标原始位置
	base.SetMousePos(currentX, currentY)
}
