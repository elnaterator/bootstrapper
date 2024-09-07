package color

type Color struct {
	Term string
	Hex  string
}

var Reset = Color{Term: "\033[0m", Hex: "#000000"}
var Red = Color{Term: "\033[31m", Hex: "#FF0000"}
var Green = Color{Term: "\033[32m", Hex: "#00FF00"}
var Yellow = Color{Term: "\033[33m", Hex: "#FFFF00"}
var Blue = Color{Term: "\033[34m", Hex: "#0000FF"}
var Magenta = Color{Term: "\033[35m", Hex: "#FF00FF"}
var Cyan = Color{Term: "\033[36m", Hex: "#00FFFF"}
var Gray = Color{Term: "\033[37m", Hex: "#808080"}
var White = Color{Term: "\033[97m", Hex: "#FFFFFF"}

func Term(s string, c Color) string {
	return c.Term + s + Reset.Term
}
