package noise

import "math"

func Dot3(g Grad3, x, y, z float64) float64 {
	return g.x*x + g.y*y + g.z*z
}

func Noise3d(xin, yin, zin float64) float64 {

	var (
		n, x, y, z, t [4]float64
		gi            [4]int
		i, j, k       [3]int
	)

	s := (xin + yin + zin) * F3
	i[0] = int(math.Floor(xin + s))
	j[0] = int(math.Floor(yin + s))
	k[0] = int(math.Floor(zin + s))
	tt := float64(i[0]+j[0]+k[0]) * G3

	x[0] = xin - (float64(i[0]) - tt)
	y[0] = yin - (float64(j[0]) - tt)
	z[0] = zin - (float64(k[0]) - tt)

	if x[0] >= y[0] {
		if y[0] >= z[0] {
			i[1], j[1], k[1], i[2], j[2], k[2] = 1, 0, 0, 1, 1, 0
		} else if x[0] >= z[0] {
			i[1], j[1], k[1], i[2], j[2], k[2] = 1, 0, 0, 1, 0, 1
		} else {
			i[1], j[1], k[1], i[2], j[2], k[2] = 0, 0, 1, 1, 0, 1
		}
	} else {
		if y[0] < z[0] {
			i[1], j[1], k[1], i[2], j[2], k[2] = 0, 0, 1, 0, 1, 1
		} else if x[0] < z[0] {
			i[1], j[1], k[1], i[2], j[2], k[2] = 0, 1, 0, 0, 1, 1
		} else {
			i[1], j[1], k[1], i[2], j[2], k[2] = 0, 1, 0, 1, 1, 0
		}
	}

	x[1] = x[0] - float64(i[1]) + G3
	y[1] = y[0] - float64(j[1]) + G3
	z[1] = z[0] - float64(k[1]) + G3
	x[2] = x[0] - float64(i[2]) + 2.0*G3
	y[2] = y[0] - float64(j[2]) + 2.0*G3
	z[2] = z[0] - float64(k[2]) + 2.0*G3
	x[3] = x[0] - 1.0 + 3.0*G3
	y[3] = y[0] - 1.0 + 3.0*G3
	z[3] = z[0] - 1.0 + 3.0*G3

	ii := i[0] & 255
	jj := j[0] & 255
	kk := k[0] & 255
	gi[0] = permMod12[ii+perm[jj+perm[kk]]]
	gi[1] = permMod12[ii+i[1]+perm[jj+j[1]+perm[kk+k[1]]]]
	gi[2] = permMod12[ii+i[2]+perm[jj+j[2]+perm[kk+k[2]]]]
	gi[3] = permMod12[ii+1+perm[jj+1+perm[kk+1]]]

	for i := 0; i < len(n); i++ {
		t[i] = 0.6 - x[i]*x[i] - y[i]*y[i] - z[i]*z[i]
		if t[i] < 0 {
			n[i] = 0.0
		} else {
			n[i] = math.Pow(t[i], 4) * Dot3(grad3Table[gi[i]], x[i], y[i], z[i])
		}
	}

	return 32.0 * (n[0] + n[1] + n[2] + n[3])
}

func OctaveNoise3d(xin, yin, zin float64, octaves int, p, scale float64) float64 {

	total := 0.0
	freq := scale
	amp := 1.0
	maxAmp := 0.0

	for i := 0; i < octaves; i++ {
		total += Noise3d(xin*freq, yin*freq, zin*freq) * amp
		freq *= 2.0
		maxAmp += amp
		amp *= p
	}

	return total / maxAmp
}
