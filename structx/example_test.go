package structx_test

import (
	"fmt"
	"testing"

	"github.com/ayonli/goext"
	"github.com/ayonli/goext/structx"
	"github.com/stretchr/testify/assert"
)

func ExampleMerge() {
	type Person struct {
		Name  string
		Email string
	}

	record1 := Person{
		Name:  "A-yon Lee",
		Email: "ayonlys@gmail.com",
	}
	record2 := Person{
		Email: "the@ayon.li",
	}
	record3 := structx.Merge(record1, record2)

	fmt.Printf("%#v\n", record3)
	fmt.Printf("%#v\n", record1) // record1 is not changed
	fmt.Printf("%#v\n", record2) // record2 is not changed
	// Output:
	// structx_test.Person{Name:"A-yon Lee", Email:"the@ayon.li"}
	// structx_test.Person{Name:"A-yon Lee", Email:"ayonlys@gmail.com"}
	// structx_test.Person{Name:"", Email:"the@ayon.li"}
}

func ExampleMerge_pointer() {
	type Person struct {
		Name  string
		Email string
	}

	record1 := &Person{
		Name:  "A-yon Lee",
		Email: "ayonlys@gmail.com",
	}
	record2 := &Person{
		Email: "the@ayon.li",
	}
	record3 := structx.Merge(record1, record2)

	fmt.Printf("%#v\n", record3)
	fmt.Printf("%#v\n", record1) // record1 is also changed
	fmt.Printf("%#v\n", record2) // record2 is not changed
	// Output:
	// &structx_test.Person{Name:"A-yon Lee", Email:"the@ayon.li"}
	// &structx_test.Person{Name:"A-yon Lee", Email:"the@ayon.li"}
	// &structx_test.Person{Name:"", Email:"the@ayon.li"}
}

func ExamplePatch() {
	type Person struct {
		Name  string
		Email string
	}

	record1 := Person{
		Name: "A-yon Lee",
	}
	record2 := Person{
		Name:  "John Doe",
		Email: "the@ayon.li",
	}
	record3 := structx.Patch(record1, record2)

	fmt.Printf("%#v\n", record3)
	fmt.Printf("%#v\n", record1) // record1 is not changed
	fmt.Printf("%#v\n", record2) // record2 is not changed
	// Output:
	// structx_test.Person{Name:"A-yon Lee", Email:"the@ayon.li"}
	// structx_test.Person{Name:"A-yon Lee", Email:""}
	// structx_test.Person{Name:"John Doe", Email:"the@ayon.li"}
}

func ExamplePatch_pointer() {
	type Person struct {
		Name  string
		Email string
	}

	record1 := &Person{
		Name: "A-yon Lee",
	}
	record2 := &Person{
		Name:  "John Doe",
		Email: "the@ayon.li",
	}
	record3 := structx.Patch(record1, record2)

	fmt.Printf("%#v\n", record3)
	fmt.Printf("%#v\n", record1) // record1 is also changed
	fmt.Printf("%#v\n", record2) // record2 is not changed
	// Output:
	// &structx_test.Person{Name:"A-yon Lee", Email:"the@ayon.li"}
	// &structx_test.Person{Name:"A-yon Lee", Email:"the@ayon.li"}
	// &structx_test.Person{Name:"John Doe", Email:"the@ayon.li"}
}

func ExampleFields() {
	type Person struct {
		Name  string
		Email string
	}

	person := Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}
	fields := structx.Fields(person)

	fmt.Println(fields)
	// Output:
	// [Name Email]
}

func ExampleFields_pointer() {
	type Person struct {
		Name  string
		Email string
	}

	person := &Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}
	fields := structx.Fields(person)

	fmt.Println(fields)
	// Output:
	// [Name Email]
}

func TestFieldsNonStruct(t *testing.T) {
	person := map[string]string{
		"Name":  "A-yon Lee",
		"Email": "the@ayon.li",
	}
	fields, err := goext.Try(func() ([]string, error) {
		fields := structx.Fields(person)
		return fields, nil
	})

	assert.Equal(t, []string(nil), fields)
	assert.Equal(t, "the argument of structx.Fields() must be a struct", err.Error())
}

func ExampleValues() {
	type Person struct {
		Name  string
		Email string
	}

	person := Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}
	fields := structx.Values[string](person)

	fmt.Println(fields)
	// Output:
	// [A-yon Lee the@ayon.li]
}

func ExampleValues_pointer() {
	type Person struct {
		Name  string
		Email string
	}

	person := &Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}
	fields := structx.Values[string](person)

	fmt.Println(fields)
	// Output:
	// [A-yon Lee the@ayon.li]
}

func TestValuesNonStruct(t *testing.T) {
	person := map[string]string{
		"Name":  "A-yon Lee",
		"Email": "the@ayon.li",
	}
	fields, err := goext.Try(func() ([]string, error) {
		fields := structx.Values[string](person)
		return fields, nil
	})

	assert.Equal(t, []string(nil), fields)
	assert.Equal(t, "the argument of structx.Values() must be a struct", err.Error())
}

