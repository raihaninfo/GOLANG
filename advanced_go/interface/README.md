## What is an Interface?

An interface is an abstract concept which enables polymorphism in Go. interface is declared as a type. Here is the declaration that is used to declare an interface.

```go
type interface_name interface{}
```

## zero value of interface

The zero value of an interface is nil. That means it holds no value and type.

```go
func main() {
    var a interface{}

    fmt.Println(a)    // <nil>
}
```
