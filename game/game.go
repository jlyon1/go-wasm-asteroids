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
	Objects []objects.Object
	spawnChan chan objects.Object
}

// Render draws the canvas
func (g *Game) Render(){
	g.Context.Call("clearRect", 0, 0,g.Width, g.Height)
	g.Player.Step(g.Keys)
	g.Player.Draw()
	for _,o := range g.Objects{
		(o).Step()
		(o).Draw()
	}

	// b.Draw()
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
	if(keycode == " "){
		g.Keys.Space = true
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
	if(keycode == " "){
		g.Keys.Space = false
	}
}

func (g *Game) processRecv(spawnChan chan objects.Object){
	for obj := range g.spawnChan{
		g.Objects = append(g.Objects, obj)
	}
}

// Init inits the game
func (g *Game) Init(){
	g.spawnChan = make(chan objects.Object, 100)
	go g.processRecv(g.spawnChan)
	player := objects.NewPlayer(g.Context, 40,40, g.spawnChan)
	asteroid := &objects.Asteroid{
		Context: g.Context,
		X: 100,
		Y: 100,
		Angle: 25.0,
		Size: 22,
		Speed: 2,
	}
	g.spawnChan <- asteroid
	g.Player = player
}
