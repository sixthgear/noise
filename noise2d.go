package noise

import "math"

func Dot2(g Grad3, x, y float64) float64 {
	return g.x*x + g.y*y
}

func Noise2d(xin, yin float64) float64 {

	var (
		n, x, y, t [3]float64
		gi         [3]int
		i, j       [2]int
	)

	s := (xin + yin) * F2
	i[0] = int(math.Floor(xin + s))
	j[0] = int(math.Floor(yin + s))
	tt := float64(i[0]+j[0]) * G2
	x[0] = xin - (float64(i[0]) - tt)
	y[0] = yin - (float64(j[0]) - tt)

	if x[0] > y[0] {
		i[1], j[1] = 1, 0
	} else {
		i[1], j[1] = 0, 1
	}

	x[1] = x[0] - float64(i[1]) + G2
	y[1] = y[0] - float64(j[1]) + G2
	x[2] = x[0] - 1.0 + G2*2
	y[2] = y[0] - 1.0 + G2*2

	ii := i[0] & 255
	jj := j[0] & 255
	gi[0] = permMod12[ii+perm[jj]]
	gi[1] = permMod12[ii+i[1]+perm[jj+j[1]]]
	gi[2] = permMod12[ii+1+perm[jj+1]]

	for i := 0; i < len(n); i++ {
		t[i] = 0.5 - x[i]*x[i] - y[i]*y[i]
		if t[i] < 0 {
			n[i] = 0.0
		} else {
			n[i] = math.Pow(t[i], 4) * Dot2(grad3Table[gi[i]], x[i], y[i])
		}
	}

	return 70.0 * (n[0] + n[1] + n[2])
}
