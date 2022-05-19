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

var (
	length, breadth, height, radius float64
	choice, counter                 int
)

func main() {
	/*
	   Simulate a while loop using a for loop
	*/
	for counter = 1; ; counter++ {

		fmt.Println("\nThis program calculate the volume, lateral and total surface area for 3D shapes ")
		fmt.Println("Enter 1 for a Cube")
		fmt.Println("Enter 2 for a Cuboid")
		fmt.Println("Enter 3 for a Sphere")
		fmt.Println("Enter 4 for a Hemisphere")
		fmt.Println("Enter 5 for a Cylinder")
		fmt.Println("Enter 6 for a Cone")
		fmt.Printf("Enter 999 to quit\n\n")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		fmt.Printf("\n\n")

		switch choice {
		case 1:
			{
				fmt.Println("Calculating the volume, lateral and total surface area of a Cube:")
				fmt.Print("Enter the side length: ")
				fmt.Scan(&length)
				cube := Cube{length}
				fmt.Printf("\nVolume: %v\n", cube.getVolume())
				fmt.Printf("Lateral surface area: %v\n", cube.getLateralSurfaceArea())
				fmt.Printf("Total surface area: %v\n", cube.getTotalSurfaceArea())

			}
		case 2:
			{
				fmt.Println("Calculating the volume, lateral and total surface area of a Cuboid:")
				fmt.Print("Enter the length: ")
				fmt.Scan(&length)
				fmt.Print("Enter the breadth: ")
				fmt.Scan(&breadth)
				fmt.Print("Enter the height: ")
				fmt.Scan(&height)

				cuboid := Cuboid{length, breadth, height}

				fmt.Println("\nVolume: ", cuboid.getVolume())
				fmt.Println("Lateral surface area: ", cuboid.getLateralSurfaceArea())
				fmt.Println("Total surface area: ", cuboid.getTotalSurfaceArea())

			}
		case 3:
			{
				fmt.Println("Calculating the volume, lateral and total surface area of a Sphere:")
				fmt.Print("Enter the radius: ")
				fmt.Scan(&radius)
				sphere := Sphere{radius}
				fmt.Println("\nVolume: ", sphere.getVolume())
				fmt.Println("Lateral surface area: ", sphere.getLateralSurfaceArea())
				fmt.Println("Total surface area: ", sphere.getTotalSurfaceArea())
			}
		case 4:
			{
				fmt.Println("Calculating the volume, lateral and total surface area of a Hemisphere:")
				fmt.Print("Enter the radius: ")
				fmt.Scan(&radius)
				hemiSphere := Hemisphere{radius}
				fmt.Println("\nVolume: ", hemiSphere.getVolume())
				fmt.Println("Lateral surface area: ", hemiSphere.getLateralSurfaceArea())
				fmt.Println("Total surface area: ", hemiSphere.getTotalSurfaceArea())
			}
		case 5:
			{
				fmt.Println("Calculating the volume, lateral and total surface area of a Cylinder:")
				fmt.Print("Enter the radius: ")
				fmt.Scan(&radius)
				fmt.Print("Enter the height: ")
				fmt.Scan(&height)
				cylinder := Cylinder{radius, height}
				fmt.Println("\nVolume: ", cylinder.getVolume())
				fmt.Println("Lateral surface area: ", cylinder.getLateralSurfaceArea())
				fmt.Println("Total surface area: ", cylinder.getTotalSurfaceArea())
			}
		case 6:
			{
				fmt.Println("Calculating the volume, lateral and total surface area of a Cone:")
				fmt.Print("Enter the radius: ")
				fmt.Scan(&radius)
				fmt.Print("Enter the height: ")
				fmt.Scan(&height)

				cone := Cone{radius, height}

				fmt.Println("\nVolume: ", cone.getVolume())
				fmt.Println("Lateral surface area: ", cone.getLateralSurfaceArea())
				fmt.Println("Total surface area: ", cone.getTotalSurfaceArea())

			}

		}
		//break out of the loop
		if choice == 999 {
			break
		}
	}
}
