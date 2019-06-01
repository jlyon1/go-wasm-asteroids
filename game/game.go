package game

import "../objects"
import "syscall/js"
import "../control"

// Game represents the game
type Game struct {
	Context js.Value
	Player *objects.Player
	Width float64
	Height float64
	Keys control.KeysPressed
}

// Render draws the canvas
func (g *Game) Render(){
	g.Context.Call("clearRect", 0, 0,g.Width, g.Height)
	g.Player.Draw()
	g.Player.Step(g.Keys)
}


// HandleKeysDown handles the keystrokes down for the game
func (g *Game) HandleKeysDown(keycode string){
	
	if(keycode == "ArrowRight"){
		g.Keys.Right = true
	}
	if(keycode == "ArrowLeft"){
		g.Keys.Left = true
	}
	if(keycode == "ArrowUp"){
		g.Keys.Up = true
	}
}

// HandleKeysUp handles the keystrokes down for the game
func (g *Game) HandleKeysUp(keycode string){
	
	if(keycode == "ArrowRight"){
		g.Keys.Right = false
	}
	if(keycode == "ArrowLeft"){
		g.Keys.Left = false
	}
	if(keycode == "ArrowUp"){
		g.Keys.Up = false
	}
}
