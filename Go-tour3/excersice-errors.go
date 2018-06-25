package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number ", float64(e))
}

func Sqrt(x float64) float64 {
z:=float64(1)
for i:=0;i<10;i++{
}
return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

