package maduce

import (
	"reflect"
)

// Collection is a list of what ever type you can think of, it is even possible
// to have different types contained within the same collection.
//
// NOTE: Strive to keep a single type in a collection otherwise you may
// experience panics if your filter, map or reduce functions isn't type agnostic.
type Collection []interface{}

// From takes a slice of some type and turns it into a collection.
func From(slice interface{}) Collection {
	s := reflect.ValueOf(slice)
	collection := make(Collection, s.Len())
	c := reflect.ValueOf(collection)

	for i := range collection {
		c.Index(i).Set(s.Index(i))
	}

	return collection
}

// Filter takes a function which evaluates whether an item should be in the
// collection. Items will stay in the collection if the filter function returns
// true.
//
// The following function signatures are supported for filtering:
//	func(item <Type>) bool
//	func(item <Type>, index int) bool
//
// A common way to delete something from a collection is to simply filter based
// on index or some other uniq identifier.
//	// Delete integer with index 1
//	result := maduce.Collection{1, 2, 3}.Filter(func (_, index int) bool { return index != 1})
func (c Collection) Filter(filter interface{}) Collection {
	var result Collection

	switch t := filter.(type) {
	case func(int) bool:
		for _, item := range c {
			if t(item.(int)) {
				result = append(result, item)
			}
		}
	case func(int, int) bool:
		for i, item := range c {
			if t(item.(int), i) {
				result = append(result, item)
			}
		}
	default:
		fn := reflect.ValueOf(filter)
		for i, item := range c {
			args := []reflect.Value{reflect.ValueOf(item), reflect.ValueOf(i)}

			output := fn.Call(args[:fn.Type().NumIn()])

			if output[0].Interface().(bool) {
				result = append(result, item)
			}
		}
	}

	return result
}

// Reduce takes a receiver and a reducer function.
//
// The following function signatures are supported for reducers:
//	func(item <Type>, receiver <Type>) <Type>
//	func(item <Type>, receiver <Type>, index int) <Type>
func (c Collection) Reduce(receiver, reducer interface{}) {
	fn := reflect.ValueOf(reducer)
	r := reflect.ValueOf(receiver)

	for i, item := range c {
		args := []reflect.Value{reflect.ValueOf(item), r.Elem(), reflect.ValueOf(i)}
		output := fn.Call(args[:fn.Type().NumIn()])
		r.Elem().Set(output[0])
	}
}

// Map takes a mapper function which transforms or mutates items and returns
// them in a new collection.
//
// The following function signatures are supported for mappers
//	func(item <Type>) <Type>
//	func(item <Type>, index int) <Type>
func (c Collection) Map(mapper interface{}) Collection {
	result := make(Collection, len(c))

	switch t := mapper.(type) {
	case func(int) int:
		for i, item := range c {
			result[i] = t(item.(int))
		}
	case func(int, int) int:
		for i, item := range c {
			result[i] = t(item.(int), i)
		}
	default:
		fn := reflect.ValueOf(mapper)
		r := reflect.ValueOf(result)

		for i, item := range c {
			args := []reflect.Value{reflect.ValueOf(item), reflect.ValueOf(i)}
			r.Index(i).Set(fn.Call(args[:fn.Type().NumIn()])[0])
		}
	}

	return result
}
