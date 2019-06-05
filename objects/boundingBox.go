package objects

// BoundingBox represents a bounding box around an object
type BoundingBox struct{
	X int
	Y int
	Width int
	Height int
}

// TL returns the top left corner
func (b *BoundingBox) TL() Point{
	return Point{X: b.X, Y: b.Y}
}

// BR returns the bottom right corner
func (b *BoundingBox) BR() Point{
	return Point{X: b.X + b.Width, Y: b.Y + b.Height}
}
// Intersects checks if the bounding box is intersecting with the other bouinding box
func (b *BoundingBox) Intersects (b2 BoundingBox) bool{
	if (b.TL().X > b2.BR().X || b.BR().X > b2.TL().Y){
		return false
	}
	if (b.TL().Y < b2.BR().Y || b.BR().Y < b2.TL().Y){
		return false
	}
	return true
}