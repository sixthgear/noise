package noise

import "os"
import "fmt"
import "testing"
import "image"
import "image/png"
import "image/color"

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

func TestNoisePNG(t *testing.T) {

	fmt.Println("Testing PNG Noise (output to out.png)")

	width, height := 512, 512

	r := image.Rect(0, 0, width, height)
	i := image.NewRGBA(r)
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			n := Noise2d(float64(x)/float64(width/64), float64(y)/float64(height/64))
			ni := uint8(n * 255)
			// fmt.Println(ni)
			c := color.RGBA{ni, ni, ni, 255}
			i.SetRGBA(x, y, c)
		}
	}
	f, err := os.Create("out.png")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	png.Encode(f, i)

}
