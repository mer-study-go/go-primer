### Project Overview

![Screen Shot 2018-08-14 at 5.38.58 PM.png](resources/44C73AAFDE2E59EF186EC326E050FCD6.png)

### New Project Folder

src/cards

### Variable Declarations

src/cards/main.go

```go
func main() {
	var card = "Ace of Spades"
	fmt.Println(card)
}
```

![Screen Shot 2018-08-14 at 5.52.57 PM.png](resources/FC84E23E1F0B47E378E0E9A5ED1029B7.png)

![Screen Shot 2018-08-14 at 5.53.52 PM.png](resources/3B6EE82AAE386DCFB2973B0C6D2B62A4.png)

![Screen Shot 2018-08-14 at 5.55.43 PM.png](resources/C17D218A02B49A72968F3B838B13A1B5.png)



```go
func main() {
	// var card = "Ace of Spades"
	card := "Ace of Spades"
	fmt.Println(card)
}
```

### Functions and Return Types

```go
func main() {
	card := newCard()
	fmt.Println(card)
}
```

```go
func newCard() string {
	return "Five of Diamonds"
}
```

![Screen Shot 2019-01-31 at 17.54.05.png](resources/F0635B90301813E2AD1DED5695176B0E.png)

### Slices and For Loops

![Screen Shot 2018-08-15 at 11.31.34 AM.png](resources/08CB9693A5ED6C88F8E3B5D6826F93FD.png)

![Screen Shot 2018-08-15 at 11.33.41 AM.png](resources/47CEF9E3704A4F09657DAB6C7C4F0399.png)

**Slice example and Add new Item**

```go
func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	// Add new item
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)
}
```

**Iterate slice**

```go
func main() {
	cards := []string{"Ace of Diamonds", newCard(), "Six of Spades"}

	// Iterate slice
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
```

![Screen Shot 2019-02-04 at 12.01.36.png](resources/C5F70E4898EE9EA6C7E18C75D3F844F7.png)

### OO Approach vs Go Approach

![Screen Shot 2019-02-04 at 12.24.26.png](resources/C160717C3BB06F2607B6024A5F558ED0.png)

![Screen Shot 2019-02-04 at 13.10.08.png](resources/DA90FC36D37DFD6C5B313EEFF9DCB280.png)

![Screen Shot 2019-02-04 at 13.14.58.png](resources/B96B70D8505B6DEE14A8D271480FE869.png)



### Custom Type Declarations

/src/cards/deck.go

```go
package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

```

/src/cards/main.go

```go
package main

func main() {
    cards := deck{"Ace of Diamonds", newCard()}
    cards = append(cards, "Six of Spades")

    cards.print()
}

func newCard() string {
    return "Five of Diamonds"
}
```

```sh
$ go run main.go deck.go
```

### Receiver Functions

![Screen Shot 2019-02-04 at 14.36.13.png](resources/0D04DBB775E4B6A5BD323EAB79213744.png)

![Screen Shot 2019-02-04 at 14.39.30.png](resources/C2FF4437A7C7F9908F001CAA37AA7D4E.png)

### Creating a New Deck

![Screen Shot 2019-02-04 at 15.02.22.png](resources/4AA1A3591A88C0633C090194F15B82B8.png)

/src/cards/deck.go

```go
// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards

}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}
```

/src/cards/main.go

```go
func main() {
    cards := newDeck()
    cards.print()
}
```

### Slice Range Syntax

![Screen Shot 2019-02-04 at 15.15.07.png](resources/0FF27D6CFB8A3FE5261648F1EE246AE6.png)

![Screen Shot 2018-08-15 at 4.22.42 PM.png](resources/BC8865648E62085B2284B8AC6ADD09B9.png)

### Multipole Return Values

/src/cards/deck.go

```go
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
```

/src/cards/main.go

```go
func main() {
	cards := newDeck() 
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}
```

![Screen Shot 2019-02-04 at 15.32.36.png](resources/4CF1190C8F534EB4816D439D601FD5FD.png)

### Byte Slices

<https://golang.org/pkg/io/ioutil/>

![Screen Shot 2019-02-04 at 15.43.21.png](resources/75459D2BEFF9799773692BEF9DA57163.png)



### Deck to String

**Type Conversion**

![Screen Shot 2019-02-04 at 15.48.35.png](resources/66B4312FBD45232590FB44B965D4D7AD.png)

/src/cards/main.go

```go
func main() {
	greeting := "Hi there! "
	fmt.Println([]byte(greeting))
}
```

![Screen Shot 2019-02-04 at 15.50.39.png](resources/8C1A0F2C6263FA8C393245B29E89472C.png)



### Joining a Slice of Strings

/src/cards/deck.go

```go
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
```

[https://golang.org/pkg/strings/\#Join](https://golang.org/pkg/strings/#Join)

/src/cards/main.go

```go
func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
}
```

### Saving Data to the Hard Drive

/src/cards/deck.go

```go
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
```

/src/cards/main.go

```go
func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
}
```

### Reading From the Hard Drive

![Screen Shot 2018-08-16 at 2.10.33 PM.png](resources/85D4320A3609BB2B277C290029B393C2.png)

[https://golang.org/pkg/io/ioutil/\#ReadFile](https://golang.org/pkg/io/ioutil/#ReadFile)

/src/cards/deck.go

```go
func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 -log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	stringSlice := strings.Split(string(byteSlice), ",") // Ace of Spades,Two of Spades,Three of Spades,...
	return deck(stringSlice)
}
```

/src/cards/main.go

```go
func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
}
```

/src/cards/main.go (error)

```go
func main() {
	cards := newDeckFromFile("my")
	cards.print()
}
```

### Shuffling a Deck

![Screen Shot 2019-02-04 at 17.02.06.png](resources/610C4D522767BD8BDA0B5B3BA2653307.png)

![Screen Shot 2019-02-04 at 17.02.48.png](resources/25B1BB49B033BE04775E7957A6250DBB.png)

[https://golang.org/pkg/math/rand/\#Intn](https://golang.org/pkg/math/rand/#Intn)

/src/cards/deck.go

```go
func (d deck) shuffle() {
	for index := range d {
		newPosition := rand.Intn(len(d) - 1)
		// swap item in index and in newPosition
		d[index], d[newPosition] = d[newPosition], d[index]

	}
}
```

/src/cards/main.go

```go
func main() {
    cards := newDeck()
    cards.shuffle()
    cards.print()
}
```

### Random Number Generation

![Screen Shot 2019-02-04 at 17.14.48.png](resources/5DB10A0CD15CABCE56083CDC286303A5.png)

[https://golang.org/pkg/math/rand/\#Rand](https://golang.org/pkg/math/rand/#Rand)

Better Approach to generate random number to change numbers every time

/src/cards/deck.go

```go
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
```

### Testing With Go

![Screen Shot 2019-02-04 at 17.52.24.png](resources/BA9F9688D93E48111FD9A470C3336988.png)

### Writing Useful Tests

![Screen Shot 2019-02-04 at 17.56.49.png](resources/3398344C892A5D04248CCD552864317F.png)

To make a test, create a new file ending in \_test.go

/src/cards/deck\_test.go

```go
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Cubes" {
		t.Errorf("Expected last card of Four of Cubes, but got %v", d[len(d)-1])
	}
}
```

### Testing File IO

```go
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
```

### Assignment 1: Even and Odd

```go
func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, " is even")
		} else {
			fmt.Println(number, " is odd")
		}
	}
}
```

































