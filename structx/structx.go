// Functions used to manipulate structs.
package structx

import (
	"fmt"
	"reflect"

	"github.com/ayonli/goext/slicex"
)

func doMerge[T any](fnName string, overwrite bool, target T, sources ...T) T {
	if reflect.ValueOf(target).Kind() == reflect.Pointer {
		targetValue := reflect.ValueOf(target).Elem()
		fields := Fields(target)

		for _, source := range sources {
			sourceValue := reflect.ValueOf(source).Elem()

			for _, field := range fields {
				targetField := targetValue.FieldByName(field)
				sourceField := sourceValue.FieldByName(field)

				if targetField.CanSet() && !sourceField.IsZero() && (overwrite || targetField.IsZero()) {
					targetField.Set(sourceField)
				}
			}
		}

		return target
	} else {
		targetValue := reflect.ValueOf(new(T)).Elem()
		fields := Fields(target)
		sources = append([]T{target}, sources...)

		for _, source := range sources {
			sourceValue := reflect.ValueOf(source)

			for _, field := range fields {
				targetField := targetValue.FieldByName(field)
				sourceField := sourceValue.FieldByName(field)

				if targetField.CanSet() && !sourceField.IsZero() && (overwrite || targetField.IsZero()) {
					targetField.Set(sourceField)
				}
			}
		}

		return targetValue.Interface().(T)
	}
}

// Performs a shallow merge of two or more structs. The later field-value pairs override the
// existing ones or the ones before them.
//
// This function only works on the exported fields. If the input arguments are pointers, it mutates
// the first one and returns it, otherwise, a new struct is created.
func Merge[T any](first T, others ...T) T {
	return doMerge("Merge", true, first, others...)
}

// Performs a shallow merge of two or more structs, copies the field-value pairs that are presented
// in the later structs but are missing in the former into it, later pairs are skipped if the same
// field in the former struct is not the 0-value of its type.
//
// This function only works on the exported fields. If the input arguments are pointers, it mutates
// the first one and returns it, otherwise, a new struct is created.
func Patch[T any](first T, others ...T) T {
	return doMerge("Patch", false, first, others...)
}

func ensureValue[T any](fnName string, target T) any {
	var _target any

	if reflect.ValueOf(target).Kind() == reflect.Pointer {
		_target = reflect.ValueOf(target).Elem().Interface()
	} else {
		_target = target
	}

	if reflect.TypeOf(_target).Kind() != reflect.Struct {
		panic(fmt.Sprintf("the argument of structx.%s() must be a struct", fnName))
	}

	return _target
}

// Returns the fields of the given struct.
//
// This function only collects the exported fields, it panics if the given argument is not a struct.
func Fields[S any](target S) []string {
	targetValue := reflect.ValueOf(ensureValue("Fields", target))
	numField := targetValue.NumField()
	fields := []string{}

	for i := 0; i < numField; i++ {
		field := targetValue.Type().Field(i)

		if field.IsExported() {
			fields = append(fields, field.Name)
		}
	}

	return fields
}

// Returns the values of the given struct.
//
// This function only collects the exported fields, it panics if the given argument is not a struct.
func Values[V any](target any) []V {
	targetValue := reflect.ValueOf(ensureValue("Values", target))
	numField := targetValue.NumField()
	values := []V{}

	for i := 0; i < numField; i++ {
		fieldType := targetValue.Type().Field(i)

		if fieldType.IsExported() {
			fieldValue := targetValue.Field(i)
			values = append(values, fieldValue.Interface().(V))
		}
	}

	return values
}

// Executes a provided function once for each field-value pair.
//
// This function only loops the exported fields, it panics if the given argument is not a struct.
func ForEach[V any](target any, fn func(value V, field string)) {
	targetValue := reflect.ValueOf(ensureValue("ForEach", target))
	numField := targetValue.NumField()

	for i := 0; i < numField; i++ {
		fieldType := targetValue.Type().Field(i)

		if fieldType.IsExported() {
			fieldValue := targetValue.Field(i)
			fn(fieldValue.Interface().(V), fieldType.Name)
		}
	}
}

