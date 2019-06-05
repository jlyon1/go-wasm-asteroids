package objects

import (
	"syscall/js"
	"math"
)

// Asteroid representsthe main asteroid object
type Asteroid struct {
	Context js.Value
	X       int
	Y       int
	Angle	float64
	Size float64
	Speed float64
	
}

// NewAsteroid returns a new player object
func NewAsteroid(context js.Value, x int, y int, angle float64, size float64, speed float64) *Asteroid {
	p := &Asteroid{
		Context: context,
		X:       x,
		Y:       y,
		Angle: 10.0,
		Size: size,
		Speed: speed,

	}
	return p
}


// Draw renders the player on the screen
func (b *Asteroid) Draw() {
	b.Context.Call("beginPath")
	b.Context.Set("fillStyle", "red")
	b.Context.Call("arc", b.X, b.Y, float64(b.Size), 0 , 2 * math.Pi)
	b.Context.Call("fill")
	b.Context.Call("closePath")
	b.Context.Set("fillStyle", "white")

}

// Step Handles misc player update functions
func (b *Asteroid) Step(){
	angleRad := b.Angle * (math.Pi/180.0)
	b.X += int(b.Speed*math.Sin(angleRad))
	b.Y += int(b.Speed*math.Cos(angleRad))
}

// GetBoundingBox returns the bounding box for a bullet
func (b *Asteroid) GetBoundingBox()(*BoundingBox){
	return &BoundingBox{
		X: b.X - int(b.Size),
		Y: b.Y - int(b.Size),
		Width: 2*int(b.Size),
		Height: 2*int(b.Size),
	}
}