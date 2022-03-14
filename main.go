package goarea
import "math"
//Pi é uma proporção numérica definida pela relação dentr
// o périmetro de uma circunferência e seu diâmetro
const Pi = 3.1416
//Circ é responsável por calcular a área da circunferência
func Circ(raio float64) {
  return math.Pow(raio, 2) * Pi
}

//React é responsável por calcular a área de um retangulo
func React(base, altura float64) float64 {
  return base * altura
}
// Não é visível!
func _TrianguloEoq(base, altura float64) float64 {
  return (base * altura) / 2
}
