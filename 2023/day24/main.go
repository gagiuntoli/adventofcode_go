package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	
	"gonum.org/v1/gonum/mat"
	//utils "github.com/gagiuntoli/adventofcode_go/utils"
)

type Particle struct {
	x, y, z int
	vx, vy, vz int
}

func xy_intersection(p1, p2 Particle) (float64, float64, error) {
	x1, y1, z1, vx1, vy1, vz1 := p1.x, p1.y, p1.z, p1.vx, p1.vy, p1.vz
	x2, y2, z2, vx2, vy2, vz2 := p2.x, p2.y, p2.z, p2.vx, p2.vy, p2.vz

	det := -vx1 * vy2 + vy1 * vx2
	if det == 0 {
		return 0.0, 0.0, fmt.Errorf("they don't intersect -- parallel")
	}

	det_t := (x1 - x2) * vy2 - (y1 - y2) * vx2
	det_s := -vx1 * (y1 - y2) + vy1 * (x1 - x2)

	det_sz := -vx1 * (z1 - z2) + vz1 * (x1 - x2)
	det_tz := (x1 - x2) * vz2 - (z1 - z2) * vx2
	detz := -vx1 * vz2 + vz1 * vx2

	t := float64(det_t) / float64(det)
	s := float64(det_s) / float64(det)

	xt := float64(x1) + float64(vx1) * t
	yt := float64(y1) + float64(vy1) * t

	if t < 0 || s < 0 {
		return 0.0, 0.0, fmt.Errorf("they crossed in the past")
	}
	fmt.Println("dets", det_t, det_s, det, "dets z", det_tz, det_sz, detz)

	return xt, yt, nil
}

func newton_raphson(p1, p2, p3 Particle, max_iters int) (uint64, float64) {
	x1, y1, z1, vx1, vy1, vz1 := float64(p1.x), float64(p1.y), float64(p1.z), float64(p1.vx), float64(p1.vy), float64(p1.vz)
	x2, y2, z2, vx2, vy2, vz2 := float64(p2.x), float64(p2.y), float64(p2.z), float64(p2.vx), float64(p2.vy), float64(p2.vz)
	x3, y3, z3, vx3, vy3, vz3 := float64(p3.x), float64(p3.y), float64(p3.z), float64(p3.vx), float64(p3.vy), float64(p3.vz)

	x := mat.NewVecDense(9, []float64{0,0,0,0,0,0,0,0,0})

	var result uint64
	var norm float64
	for iter := 0; iter < max_iters; iter++ {
		xk, yk, zk, vxk, vyk, vzk, t1, t2, t3 := x.AtVec(0), x.AtVec(1), x.AtVec(2), x.AtVec(3), x.AtVec(4), x.AtVec(5), x.AtVec(6), x.AtVec(7), x.AtVec(8)

		b := mat.NewVecDense(9, []float64{
			(xk - x1) + (vxk - vx1) * t1,
			(yk - y1) + (vyk - vy1) * t1,
			(zk - z1) + (vzk - vz1) * t1,
			(xk - x2) + (vxk - vx2) * t2,
			(yk - y2) + (vyk - vy2) * t2,
			(zk - z2) + (vzk - vz2) * t2,
			(xk - x3) + (vxk - vx3) * t3,
			(yk - y3) + (vyk - vy3) * t3,
			(zk - z3) + (vzk - vz3) * t3,
		})

		norm = b.Norm(1)
		fmt.Println("norm", norm)

		A := mat.NewDense(9, 9, []float64{
			1, 0, 0, t1,  0,  0, (vxk - vx1),           0,           0,
			0, 1, 0,  0, t1,  0, (vyk - vy1),           0,           0,
			0, 0, 1,  0,  0, t1, (vzk - vz1),           0,           0,
			1, 0, 0, t2,  0,  0,           0, (vxk - vx2),           0,
			0, 1, 0,  0, t2,  0,           0, (vyk - vy2),           0,
			0, 0, 1,  0,  0, t2,           0, (vzk - vz2),           0,
			1, 0, 0, t3,  0,  0,           0,           0, (vxk - vx3),
			0, 1, 0,  0, t3,  0,           0,           0, (vyk - vy3),
			0, 0, 1,  0,  0, t3,           0,           0, (vzk - vz3),
		})

		var dy mat.VecDense
		dy.SolveVec(A, b);
		//fmt.Println("b = ", b)
		fmt.Println("A = ", A)
		fmt.Println("solution", dy)

		x.SubVec(x, dy.SliceVec(0, dy.Len()))
	}

	return result, norm
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
				//xmin, xmax, ymin, ymax := 7.0, 27.0, 7.0, 27.0
				xmin, xmax, ymin, ymax := 200000000000000.0, 400000000000000.0, 200000000000000.0, 400000000000000.0
				are_in_rectangle := point_in_rectangle(x, y, xmin, xmax, ymin, ymax)

				if are_in_rectangle {
					solution1++
				}
			}
		}
	}

	rock := Particle{24, 13, 10, -3, 1, 2}
	for i := 0; i < len(particles); i++ {
		x, y, _ := xy_intersection(particles[i], rock)
		fmt.Println(i, "Collision at", x, y)

	}

	res, norm := newton_raphson(particles[0], particles[1], particles[2], 10)
	fmt.Println(res, norm)


	solution2 := 0
	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
