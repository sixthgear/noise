package noise

import "fmt"
import "testing"

func TestNoise2d(t *testing.T) {
	fmt.Println("Testing 2D Noise")
	for j := 0; j <= 10; j++ {
		for i := 0; i <= 10; i++ {
			fmt.Printf("%6.2f ", Noise2d(float64(i)/10, float64(j)/10))
		}
		fmt.Println()
	}
	fmt.Println()
}

func TestNoise3d(t *testing.T) {
	fmt.Println("Testing 3D Noise")
	for k := 0; k <= 10; k++ {
		for j := 0; j <= 10; j++ {
			for i := 0; i <= 10; i++ {
				fmt.Printf("%6.2f ", Noise3d(float64(i)/10, float64(j)/10, float64(k)/10))
			}
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Println()
}
