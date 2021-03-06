package main

import (
	. "github.com/fogleman/pt/pt"
	"github.com/ojrac/opensimplex-go"
)

func main() {
	scene := Scene{}
	material := GlossyMaterial(Color{1, 1, 1}, 1.2, Radians(20))
	noise := opensimplex.New()
	n := 80
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n*2; k++ {
				x := float64(i - n/2)
				y := float64(j - n/2)
				z := float64(k)
				m := 0.15
				w := noise.Eval3(x*m, y*m, z*m)
				w = (w + 0.8) / 1.6
				if w <= 0.2 {
					shape := NewSphere(Vector{x, y, z}, 0.333, material)
					scene.Add(shape)
				}
			}
		}
	}
	light := NewSphere(Vector{100, 0, 50}, 5, LightMaterial(Color{1, 1, 1}, 1, NoAttenuation))
	scene.Add(light)
	camera := LookAt(Vector{0, 0, -20}, Vector{}, Vector{0, 1, 0}, 30)
	IterativeRender("out%03d.png", 1000, &scene, &camera, 2560, 1440, -1, 4, 4)
}
