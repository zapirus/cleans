package usecase

import (
	"math/rand"
)

func (u *UseCase) GenerateCode() int {
	val := rand.Intn(921343) + 100000
	return val
}
