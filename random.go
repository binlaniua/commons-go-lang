package commons_go_lang

import "math/rand"

//
//
// RandomBoolean return a boolean
//
//
func RandomBoolean() bool {
	if r := rand.Intn(1); r == 0 {
		return false
	}
	return true
}
