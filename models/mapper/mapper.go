package mapper

import "github.com/rafiulgits/go-automapper"

func Map(src, dst interface{}, args ...interface{}) {
	automapper.Map(src, dst, args...)
}
