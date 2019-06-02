package objects

// Object represents and object in the game
type Object interface{
	Draw()
	Step()
}