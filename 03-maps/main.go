package main

import "fmt"

var (
	choice, counter int
	employees       map[string]Employee
	name, country   string
	salary          float64
)

//Define struct
type Employee struct {
	country string
	salary  float64
}

//Define method
func (e Employee) displayInfo() {
	fmt.Printf("Salary: %v\n", e.salary)
	fmt.Printf("Residence: %v\n", e.country)
}

//print options to the screen
func printScreenOptions() {
	fmt.Println("\n****************************************")
	fmt.Println("This program demonstrates how maps work")
	fmt.Println("****************************************")
	fmt.Println("* Option 1: Add an item to a map")
	fmt.Println("* Option 2: Search map for a particular key")
	fmt.Println("* Option 3: Update an item in a map")
	fmt.Println("* Option 4: Delete an item from a map")
	fmt.Println("* Option 5: Display all items in a map")
	fmt.Println("* Option 6: Delete an entire map")
	fmt.Println("* Option 7: Get the length of the map")
	fmt.Println("* Option 8: Exit the program")
	fmt.Println("****************************************")
}

func searchMap(employees map[string]Employee) {
	fmt.Println("\nYou've selected option 2: [Search map for a particular key]")
	fmt.Print("\nEnter name of the employee: ")
	fmt.Scan(&name)

	_, ok := employees[name]
	if ok {
		fmt.Printf("\nEmployee %v found in the database\n", name)
	} else {
		fmt.Printf("\n%v's not an employee\n", name)

	}
}

//main function
func main() {

	// Create employees to use to initialize the map
	emp1 := Employee{
		salary:  12000,
		country: "USA",
	}
	emp2 := Employee{
		salary:  14000,
		country: "Canada",
	}
	emp3 := Employee{
		salary:  13000,
		country: "India",
	}
	emp4 := Employee{
		salary:  13000,
		country: "Zambia",
	}
	emp5 := Employee{
		salary:  13000,
		country: "Zimbabwe",
	}
	//initialize map
	employees = map[string]Employee{
		"Steve": emp1,
		"Jamie": emp2,
		"Mike":  emp3,
		"Bob":   emp4,
		"Sam":   emp5,
	}

	for counter = 1; ; counter++ {
		printScreenOptions()
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		if choice == 8 {
			break
		}
		switch choice {
		case 1:
			{
				fmt.Println("\nYou've selected option 1: [Add an item to a map]")
				fmt.Print("Enter name of the employee: ")
				fmt.Scan(&name)

				_, ok := employees[name]

				if !ok {
					//Get the rest of the details
					fmt.Print("Enter Country of residence: ")
					fmt.Scan(&country)
					fmt.Print("Enter Salary of the employee: ")
					fmt.Scan(&salary)

					//Add an item to the map
					employees[name] = Employee{country, salary}
				} else {
					//Employee found in the map. Ask the user to Select option 3 to update this item
					fmt.Printf("\nEmployee %v exists in the map ", name)
					fmt.Println("Please select option 3 to update the item ")

				}

			}

		case 2:
			{
				// Call the searchMap function
				searchMap(employees)
			}
		case 3:
			{
				var option string
				fmt.Println("\nYou've selected option 3: [Update an item in a map]")
				fmt.Print("Enter name of the employee: ")
				fmt.Scan(&name)

				//Search if name exisits in map
				emp, ok := employees[name]

				// If name exisits, run the code in the block
				if ok {
					fmt.Printf("\nEmployee %s found\n", name)
					fmt.Println("*************************************************************")
					fmt.Println("\nEnter A/a: To update Salary only")
					fmt.Println("Enter B/b: To update Country of residence")
					fmt.Println("Enter C/c: To update both the Salary and Country of residence")
					fmt.Println("*************************************************************")
					fmt.Print("\nEnter your choice: ")
					fmt.Scan(&option)

					switch option {
					case "A", "a":
						{
							fmt.Println("You've chosen option A/a [Update Salary only]")
							fmt.Printf("\nEmployee %s's old salary is %f\n", name, emp.salary)
							fmt.Printf("Enter %s's new Salary: ", name)
							fmt.Scan(&salary)
							emp.salary = salary
							employees[name] = emp
							fmt.Printf("%s's Salary updated\n", name)

						}
					case "B", "b":
						{
							fmt.Println("You've chosen option B/b [Update employee's country of residence only]")
							fmt.Printf("\nEmployee %s's old country of residence is:  %s\n", name, emp.country)
							fmt.Printf("Enter %s's new country of residence: ", name)
							fmt.Scan(&country)
							emp.country = country
							employees[name] = emp
							fmt.Printf("Employee %s's country of residence updated\n", name)
						}
					case "C", "c":
						{
							fmt.Println("You've chosen option C/c [Update Salary and Country of residence]")
							fmt.Printf("\nEmployee %s's old salary is %f\n", name, emp.salary)
							fmt.Printf("Enter %s's new Salary: ", name)
							fmt.Scan(&salary)
							emp.salary = salary
							fmt.Printf("\nEmployee %s's old country of residence is:  %s\n", name, emp.country)
							fmt.Printf("Enter %s's new country of residence: ", name)
							fmt.Scan(&country)
							emp.country = country
							//Update
							employees[name] = emp
							fmt.Printf("Employee %s's Salary and residence updated\n", name)

						}
					default:
						{
							fmt.Println("You have chosen the wrong option")
						}
					}

				} else { //if name deosn't exist
					fmt.Printf("Employee %s does not exisit in the map", name)

				}

			}
		case 4:
			{
				fmt.Print("\nEnter the name of the employee you wish to delete from the database: ")
				fmt.Scan(&name)
				//check if the name exists in the map
				_, ok := employees[name]
				if ok {
					fmt.Printf("\nDeleting %v from the database\n", name)
					// delete the key from the database
					delete(employees, name)
					fmt.Println("Operation completed successfuly")
				} else {
					fmt.Printf("\n%v not found in the database\n", name)
				}

			}
		case 5:
			{
				// Check if map is empty
				if len(employees) != 0 {
					for name, employee := range employees {
						fmt.Printf("\nName: %v\n", name)
						employee.displayInfo()
					}
				} else { //Map is empty
					fmt.Println("\nMap is empty.")
				}

			}
		case 6:
			{
				fmt.Println("\nDeleting the entire database...")
				employees = make(map[string]Employee)
				fmt.Println("Map deleted.....")
				fmt.Println(employees)
			}
		case 7:
			{
				fmt.Printf("\nThe map has %v employees\n", len(employees))
			}

		}

	}

}
