package color

var Banner = Color{Term: "\033[35m", Hex: "#000000"}
var Primary = Color{Term: "\033[36m", Hex: "#00FFFF"}   // cyan
var Secondary = Color{Term: "\033[34m", Hex: "#0000FF"} // blue

type Color struct {
	Term string
	Hex  string
}

func Term(s string, c Color) string {
	return c.Term + s + "\033[0m"
}
