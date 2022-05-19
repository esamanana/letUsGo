package main

import (
	"fmt"
	"math"
)

/*
 Define a Cube struct
*/
type Cube struct {
	a float64
}

/*
* Method to calculate the lateral surface area of a given Cube
* Return type: float64
 */
func (cube Cube) getVolume() float64 {
	return math.Pow(cube.a, 3)
}

/*
* Method to calculate the lateral surface area of a given Cube
* Return type: float64
 */
func (cube Cube) getLateralSurfaceArea() float64 {
	return 4.0 * math.Pow(cube.a, 2)
}

/*
* Method to calculate the total surface area of a given Cube
* return type: float64
 */
func (cube Cube) getTotalSurfaceArea() float64 {
	return 6.0 * math.Pow(cube.a, 2)
}

/*
 Define a Cuboid struct
*/
type Cuboid struct {
	l, b, h float64
}

/*
* Method to calculate the lateral surface area of a given Cuboid
* Return type: float64
 */
func (cuboid Cuboid) getVolume() float64 {
	return cuboid.l * cuboid.b * cuboid.h
}

/*
* Method to calculate the lateral surface area of a given Cuboid
* Return type: float64
 */
func (cuboid Cuboid) getLateralSurfaceArea() float64 {
	return 2.0 * cuboid.h * (cuboid.l + cuboid.b)
}

/*
* Method to calculate the total surface area of a given Cuboid
* return type: float64
 */
func (cuboid Cuboid) getTotalSurfaceArea() float64 {
	return 2.0 * (cuboid.l*cuboid.b + cuboid.b*cuboid.h + cuboid.h*cuboid.l)
}

/*
 Define a Sphere struct
*/
type Sphere struct {
	radius float64
}

/*
* Method to calculate the lateral surface area of a given Sphere
* Return type: float64
 */
func (sphere Sphere) getVolume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(sphere.radius, 3)
}

/*
* Method to calculate the lateral surface area of a given Sphere
* Return type: float64
 */
func (sphere Sphere) getLateralSurfaceArea() float64 {
	return 4.0 * math.Pi * math.Pow(sphere.radius, 2)
}

/*
* Method to calculate the total surface area of a given Hemisphere
* return type: float64
 */
func (sphere Sphere) getTotalSurfaceArea() float64 {
	return 4.0 * math.Pi * math.Pow(sphere.radius, 2)
}

/*
 Define a Hemisphere struct
*/
type Hemisphere struct {
	radius float64
}

/*
* Method to calculate the lateral surface area of a given Hemisphere
* Return type: float64
 */
func (hemiSphere Hemisphere) getVolume() float64 {
	return (2.0 / 3.0) * math.Pi * math.Pow(hemiSphere.radius, 3)
}

/*
* Method to calculate the lateral surface area of a given Hemisphere
* Return type: float64
 */
func (hemiSphere Hemisphere) getLateralSurfaceArea() float64 {
	return 2.0 * math.Pi * math.Pow(hemiSphere.radius, 2)
}

/*
* Method to calculate the total surface area of a given Hemishere
* return type: float64
 */
func (hemiSphere Hemisphere) getTotalSurfaceArea() float64 {
	return 3.0 * math.Pi * math.Pow(hemiSphere.radius, 2)
}

/*
 Define a Cylinder struct
*/
type Cylinder struct {
	radius float64
	height float64
}

/*
* Method to calculate the volume of a given Cylinder
* Return type: float64
 */
func (cylinder Cylinder) getVolume() float64 {
	return math.Pi * math.Pow(cylinder.radius, 2) * cylinder.height
}

/*
* Method to calculate the lateral surface area of a given Cylinder
* Return type: float64
 */
func (cylinder Cylinder) getLateralSurfaceArea() float64 {
	return 2.0 * math.Pi * cylinder.radius * cylinder.height
}

/*
* Method to calculate the total surface area of a given Cylinder
* return type: float64
 */
func (cylinder Cylinder) getTotalSurfaceArea() float64 {
	return (2.0*math.Pi*cylinder.radius*cylinder.height + 2.0*math.Pi*math.Pow(cylinder.radius, 2))
}

/*
 Define a Cone struct
*/
type Cone struct {
	radius float64
	height float64
}

/*
* Method to calculate the volume  of a given Cone
* Return type: float64
 */
func (cone Cone) getVolume() float64 {
	return (1.0 / 3.0) * math.Pi * math.Pow(cone.radius, 2) * cone.height
}

/*
* Method to calculate the lateral surface area of a given Cone
* Return type: float64
 */
func (cone Cone) getLateralSurfaceArea() float64 {
	length := math.Sqrt(math.Pow(cone.height, 2) + math.Pow(cone.radius, 2))
	return math.Pi * cone.radius * length
}

/*
* Method to calculate the total surface area of a given Cone
* return type: float64
 */
func (cone Cone) getTotalSurfaceArea() float64 {
	length := math.Sqrt(math.Pow(cone.height, 2) + math.Pow(cone.radius, 2))
	return math.Pi * cone.radius * (cone.radius + length)
}

func main() {
	/*
		Use the below url to confirm the returned values
		   https://www.calculatorsoup.com/calculators/geometry-solids/
	*/
	cube := Cube{20}
	fmt.Println("\nThe Volume of the Cube is: ", cube.getVolume(), "cubic meters")
	fmt.Println("The lateral surface area of the cube is: ", cube.getLateralSurfaceArea(), "square meters")
	fmt.Println("The total surface area of the cube is: ", cube.getTotalSurfaceArea(), "square meters")

	cuboid := Cuboid{4, 2.5, 1.5}
	fmt.Println("\nThe Volume of the Cuboid is: ", cuboid.getVolume(), "cubic meters")
	fmt.Println("The lateral surface area of the cuboid is: ", cuboid.getLateralSurfaceArea(), "square meters")
	fmt.Println("The total surface area of the cuboid is: ", cuboid.getTotalSurfaceArea(), "square meters")

	sphere := Sphere{3.5}
	fmt.Println("\nThe Volume of the sphere is: ", sphere.getVolume(), "cubic meters")
	fmt.Println("The lateral surface area of the sphere is: ", sphere.getLateralSurfaceArea(), "square meters")
	fmt.Println("The total surface area of the sphere is: ", sphere.getTotalSurfaceArea(), "square meters")

	hemiSphere := Hemisphere{3}
	fmt.Println("\nThe Volume of the hemiSphere  is: ", hemiSphere.getVolume(), "cubic meters")
	fmt.Println("The lateral surface area of the hemiSphere is: ", hemiSphere.getLateralSurfaceArea(), "square meters")
	fmt.Println("The total surface area of the hemiSphere is: ", hemiSphere.getTotalSurfaceArea(), "square meters")

	cylinder := Cylinder{5, 10}
	fmt.Println("\nThe Volume of the cylinder  is: ", cylinder.getVolume(), "cubic meters")
	fmt.Println("The lateral surface area of the cylinder is: ", cylinder.getLateralSurfaceArea(), "square meters")
	fmt.Println("The total surface area of the cylinder is: ", cylinder.getTotalSurfaceArea(), "square meters")

	cone := Cone{5, 10}
	fmt.Println("\nThe Volume of the cone   is: ", cone.getVolume(), "cubic meters")
	fmt.Println("The lateral surface area of the cone  is: ", cone.getLateralSurfaceArea(), "square meters")
	fmt.Println("The total surface area of the cone  is: ", cone.getTotalSurfaceArea(), "square meters")
}
