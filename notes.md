# GO

### Var Types
- string
- int
  - i8/16/32/64
  - u8/16/32/64
- float
  - 32/64
- bool
- runes (char)
- nil

## Declaration
```Go
// Explicit
var name string = "Hello"
var isErorr bool = false

// Inferred
name := world
isError := false
```

## Err handling
```Go
// Err value can be assigned from common methods
decimal := "3.14259"
decAsFloat, err := strconv.ParseFloat(decimal, 64)

// Handle error as value
// Guard clause
if err != nil {
	log.Fatal(err)
}
fmt.Println(decAsFloat)

// Declare and error handle in clause
if decAsFloat, err := strconv.ParseFloat(decimal, 64); err == nil {
    fmt.Println(decAsFloat)
} else {
    log.Fatal(err)
}
```