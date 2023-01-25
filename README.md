# env
A small library for populating structs with environment values.

> No external dependencies, this is only dealing with environment and no files.

An example how to use it would be 
```go
type MyStruct struct {
	Field1 string `env:"FIELD_1"`
	Field2 int    `env:"FIELD_2"`
}

myStruct := MyStruct{}
env.PopulateWithEnv(&myStruct)

//you should see the data in your struct
fmt.Println(myStruct)

```

It supports also structs inside structs by taging it with `obj:""`
```go
type MyStruct struct {
	Field1 string `env:"FIELD_1"`
	Field2 int    `env:"FIELD_2"`
}

type MainStruct struct {
	MyStruct `obj:"ref"`
}

struct := MainStruct{}
env.PopulateWithEnv(&struct)

//you should see the data in your struct
fmt.Println(struct)
```
