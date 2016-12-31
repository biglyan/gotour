package main

import (
    "fmt"
)

// error type
type ErrNegativeSqrt float64

// error function for the error type
func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf(
        "Can't find the Sqrt of a negative number: %v", float64(e))
}

func Sqrt(f float64) (float64, error) {

    // return error if number is negative
    if f < 0 {
        return 0, ErrNegativeSqrt(f)

    // otherwise find an approximation for the square root
    } else {
        z := float64(1)
        for i := 0; i < 1000; i++ {
            z = z - (z * z - f)/(2 * z)
        }
        return z, nil
    }
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}