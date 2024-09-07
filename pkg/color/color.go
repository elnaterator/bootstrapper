package color

type Color struct {
	value string
}

var reset = Color{value: "\033[0m"}
var Red = Color{value: "\033[31m"}
var Green = Color{value: "\033[32m"}
var Yellow = Color{value: "\033[33m"}
var Blue = Color{value: "\033[34m"}
var Magenta = Color{value: "\033[35m"}
var Cyan = Color{value: "\033[36m"}
var Gray = Color{value: "\033[37m"}
var White = Color{value: "\033[97m"}

func Wrap(s string, c Color) string {
	return c.value + s + reset.value
}
