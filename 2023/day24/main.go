package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	
	"gonum.org/v1/gonum/mat"
)

type Particle struct {
	x, y, z int
	vx, vy, vz int
}

func xy_intersection(p1, p2 Particle) (float64, float64, error) {
	x1, y1, vx1, vy1 := p1.x, p1.y, p1.vx, p1.vy
	x2, y2, vx2, vy2 := p2.x, p2.y, p2.vx, p2.vy

	det := -vx1 * vy2 + vy1 * vx2
	if det == 0 {
		return 0.0, 0.0, fmt.Errorf("they don't intersect -- parallel")
	}

	det_t := (x1 - x2) * vy2 - (y1 - y2) * vx2
	det_s := -vx1 * (y1 - y2) + vy1 * (x1 - x2)

	t := float64(det_t) / float64(det)
	s := float64(det_s) / float64(det)

	xt := float64(x1) + float64(vx1) * t
	yt := float64(y1) + float64(vy1) * t

	if t < 0 || s < 0 {
		return 0.0, 0.0, fmt.Errorf("they crossed in the past")
	}

	return xt, yt, nil
}

func intersection(p1, p2, p3 Particle) uint64 {
	x1, y1, z1, vx1, vy1, vz1 := float64(p1.x), float64(p1.y), float64(p1.z), float64(p1.vx), float64(p1.vy), float64(p1.vz)
	x2, y2, z2, vx2, vy2, vz2 := float64(p2.x), float64(p2.y), float64(p2.z), float64(p2.vx), float64(p2.vy), float64(p2.vz)
	x3, y3, z3, vx3, vy3, vz3 := float64(p3.x), float64(p3.y), float64(p3.z), float64(p3.vx), float64(p3.vy), float64(p3.vz)


	// (rk - r1) + (vk - r1) * t = 0
	// (rk - r1) x (vk - r1) = 0 (cross product on both sides)
	// (rk - r2) x (vk - r2) = 0 (cross product on both sides)
	// 
	// (v2 - v1) x rk + (r2 - r1) x vk = r2 x v2 - r1 x v1 (substract both)
	// (v3 - v1) x rk + (r3 - r1) x vk = r3 x v3 - r1 x v1 (substract both)
	// 
	rhs := mat.NewVecDense(6, []float64{
		(y2 * vz2 - vy2 * z2)  - (y1 * vz1 - vy1 * z1),
	       -(x2 * vz2 - vx2 * z2)  + (x1 * vz1 - vx1 * z1),
		(x2 * vy2 - vx2 * y2)  - (x1 * vy1 - vx1 * y1),
		(y3 * vz3 - vy3 * z3)  - (y1 * vz1 - vy1 * z1),
	       -(x3 * vz3 - vx3 * z3)  + (x1 * vz1 - vx1 * z1),
		(x3 * vy3 - vx3 * y3)  - (x1 * vy1 - vx1 * y1),
	})


	A := mat.NewDense(6, 6, []float64{
		           0,   (vz2 - vz1), -(vy2 - vy1),          0,  (z2 - z1), -(y2 - y1),
		-(vz2 - vz1),             0,  (vx2 - vx1), -(z2 - z1),          0,  (x2 - x1),
		 (vy2 - vy1),  -(vx2 - vx1),            0,  (y2 - y1), -(x2 - x1),          0,
		           0,   (vz3 - vz1), -(vy3 - vy1),          0,  (z3 - z1), -(y3 - y1),
		-(vz3 - vz1),             0,  (vx3 - vx1), -(z3 - z1),          0,  (x3 - x1),
		 (vy3 - vy1),  -(vx3 - vx1),            0,  (y3 - y1), -(x3 - x1),          0,
	})

	var x mat.VecDense
	x.SolveVec(A, rhs);

	tol := 1.0e-1
	return uint64(x.AtVec(0) + tol) + uint64(x.AtVec(1) + tol) + uint64(x.AtVec(2) + tol)
}

func point_in_rectangle(x, y, xmin, xmax, ymin, ymax float64) bool {
	if xmin <= x && x <= xmax && ymin <= y && y <= ymax {
		return true
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		panic("The program requires the input file path as argument")
	}
	input := os.Args[1]
	dat, err := os.ReadFile(input)
	if err != nil {
		panic("Input file not found")
	}

	words := strings.Split(strings.Trim(string(dat), "\n"), "\n")
	particles := []Particle{}

	for _, line := range words {
		line = strings.ReplaceAll(line, " ", "")
		line_s := strings.Split(line, "@")
		p := strings.Split(line_s[0], ",")
		x, _ := strconv.Atoi(p[0])
		y, _ := strconv.Atoi(p[1])
		z, _ := strconv.Atoi(p[2])
		v := strings.Split(line_s[1], ",")
		vx, _ := strconv.Atoi(v[0])
		vy, _ := strconv.Atoi(v[1])
		vz, _ := strconv.Atoi(v[2])
		particles = append(particles, Particle{x,y,z,vx,vy,vz})
	}

	solution1 := 0
	for i := 0; i < len(particles)-1; i++ {
		for j := i+1; j < len(particles); j++ {
			x, y, err := xy_intersection(particles[i], particles[j])

			if err == nil {
				xmin, xmax, ymin, ymax := 200000000000000.0, 400000000000000.0, 200000000000000.0, 400000000000000.0
				are_in_rectangle := point_in_rectangle(x, y, xmin, xmax, ymin, ymax)

				if are_in_rectangle {
					solution1++
				}
			}
		}
	}

	solution2 := intersection(particles[0], particles[1], particles[2])

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
