// O pacote Dreambridge_utils contém funções úteis para várias coisas diversas
package stringutils

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// CriptografaString - Criptografa uma senha
func CriptografaString(senha string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), cost)

	fmt.Println("Hash: " + string(bytes))
	return string(bytes), err
}

// checkStringHash - Verifica se a criptografia corresponde à string passada
func checkStringHash(str, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))

	if err != nil {
		log.Println("checkStringHash - Erro ao validar string.")
		return false, err
	}

	return true, nil
}