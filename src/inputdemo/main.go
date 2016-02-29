package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println("Circl Area calculation")
  fmt.Print("Enter a radious value: ")
  var radius float64
  fmt.Scanf("%f", &radius)

  area := math.Pi * math.Pow(radius, 2)
  fmt.Printf("Circle area with radious %.2f = %.2f \n", radius, area)
}

