// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: copy.gen.go.tmpl

package arrowutil

import (
	"fmt"

	"github.com/influxdata/flux/array"
)

// CopyTo will copy the contents of the array into a new array builder.
func CopyTo(b array.Builder, arr array.Interface) {
	switch arr := arr.(type) {

	case *array.Int:
		CopyIntsTo(b.(*array.IntBuilder), arr)

	case *array.Uint:
		CopyUintsTo(b.(*array.UintBuilder), arr)

	case *array.Float:
		CopyFloatsTo(b.(*array.FloatBuilder), arr)

	case *array.Boolean:
		CopyBooleansTo(b.(*array.BooleanBuilder), arr)

	case *array.String:
		CopyStringsTo(b.(*array.StringBuilder), arr)

	default:
		panic(fmt.Errorf("unsupported array data type: %s", arr.DataType()))
	}
}

// CopyByIndexTo will copy the contents of the array at the given indices into a new array builder.
func CopyByIndexTo(b array.Builder, arr array.Interface, indices *array.Int) {
	switch arr := arr.(type) {

	case *array.Int:
		CopyIntsByIndexTo(b.(*array.IntBuilder), arr, indices)

	case *array.Uint:
		CopyUintsByIndexTo(b.(*array.UintBuilder), arr, indices)

	case *array.Float:
		CopyFloatsByIndexTo(b.(*array.FloatBuilder), arr, indices)

	case *array.Boolean:
		CopyBooleansByIndexTo(b.(*array.BooleanBuilder), arr, indices)

	case *array.String:
		CopyStringsByIndexTo(b.(*array.StringBuilder), arr, indices)

	default:
		panic(fmt.Errorf("unsupported array data type: %s", arr.DataType()))
	}
}

// CopyValue will copy an individual value from the array into the builder.
func CopyValue(b array.Builder, arr array.Interface, i int) {
	switch arr := arr.(type) {

	case *array.Int:
		CopyIntValue(b.(*array.IntBuilder), arr, i)

	case *array.Uint:
		CopyUintValue(b.(*array.UintBuilder), arr, i)

	case *array.Float:
		CopyFloatValue(b.(*array.FloatBuilder), arr, i)

	case *array.Boolean:
		CopyBooleanValue(b.(*array.BooleanBuilder), arr, i)

	case *array.String:
		CopyStringValue(b.(*array.StringBuilder), arr, i)

	default:
		panic(fmt.Errorf("unsupported array data type: %s", arr.DataType()))
	}
}

func CopyIntsTo(b *array.IntBuilder, arr *array.Int) {
	b.Reserve(arr.Len())

	for i, n := 0, arr.Len(); i < n; i++ {
		if arr.IsNull(i) {
			b.AppendNull()
			continue
		}
		b.Append(arr.Value(i))
	}
}

func CopyIntsByIndexTo(b *array.IntBuilder, arr *array.Int, indices *array.Int) {
	b.Reserve(indices.Len())

	for i, n := 0, indices.Len(); i < n; i++ {
		offset := int(indices.Value(i))
		if arr.IsNull(offset) {
			b.AppendNull()
			continue
		}
		b.Append(arr.Value(offset))
	}
}

func CopyIntValue(b *array.IntBuilder, arr *array.Int, i int) {
	if arr.IsNull(i) {
		b.AppendNull()
		return
	}
	b.Append(arr.Value(i))
}

func CopyUintsTo(b *array.UintBuilder, arr *array.Uint) {
	b.Reserve(arr.Len())

	for i, n := 0, arr.Len(); i < n; i++ {
		if arr.IsNull(i) {
			b.AppendNull()
			continue
		}
		b.Append(arr.Value(i))
	}
}

func CopyUintsByIndexTo(b *array.UintBuilder, arr *array.Uint, indices *array.Int) {
	b.Reserve(indices.Len())

	for i, n := 0, indices.Len(); i < n; i++ {
		offset := int(indices.Value(i))
		if arr.IsNull(offset) {
			b.AppendNull()
			continue
		}
		b.Append(arr.Value(offset))
	}
}

func CopyUintValue(b *array.UintBuilder, arr *array.Uint, i int) {
	if arr.IsNull(i) {
		b.AppendNull()
		return
	}
	b.Append(arr.Value(i))
}

func CopyFloatsTo(b *array.FloatBuilder, arr *array.Float) {
	b.Reserve(arr.Len())

	for i, n := 0, arr.Len(); i < n; i++ {
		if arr.IsNull(i) {
			b.AppendNull()
			continue
		}
		b.Append(arr.Value(i))
	}
}

func CopyFloatsByIndexTo(b *array.FloatBuilder, arr *array.Float, indices *array.Int) {
	b.Reserve(indices.Len())

	for i, n := 0, indices.Len(); i < n; i++ {
		offset := int(indices.Value(i))
		if arr.IsNull(offset) {
			b.AppendNull()
			continue
		}
		b.Append(arr.Value(offset))
	}
}

func CopyFloatValue(b *array.FloatBuilder, arr *array.Float, i int) {
	if arr.IsNull(i) {
		b.AppendNull()
		return
	}
	b.Append(arr.Value(i))
}

func CopyBooleansTo(b *array.BooleanBuilder, arr *array.Boolean) {
	b.Reserve(arr.Len())

	for i, n := 0, arr.Len(); i < n; i++ {
		if arr.IsNull(i) {
			b.AppendNull()
			continue
		}
		b.Append(arr.Value(i))
	}
}

func CopyBooleansByIndexTo(b *array.BooleanBuilder, arr *array.Boolean, indices *array.Int) {
	b.Reserve(indices.Len())

	for i, n := 0, indices.Len(); i < n; i++ {
		offset := int(indices.Value(i))
		if arr.IsNull(offset) {
			b.AppendNull()
			continue
		}
		b.Append(arr.Value(offset))
	}
}

func CopyBooleanValue(b *array.BooleanBuilder, arr *array.Boolean, i int) {
	if arr.IsNull(i) {
		b.AppendNull()
		return
	}
	b.Append(arr.Value(i))
}

func CopyStringsTo(b *array.StringBuilder, arr *array.String) {
	b.Reserve(arr.Len())

	{
		sz := 0
		for i, n := 0, arr.Len(); i < n; i++ {
			if arr.IsNull(i) {
				continue
			}
			sz += arr.ValueLen(i)
		}
		b.ReserveData(sz)
	}

	for i, n := 0, arr.Len(); i < n; i++ {
		if arr.IsNull(i) {
			b.AppendNull()
			continue
		}
		b.Append(arr.Value(i))
	}
}

func CopyStringsByIndexTo(b *array.StringBuilder, arr *array.String, indices *array.Int) {
	b.Reserve(indices.Len())

	{
		sz := 0
		for i, n := 0, indices.Len(); i < n; i++ {
			offset := int(indices.Value(i))
			if arr.IsNull(offset) {
				continue
			}
			sz += arr.ValueLen(offset)
		}
		b.ReserveData(sz)
	}

	for i, n := 0, indices.Len(); i < n; i++ {
		offset := int(indices.Value(i))
		if arr.IsNull(offset) {
			b.AppendNull()
			continue
		}
		b.Append(arr.Value(offset))
	}
}

func CopyStringValue(b *array.StringBuilder, arr *array.String, i int) {
	if arr.IsNull(i) {
		b.AppendNull()
		return
	}
	b.Append(arr.Value(i))
}
