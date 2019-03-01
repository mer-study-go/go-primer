### Structs in Go

![Screen Shot 2019-02-04 at 18.10.54.png](resources/A03BDF9E75EE72330C3B4EA0CED955E4.png)



**Struct**

![Screen Shot 2019-02-04 at 18.12.31.png](resources/7722D64C77902EE720C4020C7B989E0B.png)

![Screen Shot 2019-02-04 at 18.12.57.png](resources/0AEFA150BCF9C9D3F6A167CCEC567EA7.png)

### Defining Structs

![Screen Shot 2019-02-04 at 18.16.09.png](resources/91D48BBC0916B4FC98E67594F76CB35D.png)

/src/struct/main.go

```go
type person struct {
	firstName string
	lastName  string
}
```

### Declaring Structs

/src/struct/main.go

```go
func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
```

### Updating Struct Values

![Screen Shot 2019-02-04 at 18.21.46.png](resources/FFE90363767AF6037B9B2294771C6572.png)

/src/struct/main.go

```go
func main() {
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
```

### Embedding Structs

![Screen Shot 2019-02-04 at 18.23.38.png](resources/644090550157E402E3D39371AF36973D.png)

/src/struct/main.go

```go
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	mer := person{
		firstName: "Mer",
		lastName:  "JQ",
		contact: contactInfo{
			email:   "mer@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Println(mer)
}
```

### Structs with Receiver Functions

src/struct/main.go

```go
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Println(jim)
}
```

**Add receiver functions**

/src/struct/main.go

```go
func main() {
  mer.updateName("Jieqiong")
  mer.print()
}

func (p person) print() {
    fmt.Println(p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
```

### Pass By Value

![Screen Shot 2019-02-04 at 18.31.43.png](resources/9E9F5B93C6291493D6541EE036C594B7.png)

![Screen Shot 2019-02-04 at 18.33.33.png](resources/D4BDB12F2D6555FEB01D1AC357BD2509.png)

### Structs with Pointers

/src/struct/main.go

```go
func main() {
  merPointer := &mer
	merPointer.updateName("Jieqiong")
	mer.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
```

### Pointer Operations

**&variable**

Give me the memory address of the value this variable is pointing at

**\*pointer**

Give me the value this memory address is pointing at

![Screen Shot 2019-02-04 at 18.36.34.png](resources/73C079C8DB52EA69377F67DF1DFC4054.png)

![Screen Shot 2019-02-04 at 18.36.52.png](resources/49099D839CF81F4636C92B3BFF00A57B.png)

![Screen Shot 2018-08-20 at 3.48.16 PM.png](resources/B95878E5A4845D95C86A80F77E338266.png)

![Screen Shot 2018-08-20 at 3.52.11 PM.png](resources/AA5CF16D412A38DA6964537EF2718CC7.png)

### Pointer Shortcut

main.go

```go
func main() {
	mer.updateName("Jieqiong")
	mer.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
```

![Screen Shot 2019-02-04 at 18.42.06.png](resources/2F66C9C0220E83C71F7AC9ABF5CD20C6.png)

### **Gotchas With Pointers**

/src/struct/main.go

```go
func main() {
    mySlice := []string{"Hi", "There", "How", "Are", "You"}
    updateSlice(mySlice)
    fmt.Println(mySlice)
}

func updateSlice(s []string) {
    s[0] = "Bye"
}
```

### Reference vs Value Types

![Screen Shot 2018-08-20 at 4.24.43 PM.png](resources/48DB9A1BE7D8E73ABE916DA72D87C959.png)

![Screen Shot 2018-08-20 at 4.25.56 PM.png](resources/3D7395B7D44F42FE2F08614E372983BE.png)

![Screen Shot 2018-08-20 at 4.27.32 PM.png](resources/69B3DB94532773CA2177A04D8608EBFA.png)





































