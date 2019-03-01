### Purpose of Interfaces

![Screen Shot 2018-08-21 at 11.11.23 AM.png](resources/3D9C8EAD1F6D2D3FD7BCD9B3AAEF7780.png)

![Screen Shot 2018-08-21 at 11.13.15 AM.png](resources/9442D423E15AC452D1B7B2187AA3DECF.png)

![Screen Shot 2019-02-04 at 20.19.36.png](resources/057411EF283F5F21C36BEAD61F9AAAF1.png)

### Interfaces in Practice

main.go

```go
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating an spanish greeting
	return "Hola!"
}
```

### Rules of Interfaces

![Screen Shot 2018-08-21 at 5.25.32 PM.png](resources/CC1C6254A6EF904491CFCB4DF33A2039.png)

### Extra Interface Notes

![Screen Shot 2018-08-21 at 5.30.49 PM.png](resources/478A69889C368365E3B4D286879400FB.png)

### The HTTP Package

![Screen Shot 2018-08-21 at 5.38.38 PM.png](resources/2AA222779AABF95C545FDF739485F25C.png)

### The Reader Interface

![Screen Shot 2018-08-21 at 6.25.18 PM.png](resources/4172FA60A6CF7CCF2188CA5375A3E4B3.png)

![Screen Shot 2018-08-21 at 6.28.31 PM.png](resources/66E25EF0373983FDD97FAA043D372B77.png)

### More on the Reader Interface

![Screen Shot 2018-08-21 at 6.34.02 PM.png](resources/483A06828AEC80C91804E15478C4EA7D.png)

### Working with the Read Function

main.go

```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
```

### The Writer Interface

main.go

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}
```

### A Custom Writer

main.go

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
```































