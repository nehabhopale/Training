package Calc
import "math"

func Add(num1,num2 int)int{
	return num1+num2
}
func Sub(num1,num2 int)int{
 	return num1-num2
 }
func Mul(num1,num2 int)int{
	 return num1*num2
}
func SquareRoot(f float64) float64 {
    return math.Pow(f, 0.5)
}
func Power(num1 float64,pow float64) float64{
	return math.Pow(num1,pow)
}