package goarea
//foi alterado apenas o nome do pacote para goarea e o resto foi copiado do código fonte
import (
	"math"
)

//Pi constante para cálculo da área do círculo
const Pi = 3.1416

//Circulo área do círculo
func Circulo(r float64) float64 {
	return Pi * math.Pow(r, 2)
}

//Retangulo área do retângulo
func Retangulo(b, a float64) float64 {
	return b * a
}

//Função _TriEqui não é pública logo não precisa de comentário
//Funções começando por _ ou letra minúscula não são visíveis fora do pacote (privadas)
//Logo _TriEqui não irá funcionar no main
func _TriEqui(b, a float64) float64 {
	return (b * a) / 2
}
