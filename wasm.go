package main

import (
	"syscall/js"
	"./game"
	"sync"
)

var (
	window, doc, body, canvas, laserCtx, beep js.Value
	windowSize                                struct{ w, h float64 }
	
	gameobj game.Game
	gs = gameState{laserSize: 35, directionX: 3.7, directionY: -3.7, laserX: 40, laserY: 40}
)
var wg sync.WaitGroup
func main() {


	setup()

	// declare renderer at compile time
	var renderer js.Func

	renderer = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		updateGame()
		// for the 60fps anims
		window.Call("requestAnimationFrame", renderer)
		return nil
	})
	window.Call("requestAnimationFrame", renderer)

	// let's handle that mouse/touch down
	var keyEventHandler js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// updatePlayer(args[0])
		log(args[0])
		gameobj.HandleKeysDown(args[0].Get("key").String())
		return nil
	})
	var keyEventUpHandler js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// updatePlayer(args[0])
		gameobj.HandleKeysUp(args[0].Get("key").String())
		return nil
	})
	window.Call("addEventListener", "keydown", keyEventHandler)
	window.Call("addEventListener", "keyup", keyEventUpHandler)
	wg.Add(1)
	wg.Wait()

}

func updateGame() {
	gameobj.Render()

}


func setup() {
	window = js.Global()
	doc = window.Get("document")
	body = doc.Get("body")

	windowSize.h = window.Get("innerHeight").Float()
	windowSize.w = window.Get("innerWidth").Float()

	canvas = doc.Call("createElement", "canvas")
	canvas.Set("height", 800)
	canvas.Set("width", 800)
	body.Call("appendChild", canvas)
	log("hi")
	laserCtx = canvas.Call("getContext", "2d")
	laserCtx.Set("fillStyle", "white")
	gameobj = game.Game{
		Context: laserCtx,
		Width: windowSize.w,
		Height: windowSize.h,
	}
	gameobj.Init()

}

func playSound() {
}

type gameState struct{ laserX, laserY, directionX, directionY, laserSize float64 }

func log(args ...interface{}) {
	window.Get("console").Call("log", args...)
}