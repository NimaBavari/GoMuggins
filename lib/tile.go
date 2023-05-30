package muggins

type Tile struct {
	left  int
	right int
}

func (t Tile) Face() int {
	return t.left + t.right
}

func (t Tile) IsDouble() bool {
	return t.left == t.right
}

func (t Tile) IsSame(o Tile) bool {
	return (t.left == o.left && t.right == o.right) || (t.left == o.right && t.right == o.left)
}

func (t Tile) IsPlayable(e int) bool {
	return t.left == e || t.right == e
}