func ExampleForEach() {
	type Person struct {
		Name  string
		Email string
	}

	person := Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}

	structx.ForEach[string](person, func(value, key string) {
		fmt.Println(key, "=>", value)
	})
	// Output:
	// Name => A-yon Lee
	// Email => the@ayon.li
}

func ExampleForEach_pointer() {
	type Person struct {
		Name  string
		Email string
	}

	person := &Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}

	structx.ForEach[string](person, func(value, key string) {
		fmt.Println(key, "=>", value)
	})
	// Output:
	// Name => A-yon Lee
	// Email => the@ayon.li
}

func TestForEachNonStruct(t *testing.T) {
	person := map[string]string{
		"Name":  "A-yon Lee",
		"Email": "the@ayon.li",
	}
	pairs, err := goext.Try(func() ([]string, error) {
		pairs := []string{}

		structx.ForEach[string](person, func(value, key string) {
			pairs = append(pairs, fmt.Sprint(key, "=>", value))
		})

		return pairs, nil
	})

	assert.Equal(t, []string(nil), pairs)
	assert.Equal(t, "the argument of structx.ForEach() must be a struct", err.Error())
}

func ExamplePick() {
	type Person struct {
		Name   string
		Email  string
		Gender int
		Age    int
	}

	record1 := Person{
		Name:   "A-yon Lee",
		Email:  "the@ayon.li",
		Gender: 1,
		Age:    28,
	}
	record2 := structx.Pick(record1, []string{
		"Name",
		"Email",
		"Other", // unrecognized field is ignored
	})

	fmt.Printf("%#v\n", record2)
	// Output:
	// structx_test.Person{Name:"A-yon Lee", Email:"the@ayon.li", Gender:0, Age:0}
}

func ExamplePick_pointer() {
	type Person struct {
		Name   string
		Email  string
		Gender int
		Age    int
	}

	record1 := &Person{
		Name:   "A-yon Lee",
		Email:  "the@ayon.li",
		Gender: 1,
		Age:    28,
	}
	record2 := structx.Pick(record1, []string{"Name", "Email"})

	fmt.Printf("%#v\n", record2)
	// Output:
	// &structx_test.Person{Name:"A-yon Lee", Email:"the@ayon.li", Gender:0, Age:0}
}

func ExampleOmit() {
	type Person struct {
		Name   string
		Email  string
		Gender int
		Age    int
	}

	record1 := Person{
		Name:   "A-yon Lee",
		Email:  "the@ayon.li",
		Gender: 1,
		Age:    28,
	}
	record2 := structx.Omit(record1, []string{"Gender", "Age"})

	fmt.Printf("%#v\n", record2)
	// Output:
	// structx_test.Person{Name:"A-yon Lee", Email:"the@ayon.li", Gender:0, Age:0}
}

func ExampleOmit_pointer() {
	type Person struct {
		Name   string
		Email  string
		Gender int
		Age    int
	}

	record1 := &Person{
		Name:   "A-yon Lee",
		Email:  "the@ayon.li",
		Gender: 1,
		Age:    28,
	}
	record2 := structx.Omit(record1, []string{"Gender", "Age"})

	fmt.Printf("%#v\n", record2)
	// Output:
	// &structx_test.Person{Name:"A-yon Lee", Email:"the@ayon.li", Gender:0, Age:0}
}

func ExampleSet() {
	type Person struct {
		Name  string
		Email string
	}

	person := Person{}
	ok1 := structx.Set(&person, "Name", "A-yon Lee")
	ok2 := structx.Set(&person, "Gender", "MALE")

	fmt.Printf("%#v\n", person)
	fmt.Println(ok1)
	fmt.Println(ok2)
	// Output:
	// structx_test.Person{Name:"A-yon Lee", Email:""}
	// true
	// false
}

func TestSetNonPointer(t *testing.T) {
	type Person struct {
		Name  string
		Email string
	}

	person := Person{}

	ok, err := goext.Try(func() (bool, error) {
		return structx.Set(person, "Name", "A-yon Lee"), nil
	})

	assert.Equal(t, Person{}, person)
	assert.False(t, ok)
	assert.Equal(t,
		"the first argument passed to structx.Set() must be a pointer of a struct",
		err.Error())
}

func ExampleHas() {
	type Person struct {
		Name   string
		Email  string
		Gender string
	}

	person := Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}
	ok1 := structx.Has(person, "Name")
	ok2 := structx.Has(person, "Gender")
	ok3 := structx.Has(person, "Age")

	fmt.Println(ok1)
	fmt.Println(ok2)
	fmt.Println(ok3)
	// Output:
	// true
	// true
	// false
}

func ExampleHas_pointer() {
	type Person struct {
		Name   string
		Email  string
		Gender string
	}

	person := &Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}
	ok1 := structx.Has(person, "Name")
	ok2 := structx.Has(person, "Gender")
	ok3 := structx.Has(person, "Age")

	fmt.Println(ok1)
	fmt.Println(ok2)
	fmt.Println(ok3)
	// Output:
	// true
	// true
	// false
}

