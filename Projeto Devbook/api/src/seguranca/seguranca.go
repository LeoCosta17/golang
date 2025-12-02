package seguranca

import "golang.org/x/crypto/bcrypt"

// Recebe uma string (senha) e coloca um Hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost) // Recebe 2 parametros: a senha a receber o Hash
} // e o "custo" do Hash - que informa o nível de Hash que será feito
// quanto maior, mais dificil fica a conversão (mais seguro)

// Compara uma senha com hash e retorna se ambos são iguais
func VerificarSenha(senha_criptografada string, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(senha_criptografada), []byte(senha))
}
