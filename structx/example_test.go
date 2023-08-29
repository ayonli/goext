package structx_test

import (
	"fmt"
	"testing"

	"github.com/ayonli/goext/async"
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
	fields, err := async.Try(func() ([]string, error) {
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
	fields, err := async.Try(func() ([]string, error) {
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
	pairs, err := async.Try(func() ([]string, error) {
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
	record2 := structx.Pick(record1, []string{"Name", "Email"})

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
