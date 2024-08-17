package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

// Interface Men implemented by Human, Student and Employee
type Men interface {
	SayHi()
	Sing(lyrics string)
}

// method
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// method
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// method
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
	student1 := Student{Human{"Debopriya Deb Roy", 24, "01990288897"}, "Aust", 0.00}
	student2 := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
    employee1 := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
    employee2 := Employee{Human{"Tom", 36, "444-222-XXX"}, "Things Ltd.", 5000}

	// define interface i
    var i Men

    //i can store Student
    i = student1
    fmt.Println("This is Mike, a Student:")
    i.SayHi()
    i.Sing("November rain")

    //i can store Employee
    i = employee1
    fmt.Println("This is Tom, an Employee:")
    i.SayHi()
    i.Sing("Born to be wild")

    // slice of Men
    fmt.Println("Let's use a slice of Men and see what happens")
    x := make([]Men, 3)
    // these three elements are different types but they all implemented interface Men
    x[0], x[1], x[2] = student2, employee2, employee1

    for _, value := range x {
        value.SayHi()
    }

}
