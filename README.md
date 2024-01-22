# golang_material
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
1. int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64 -- default value = 0
2. float32, float64 -- default value = 0
3. rune -- default value = 0
4. string -- default value = ''
5. bool -- default value = false

>**Variables :**
There are multiple ways of declaring a variable
1. _var_ \<variable-name> _data_type_
2. _var_ \<variable-name> = <value>
3. \<variable-name> := \<value>
