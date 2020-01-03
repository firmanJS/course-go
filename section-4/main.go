package main
import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	// contact contactInfo
	contactInfo
}

// func main(){
// 	var alex person
// 	alex.firstName = "Alex"
// 	alex.lastName = "Anderson"
// 	fmt.Println(alex)
// 	fmt.Printf("%+v",alex)
// }

// # print in main function
// func main(){
// 	jim := person {
// 		firstName:"Jim",
// 		lastName:"Party",
// 		contactInfo : contactInfo {
// 			email: "jim@yahoo.com",
// 			zipCode: 34,
// 		},
// 	}

// 	fmt.Printf("%+v",jim)
// }

// # struct with receiver function
// func main(){
// 	jim := person {
// 		firstName:"Jim",
// 		lastName:"Party",
// 		contactInfo : contactInfo {
// 			email: "jim@yahoo.com",
// 			zipCode: 34,
// 		},
// 	}

// 	jim.updateName("Jimmy")
// 	jim.print()
// }

// func (p person) updateName(newFirstName string){
// 	p.firstName = newFirstName
// }

// func (p person) print(){
// 	fmt.Printf("%+v",p)
// }

func main(){
	jim := person {
		firstName:"Jim",
		lastName:"Party",
		contactInfo : contactInfo {
			email: "jim@yahoo.com",
			zipCode: 34,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v",p)
}