package interfaces

//go:generate mockgen -source=bcrypt.go -destination=../mocks/bcrypt.go
type BCrypt interface {
	CompareAndHash([]byte, []byte) error
	GenerateFromPassword([]byte, int) ([]byte, error)
}
