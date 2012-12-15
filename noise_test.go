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

	fmt.Println("Testing PNG Noise (outputing frames to ./out/out-NNN.png)")

	if err := os.MkdirAll("out", os.FileMode(0777)); err != nil {
		t.Fatal(err.Error())
	}

	width, height, depth := 256, 256, 128
	r := image.Rect(0, 0, width, height)
	i := image.NewRGBA(r)

	for z := 0; z <= depth; z++ {
		for y := 0; y <= height; y++ {
			for x := 0; x <= width; x++ {
				xx := float64(x)
				yy := float64(y)
				zz := float64(z)
				n := OctaveNoise3d(xx, yy, zz, 4, 0.6, 1.0/16)
				n = (n + 1.0) * 0.5
				ni := uint8(n * 255)
				c := color.RGBA{ni, ni, ni, 255}
				i.SetRGBA(x, y, c)
			}
		}

		f, err := os.Create(fmt.Sprintf("out/out-%03d.png", z))
		if err != nil {
			t.Fatal(err.Error())
		}
		png.Encode(f, i)
		f.Close()

	}

}
