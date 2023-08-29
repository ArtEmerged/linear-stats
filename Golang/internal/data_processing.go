package internal

import (
	"fmt"
	"math"
)

var sumX, sumY, sumXY, sumXPow2, sumYPow2, n float64

type Analytical_Set struct {
	Dx               float64
	Dy               float64
	COVxy            float64
	PCC              float64
	A                float64
	B                float64
	LinearRegression string
}

var Set Analytical_Set

func Dataprocessing() {
	n = float64(len(Data))
	for i, y := range Data {
		x := float64(i)
		sumY += y
		sumX += x
		sumXY += (x * y)
		sumXPow2 += math.Pow(x, 2)
		sumYPow2 += math.Pow(y, 2)
	}
	Set.Dx = sumXPow2 - math.Pow(sumX, 2)/n
	Set.Dy = sumYPow2 - math.Pow(sumY, 2)/n
	Set.COVxy = sumXY - sumX*sumY/n
	Set.A = Set.COVxy / Set.Dx
	Set.B = (sumY - Set.A*sumX) / n

	Set.PCC = Set.COVxy / math.Sqrt(Set.Dx*Set.Dy)
	Set.LinearRegression = fmt.Sprintf("y = %.6fx + %.6f", Set.A, Set.B)
}
