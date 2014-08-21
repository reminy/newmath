// Package newmath is an example library package newmath 
package newmath

// Sqrt returns approximation to square root of x.
func Sqrt(x float64) float64	{
	z := 1.0
	for i := 0; i < 1000; i++	{

		z -= (z*z - x) / (2 * z)
	}
	return z
}

