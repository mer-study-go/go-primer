### What’s a Map?

![Screen Shot 2019-02-04 at 19.01.03.png](resources/C97FBCE14A26FE731ACDEAB7CC615CEB.png)

![Screen Shot 2019-02-04 at 19.01.56.png](resources/C2239145D3AEC2EF9DDE5C825147212B.png)

/src/map/main.go

```go
func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors)
}
```

### Manipulating Maps

Another methods to create a map

```go
var colors map[string]string
```

```go
colors := make(map[string]string)
```

Assign value to the map

```go
colors["white"] = "#ffffff"
```

To delete key in map

```go
delete(colors, "white")
```

### Iterating Over Maps

![Screen Shot 2018-08-20 at 5.29.22 PM.png](resources/E1C0529A4EB558329CE50443F661F5DC.png)

main.go

```go
func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#rbf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
```

### Diffrerences Between Maps and Structs

![Screen Shot 2018-08-20 at 5.33.57 PM.png](resources/62C1AC0F5B1CCF7BF96A63C15C323BBD.png)



































