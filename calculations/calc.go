package calc

import "math"

/*Перевод градусов в радианы*/
func DegToRadians(D float64) float64 {
	radians := D * (math.Pi / 180)
	return radians
}

/*Перевод из радиан в градусы*/
func RadiansToDegrees(R float64) float64 {
	degrees := R * (180 / math.Pi)
	return degrees
}

/* Вычисление дистанции до объекта, исходя из его
линейного размера и углового размера.*/
func Lcalc(D, a float64) float64 {
	L := D / (2 * RadiansToDegrees(math.Tan(DegToRadians(a/2))))
	return L
}

/*Вычисление линейного размера, исходя из его
углово размера и дистанции до него.*/
func Dcalc(L, a float64) float64 {
	D := 2 * L * RadiansToDegrees(math.Tan(DegToRadians(a/2)))
	return D
}

/*Вычисление углового размера, исходя из его
линейного размера и дистанции до него*/
func Acalc(D, L float64) float64 {
	a := 2 * RadiansToDegrees(math.Atan(D/(2*L)))
	return a
}
