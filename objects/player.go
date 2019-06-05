package objects

import (
	"syscall/js"
	"math"
	"../control"
)

// Player representsthe main player object
type Player struct {
	Context js.Value
	X       int
	Y       int
	angle	float64
	speedx float64
	speedy float64
	spawnChan chan Object
	
}

// NewPlayer returns a new player object
func NewPlayer(context js.Value, x int, y int, spawnChan chan Object) *Player {
	p := &Player{
		Context: context,
		X:       x,
		Y:       y,
		angle: 0,
		speedx: 0,
		speedy: 0,
		spawnChan: spawnChan,
	}
	return p
}

type pv struct {x, y float64}

func calcRotatedPoint(angle float64, x int, y int, originx int, originy int)  pv{
	newX := math.Cos(angle) * float64(x - originx) - math.Sin(angle) * float64(y-originy) + float64(originx);
	newY := math.Sin(angle) * float64(x - originx) + math.Cos(angle) * float64(y - originy) + float64(originy);
	return pv{
		newX, newY,
	}
}

// Draw renders the player on the screen
func (p *Player) Draw() {
	angleRad := p.angle * (math.Pi/180)
	// _= angleRad
	p.Context.Call("beginPath")
	pt := calcRotatedPoint(angleRad, p.X + 10, p.Y + 10, p.X, p.Y)
	p.Context.Call("moveTo", pt.x, pt.y)
	pt2 := calcRotatedPoint(angleRad, p.X - 10, p.Y + 10, p.X, p.Y)
	p.Context.Call("lineTo", pt2.x, pt2.y)
	pt3 := calcRotatedPoint(angleRad, p.X, p.Y - 15, p.X, p.Y)
	p.Context.Call("lineTo", pt3.x,pt3.y)
	p.Context.Call("lineTo", pt.x, pt.y)
	p.Context.Call("fill")

	p.Context.Call("closePath")
}

// Step Handles misc player update functions
func (p *Player) Step(keys control.KeysPressed){

	accel := 1.0
	angleRad := p.angle * (math.Pi/180.0)
	maxSpeed := 10.0
	if(p.X > 800){
		p.X =0
	}else if (p.X < 0){
		p.X = 800
	}
	if(p.Y > 800){
		p.Y = 0
	}else if(p.Y <0){
		p.Y = 800
	}
	if(keys.Left){
		p.angle -= 4;
	}else if(keys.Right){
		p.angle += 4;
	}
	if(keys.Up){
		p.speedx += accel*math.Sin(angleRad)
		if(p.speedx > maxSpeed){
			p.speedx = maxSpeed
		}else if(p.speedx < -maxSpeed){
			p.speedx = -maxSpeed
		}
		p.speedy -= accel*math.Cos(angleRad)
		if(p.speedy > maxSpeed){
			p.speedy = maxSpeed
		}else if(p.speedy < -maxSpeed){
			p.speedy = -maxSpeed
		}
	}
	
	p.X += int(p.speedx)
	p.Y += int(p.speedy)
	if(keys.Space){
		b := &Bullet{
			Context: p.Context,
			X: int(p.X),
			Y: int(p.Y),
			Angle: 180.0-p.angle,
		}

		p.spawnChan <- b

	}
}
