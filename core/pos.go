package core

// NoPos represents an invalid position.
const NoPos Pos = 0

// Pos represents a position in the file set.
type Pos int

// IsValid returns true if the position is valid.
func (p Pos) IsValid() bool {
	return p != NoPos
}