func ExampleGet() {
	type Person struct {
		Name   string
		Email  string
		Gender string
	}

	person := Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}
	name, ok1 := structx.Get[string](person, "Name")
	gender, ok2 := structx.Get[string](person, "Gender")
	age, ok3 := structx.Get[int](person, "Age")

	fmt.Println(name, ok1)
	fmt.Println(gender, ok2)
	fmt.Println(age, ok3)
	// Output:
	// A-yon Lee true
	//  true
	// 0 false
}

func ExampleGet_pointer() {
	type Person struct {
		Name   string
		Email  string
		Gender string
	}

	person := &Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}
	name, ok1 := structx.Get[string](person, "Name")
	gender, ok2 := structx.Get[string](person, "Gender")
	age, ok3 := structx.Get[int](person, "Age")

	fmt.Println(name, ok1)
	fmt.Println(gender, ok2)
	fmt.Println(age, ok3)
	// Output:
	// A-yon Lee true
	//  true
	// 0 false
}

type Person struct {
	Name   string
	Email  string
	Gender string
}

func (self Person) GetName() string {
	return self.Name
}

func (self *Person) Greet(verb string) string {
	return verb + ", " + self.Name
}

func ExampleHasMethod() {
	person := Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}
	ok1 := structx.HasMethod(person, "GetName")
	ok2 := structx.HasMethod(person, "Greet")
	ok3 := structx.HasMethod(person, "GetEmail")

	fmt.Println(ok1)
	fmt.Println(ok2)
	fmt.Println(ok3)
	// Output:
	// true
	// false
	// false
}

func ExampleHasMethod_pointer() {
	person := &Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}
	ok1 := structx.HasMethod(person, "GetName")
	ok2 := structx.HasMethod(person, "Greet")
	ok3 := structx.HasMethod(person, "GetEmail")

	fmt.Println(ok1)
	fmt.Println(ok2)
	fmt.Println(ok3)
	// Output:
	// true
	// true
	// false
}

func TestHasMethodNonStruct(t *testing.T) {
	_, err := goext.Try(func() (bool, error) {
		return structx.HasMethod("", "method"), nil
	})

	assert.Equal(t, "the argument of structx.HasMethod() must be a struct", err.Error())
}

func ExampleCallMethod() {
	person := &Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}
	returns1 := structx.CallMethod(person, "GetName")
	returns2 := structx.CallMethod(person, "Greet", "Hello")
	_, err := goext.Try(func() ([]any, error) {
		return structx.CallMethod(person, "GetEmail"), nil
	})

	fmt.Println(returns1...)
	fmt.Println(returns2...)
	fmt.Println(err)
	// Output:
	// A-yon Lee
	// Hello, A-yon Lee
	// method GetEmail() doesn't exist on *structx_test.Person
}

func TestCallMethodNonPointer(t *testing.T) {
	person := Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}
	returns1 := structx.CallMethod(person, "GetName")
	_, err1 := goext.Try(func() ([]any, error) {
		return structx.CallMethod(person, "Greet", "Hello"), nil
	})
	_, err2 := goext.Try(func() ([]any, error) {
		return structx.CallMethod(person, "GetEmail"), nil
	})

	assert.Equal(t, []any{"A-yon Lee"}, returns1)
	assert.Equal(t, "method Greet() doesn't exist on structx_test.Person", err1.Error())
	assert.Equal(t, "method GetEmail() doesn't exist on structx_test.Person", err2.Error())
}

func TestCallMethodNonStruct(t *testing.T) {
	_, err := goext.Try(func() ([]any, error) {
		return structx.CallMethod("", "method"), nil
	})

	assert.Equal(t, "the argument of structx.CallMethod() must be a struct", err.Error())
}

func ExampleToMap() {
	type Person struct {
		Name  string
		Email string
	}

	person := Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}
	record := structx.ToMap[string](person)

	fmt.Println(record)
	// Output:
	// map[Email:the@ayon.li Name:A-yon Lee]
}

func ExampleToMap_pointer() {
	type Person struct {
		Name  string
		Email string
	}

	person := &Person{
		Name:  "A-yon Lee",
		Email: "the@ayon.li",
	}
	record := structx.ToMap[string](person)

	fmt.Println(record)
	// Output:
	// map[Email:the@ayon.li Name:A-yon Lee]
}

func ExampleFromMap() {
	type Person struct {
		Name  string
		Email string
	}

	record := map[string]string{
		"Name":  "A-yon Lee",
		"Email": "the@ayon.li",
	}
	person := structx.FromMap[string, Person](record)

	fmt.Printf("%#v\n", person)
	// Output:
	// structx_test.Person{Name:"A-yon Lee", Email:"the@ayon.li"}
}

func ExampleFromMap_pointer() {
	type Person struct {
		Name  string
		Email string
	}

	record := map[string]string{
		"Name":  "A-yon Lee",
		"Email": "the@ayon.li",
	}
	person := structx.FromMap[string, *Person](record)

	fmt.Printf("%#v\n", person)
	// Output:
	// &structx_test.Person{Name:"A-yon Lee", Email:"the@ayon.li"}
}
