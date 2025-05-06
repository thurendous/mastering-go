# Go Multiple Variable Declaration

```go
package main
import ("fmt")

func main() {
    var one, two, three, four, five int = 1, 2, 3, 4, 5

    fmt Println(one)
    fmt Println(two)
    // write more text here:

}

```


# Go Vairable Decaration in a Block
```go
package main

import ("fmt")

func main() {
    var (
        num int
        number int = 1
        greetings string = "hello man"
    )

    fmt.Println(num)
    fmt.Pringln(number)
    fmt.Println(greetings)
}
```

# Go Variable Naming
1. A variable name must start with a letter or an underscore character(_)
2. A variable name cannot start with a digit
3. A variable name can only contain alpha-numeric characters and underscors(a-z, A-Z, 0-9 and _)
4. Variable names are case-sensitive(age, Age and AGE are three different variables)
5. There is no limit on the length of the variable name
6. A variable name cannot contain spaces
7. The variable name cannot be any Go keywords




