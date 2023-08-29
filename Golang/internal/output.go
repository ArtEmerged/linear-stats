package internal

import "fmt"

const (
	LRL = "Linear Regression Line: "
	PCC = "Pearson Correlation Coefficient: "
)

func Output() {
	fmt.Printf("%s%s\n", LRL, Set.LinearRegression)
	fmt.Printf("%s%.10f\n", PCC, Set.PCC)
}
