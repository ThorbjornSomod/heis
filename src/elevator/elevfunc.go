package elevator

import (
	"math/rand"
)



func RandSeq(n int) string{ //Function generating a random string of length n.

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b:= make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return(string(b))
}

