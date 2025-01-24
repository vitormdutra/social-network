package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e coloca um hash
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificaSenha compara uma senha e um hash e retorna se elas sao iguais
func VerificaSenha(senhaString, senhaHash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString)) == nil
}
