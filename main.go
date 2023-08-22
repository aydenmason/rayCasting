//TODO dont short hand type
package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type sphere struct { 
	radius float32
	color rl.Color //this needs to be of type raylib.Color.RGBA 
						//then we pass the RGBA values of it into the function from the object declaration
	center [3]float32
}

func main() {

	var sphere1 = new(sphere)
	sphere1.radius = 1
	sphere1.color = rl.Color{255,0,0,255} //
	sphere1.center = [3]float32{0,-1,3}

	var sphere2 = new(sphere)
	sphere2.radius = 1
	sphere2.color = rl.Color{0,255,0,255} //
	sphere2.center = [3]float32{-2,0,4}
	
	var sphere3 = new(sphere)
	sphere3.radius = 1
	sphere3.color = rl.Color{0,0,255,255} //
	sphere3.center = [3]float32{2,-0,4}

	//var viewport_size = 1
	//var projection_plane_z = 1

	var scene = []sphere{*sphere1,*sphere2,*sphere3}
	
	rl.InitWindow(1352, 878, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	origin := [3]float32{0, 0, 0}
	//i think im missing something here for the projection plane... i was fixed it lol.
	//projection plane_z is = 1 but im still not really sure why (i just have it hard coded in CanvasToViewport)

	//things should be getting graphed on the screen
	//some math is wrong for my converesions across coordinate systems
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for x := -960; x < 960; x++{
			for y := -540 ; y < 540; y++{
				D := CanvasToViewPort(x, y)
				color := TraceRay(origin, D, 1, 9999999, scene)
				rl.DrawPixel(int32(x),int32(y), color)
			}
		}
		//draw_triangle(0,0,0,50,50,50)

		rl.EndDrawing()
	}
}
func CanvasToViewPort(x int, y int)[3]float32{
	return [3]float32{float32(x)*float32(1)/float32(1352),float32(y)*float32(1)/float32(878),1}
}

func  IntersectRaySphere(O [3]float32, D [3]float32, shape sphere)(float32,float32){
	//TODO, refactor so we  use /vec_dotproduct, 
	//TODO stop using implicit typing!
	r := shape.radius
	CO := vec_subtract(O, shape.center)
	//a := (D[0]*D[0] + D[1]* D[1]) 
	a := vec_dotproduct(D, D)
	//b := 2* (CO[0]*D[0] + CO[1]*D[1])
	b := 2 * vec_dotproduct(CO, D)
	//c := (CO[0]*CO[0] + CO[1]*CO[1]) - r*r
	c := vec_dotproduct(CO, CO) - (r * r)

	disc := (b * b) - (4 * a* c)
	if disc < 0 {
		return 99999999,99999999
	}
	t1 := (-1*b + float32(math.Sqrt(float64(disc)))) / (2*a)
	t2 := (-1*b - float32(math.Sqrt(float64(disc)))) / (2*a)
	return t1,t2

}

func vec_subtract(vec1 [3]float32, vec2 [3]float32)[3]float32{
	return [3]float32{vec1[0]-vec2[0],vec1[1]-vec2[1],vec1[2]-vec2[2]}
}

func vec_dotproduct(vec1 [3]float32, vec2 [3]float32)float32{
	return (vec1[0]*vec2[0] + vec1[1]*vec2[1] + vec1[2]*vec2[2])
}

func TraceRay(O [3]float32, D [3]float32, t_min float32, t_max float32, scene []sphere)rl.Color{
	//something is fucked up here with the logic for over writing the closest sphere. 
	// hard 
	var closest_t float32 = 99999999
	var closest_sphere sphere
	closest_sphere.color = rl.Color{25,25,25,25}
	//closest_sphere.color = rl.Color{255,255,0,255}
	//closest_sphere.center = [3]float32{-2,0,4}
	//closest_sphere.radius = 1

	for i := 0; i < len(scene); i++{
		t1,t2 := IntersectRaySphere(O, D, scene[i])
		if t1 < t_max && t1 > t_min && t1 < closest_t{x
			closest_t = t1
			closest_sphere = scene[i]
		}
		if t2 < t_max && t2 > t_min && t2 < closest_t{
			closest_t = t2 
			closest_sphere = scene[i] 
		}
	}

	if cloeset_sphere.color == r1.Color{25,25,25,25}
	
	return closest_sphere.color	
}