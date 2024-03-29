package objects

import (
	"syscall/js"
	"math"
)

// Bullet representsthe main bullet object
type Bullet struct {
	Context js.Value
	X       int
	Y       int
	Speed  float64
	SpeedY  float64
	Angle	float64

	
}

// NewBullet returns a new player object
func NewBullet(context js.Value, x int, y int, angle float64, speed float64) *Bullet {
	p := &Bullet{
		Context: context,
		X:       x,
		Y:       y,
		Speed:  15.0+speed,
		Angle: 10.0,

	}
	return p
}


// Draw renders the player on the screen
func (b *Bullet) Draw() {
	b.Context.Call("beginPath")
	b.Context.Call("arc", b.X, b.Y, 2, 0 , 2 * math.Pi)
	b.Context.Call("fill")
	b.Context.Call("closePath")
}

// GetBoundingBox returns the bounding box for a bullet
func (b *Bullet) GetBoundingBox()(*BoundingBox){
	return &BoundingBox{
		X: b.X - 2,
		Y: b.Y - 2,
		Width: 4,
		Height: 4,
	}
}

// Step Handles misc player update functions
func (b *Bullet) Step(){
	angleRad := b.Angle * (math.Pi/180)
	b.X += int(25.0*math.Sin(angleRad))
	b.Y += int(25.0*math.Cos(angleRad))
}
