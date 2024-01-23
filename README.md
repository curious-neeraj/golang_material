# Golang Study Material
Repository to keep some basic teaching material for golang handy.

Some Salient Feature of GO -

- Go is **Statically Typed Language**\
You either have to tell the type explicitly while declaring, or provide a value so that it may infer the type from the value.

- Go is **Strongly Typed Language**\
The operations allowed to be performed on variables are strictly based on their types (example, you cannot add (+) a string and int value)

- Go is a **Compiled Language**

- Go has a **Fast Compile Time**\
The compilation time of Go programs is pretty less. It is thus suitable even for production setup.

- Go has **Built-In Concurrency**\
Go inheritly supports concurrency via concepts like *go routines*.

- Go has **Built-In Garbage Collector**\
The garbage collection in Go is automatic which makes it simple to use for the users as they don't have to handle all this by themselves.

## Basics
Initialize Go Module
```
go mod init <module-name>
```

Build a Go program file to generate a binary file
```
go build <file-name-with-extension>
```

Run the compiled binary file
```
./<file-name>
```

Build and Run in one *Go* 😉
```
go run <file-name-with-extension>
```

## Data Types and Variables

>**Data-types :**
1. int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64 -- _default value = 0_
2. float32, float64 -- _default value = 0_
3. rune -- _default value = 0_
4. string -- _default value = ''_
5. bool -- _default value = false_

>**Variables :**

There are multiple ways of declaring a variable
1. _var_ \<variable-name> _data_type_
2. _var_ \<variable-name> = <value>
3. \<variable-name> := \<value>

## Array
- Same Type
- Fixed Length
- Indexable (Zero indexed)
- Contiguous Memory Allocation
```
var intArr1 [3]int // Filled with Default value of datatype
intArr2 := [3]int{1, 2, 3}
intArr3 := [...]int{1, 2, 3, 4, 5} // Without specifying exact length

```

## Slice
- Wrapper around Array
- Convenient Interface to sequence of data
```
var intSlice []int = []int{1, 2, 3, 4}
intSlice = append(intSlice, 5)

// Appending multiple elements to slice
intSlice2 := []int{6, 7}
intSlice = append(intSlice, intSlice2...)	
```

## Map
- Key, Value pair
- Always return a value, even if key doesn't exist
```
// Empty Map - map1 = map[]
// key type -> string, value type -> int
var map1 map[string]int = make(map[string]int)

//Initialized Map
map2 := map[string]int{"Neeraj": 11, "AI": 10}
	
// Delete value from map
delete(map2, "AI")
```

## String, Rune
- String is immutable in Go
- String is a collection of bytes in Go (instead of char)
```
// This would tell us that the type of indexed value in string isnt char
var str string = "neeraj"
fmt.Printf("value = %v, type = %T \n", str[0], str[0])
```