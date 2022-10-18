Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```

<nil>
false

An interface value is equal to nil only if both its value and dynamic type are nil

You can think of the interface value nil as typed, and nil without type doesn’t equal nil with type. If we convert nil to the correct type, the values are indeed equal.

fmt.Println(err == (*os.PathError)(nil)) // true

```
