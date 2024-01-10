package bcyrpt

import (
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
	bcryptLib "golang.org/x/crypto/bcrypt"
)

type bcrypt struct{}

var _ interfaces.BCrypt = (*bcrypt)(nil)

func New() *bcrypt {
	return &bcrypt{}
}

func (b *bcrypt) CompareAndHash(hashedPassword []byte, password []byte) error {
	return bcryptLib.CompareHashAndPassword(hashedPassword, password)
}

func (b *bcrypt) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcryptLib.GenerateFromPassword(password, cost)
}
