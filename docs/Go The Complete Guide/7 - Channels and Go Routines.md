### Website Status Checker

Non concurrent method

Main.go

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
```

### Receiving Messages

main.go

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    links := []string{
        "http://linkedin.com",
        "http://google.com",
        "http://facebook.com",
        "http://stackoverflow.com",
        "http://golang.org",
        "http://amazon.com",
    }

    c := make(chan string)

    for _, link := range links {
        go checkLink(link, c)
    }

    for i := 0; i < len(links); i++ {
        fmt.Println(<-c)
    }
}

func checkLink(link string, c chan string) {
    _, err := http.Get(link)
    if err != nil {
        fmt.Println(link, "might be down!")
        c <- "Might be down I think"
        return
    }

    fmt.Println(link, "is up!")
    c <- "Yep it's up"
}
```

### Repeating Routines

main.go

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://linkedin.com",
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
```

### Alternative Loop Syntax

main.go

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    links := []string{
        "http://linkedin.com",
        "http://google.com",
        "http://facebook.com",
        "http://stackoverflow.com",
        "http://golang.org",
        "http://amazon.com",
    }

    c := make(chan string)

    for _, link := range links {
        go checkLink(link, c)
    }

    for l := range c {
        go checkLink(l, c)
    }
}

func checkLink(link string, c chan string) {
    _, err := http.Get(link)
    if err != nil {
        fmt.Println(link, "might be down!")
        c <- link
        return
    }

    fmt.Println(link, "is up!")
    c <- link
}

```

### Function Literals

main.go

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://linkedin.com",
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
```









































