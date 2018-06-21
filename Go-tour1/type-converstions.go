package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z)
}

//Este no funciona, ya que no puede usar type int as type float64
/*
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f = math.Sqrt((x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z)
}
*/