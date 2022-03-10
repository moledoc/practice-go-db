package data

import (
	"fmt"
	"math/rand"
)

type Data struct {
	Id   int
	Name string
}

func Gen(n int) []Data {
	data := make([]Data, n)
	for i := 0; i < n; i++ {
		rnd := rand.Intn(1000)
		data[i] = Data{Id: rnd, Name: fmt.Sprintf("test%v", rnd)}
	}
	return data
}
