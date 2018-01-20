package main
import (
	"github.com/hoisie/web"
)

func hello2(val string) string { return "hello " + val }

func main() {
	web.Get("/(.*)", hello2)
	web.Run("0.0.0.0:9999")
}