package Bridge

import "testing"

func TestCircle_Draw(t *testing.T) {
	red := Circle{}
	red.Constructor(2, 5, 3, &RedCircle{})
	red.Draw()

	yellow := Circle{}
	yellow.Constructor(6, 8, 10, &YellowCircle{})
	yellow.Draw()
}