// Creates a new struct based on the original struct but only contains the specified fields. Omitted
// fields will be set to the 0-values of their types.
//
// This function only works on the exported fields.
func Pick[S any](original S, fields []string) S {
	originalValue := reflect.ValueOf(ensureValue("Pick", original))
	var newValue reflect.Value
	var isPtr bool

	if reflect.ValueOf(original).Kind() == reflect.Pointer {
		newValue = reflect.New(reflect.TypeOf(original).Elem())
		isPtr = true
	} else {
		newValue = reflect.ValueOf(new(S))
	}

	for _, field := range fields {
		newField := newValue.Elem().FieldByName(field)
		originalField := originalValue.FieldByName(field)

		if newField.CanSet() && originalField.IsValid() {
			newField.Set(originalField)
		}
	}

	if isPtr {
		return newValue.Interface().(S)
	} else {
		return newValue.Elem().Interface().(S)
	}
}

// Creates a new struct based on the original struct but without the specified fields. Omitted
// fields will be set to the 0-values of their types.
//
// This function only works on the exported fields.
func Omit[S any](original S, fields []string) S {
	allFields := Fields(original)
	keptFields := slicex.Diff(allFields, fields)
	return Pick(original, keptFields)
}

// Sets the field of the struct to be the given value.
//
// This function takes a pointer of the struct instead of its value, and it only supports the
// exported fields of the struct.
func Set[S any](target S, field string, value any) bool {
	targetValue := reflect.ValueOf(target)

	if targetValue.Kind() != reflect.Pointer || targetValue.Elem().Kind() != reflect.Struct {
		panic("the first argument passed to structx.Set() must be a pointer of a struct")
	}

	fieldValue := targetValue.Elem().FieldByName(field)

	if fieldValue.CanSet() {
		fieldValue.Set(reflect.ValueOf(value))
		return true
	} else {
		return false
	}
}

func Has[S any](target S, field string) bool {
	targetValue := reflect.ValueOf(ensureValue("Has", target))
	return targetValue.FieldByName(field).IsValid()
}

// Returns the value of the specified field of the struct.
//
// This function only supports the exported fields, it panics if the given argument is not a struct.
func Get[V any](target any, field string) (V, bool) {
	targetValue := reflect.ValueOf(ensureValue("Get", target))
	fieldValue := targetValue.FieldByName(field)

	if fieldValue.IsValid() {
		return fieldValue.Interface().(V), true
	} else {
		return *new(V), false
	}
}

// Checks if the target struct has a method of the given name.
func HasMethod[S any](target S, method string) bool {
	ensureValue("HasMethod", target)
	targetValue := reflect.ValueOf(target)
	return targetValue.MethodByName(method).IsValid()
}

// Calls the method of the target struct.
//
// This function panics if the given method doesn't exist.
func CallMethod[S any](target S, method string, args ...any) []any {
	ensureValue("CallMethod", target)
	targetValue := reflect.ValueOf(target)
	fn := targetValue.MethodByName(method)

	if !fn.IsValid() {
		typeName := reflect.TypeOf(target).String()
		panic(fmt.Sprintf("method %s() doesn't exist on %s", method, typeName))
	} else {
		returns := fn.Call(slicex.Map(args, func(arg any, _ int) reflect.Value {
			return reflect.ValueOf(arg)
		}))
		return slicex.Map(returns, func(res reflect.Value, _ int) any {
			return res.Interface()
		})
	}
}

// Creates a map based on the given struct.
//
// This function only collects the exported fields.
func ToMap[V any](s any) map[string]V {
	m := make(map[string]V)
	originalValue := reflect.ValueOf(ensureValue("ToMap", s))
	numField := originalValue.NumField()

	for i := 0; i < numField; i++ {
		fieldType := originalValue.Type().Field(i)

		if fieldType.IsExported() {
			fieldValue := originalValue.Field(i)
			m[fieldType.Name] = fieldValue.Interface().(V)
		}
	}

	return m
}

// Creates a struct based on the given map.
//
// This function only supports the exported fields.
func FromMap[V any, S any](m map[string]V) S {
	var s S
	var oldValue = reflect.ValueOf(s)
	var ptrValue reflect.Value
	var isPtr bool

	if oldValue.Kind() == reflect.Pointer {
		ptrValue = reflect.New(oldValue.Type().Elem())
		isPtr = true
	} else {
		ptrValue = reflect.ValueOf(new(S))
	}

	for key, value := range m {
		field := ptrValue.Elem().FieldByName(key)

		if field.CanSet() {
			field.Set(reflect.ValueOf(value))
		}
	}

	if isPtr {
		return ptrValue.Interface().(S)
	} else {
		return ptrValue.Elem().Interface().(S)
	}
}
