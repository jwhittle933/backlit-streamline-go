// Package box (internal) dynamically parses ISO/IEC BMFF boxes/atoms
package box

import "reflect"

func Tags(i interface{}) []reflect.StructTag {
	t := reflect.TypeOf(i)
	numFields := t.NumField()

	tags := make([]reflect.StructTag, 0)

	for i := 0; i < numFields; i++ {
		tag := t.Field(i).Tag
		tags = append(tags, tag)
	}

	return tags
}
