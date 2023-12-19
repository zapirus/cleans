package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"

	"github.com/google/uuid"
)

func (u *UseCase) generateCode() int {
	val := rand.Intn(921343) + 100000
	return val
}

func (u *UseCase) generateGuid() uuid.UUID {
	newGuid := uuid.New()
	return newGuid
}

func (u *UseCase) generateHashPass(pass string) string {
	hash := sha256.New()
	hash.Write([]byte(pass))
	hashValue := hash.Sum(nil)
	hashString := hex.EncodeToString(hashValue)
	return hashString
}
