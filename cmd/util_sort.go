package cmd

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

// sortableGeneric is a struct containing a sort key and a generic value that
// can be sorted so as to preserve the original value. This could also be done
// by implementing the Sort interface on a type alias for interface{}, however,
// we want to be able to return errors if we don't understand the type.
type sortableGeneric struct {
	sortKey reflect.Value
	value   interface{}
}

// NewSortableGeneric returns a sortableGeneric or an error if we don't know
// how to sort by the type of v.
func newSortableGeneric(sortKey reflect.Value, v interface{}) (sortableGeneric, error) {
	if sortKey.Kind() != reflect.Int && sortKey.Kind() != reflect.String && reflectValueToStringer(sortKey) == nil {
		return sortableGeneric{}, fmt.Errorf("cannot sort by unknown field type: %v", sortKey.Kind())
	}
	return sortableGeneric{sortKey, v}, nil
}

type sortableGenericSlice []sortableGeneric

func (s sortableGenericSlice) Len() int { return len(s) }

func (s sortableGenericSlice) Less(i, j int) bool {
	if s[i].sortKey.Kind() != s[j].sortKey.Kind() {
		return false
	}
	switch s[i].sortKey.Kind() {
	case reflect.Int:
		return s[i].sortKey.Int() < s[j].sortKey.Int()
	case reflect.String:
		return cmpStr(s[i].sortKey.String(), s[j].sortKey.String())
	case reflect.Struct:
		s1 := reflectValueToStringer(s[i].sortKey)
		if s1 == nil {
			return false
		}
		s2 := reflectValueToStringer(s[j].sortKey)
		if s2 == nil {
			return false
		}

		return cmpStr(s1.String(), s2.String())
	}

	return false
}

func (s sortableGenericSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s sortableGenericSlice) Values() []interface{} {
	res := make([]interface{}, len(s))
	for i, v := range s {
		res[i] = v.value
	}
	return res
}

func sortByField(s []interface{}, path string, reverse bool) ([]interface{}, error) {
	// Lookup sort key from each obj field
	res := sortableGenericSlice{}
	for _, v := range s {
		sortKey, err := lookupField(v, path)
		if err != nil {
			return nil, err
		}

		g, err := newSortableGeneric(sortKey, v)
		if err != nil {
			return nil, err
		}

		res = append(res, g)
	}

	if reverse {
		sort.Sort(sort.Reverse(res))
	} else {
		sort.Sort(res)
	}

	return res.Values(), nil
}

func cmpStr(i, j string) bool {
	// If strings can be compared as ints, do that first. This is to handle
	// the fact that all Hue Bridge IDs are generally numbers. We want the
	// sort to treat them as if they are real integers.
	if res, err := cmpStrInt(i, j); err == nil {
		return res
	}
	return i < j
}

func cmpStrInt(i, j string) (bool, error) {
	intI, err := strconv.Atoi(i)
	if err != nil {
		return false, err
	}
	intJ, err := strconv.Atoi(j)
	if err != nil {
		return false, err
	}

	return intI < intJ, nil
}

func reflectValueToStringer(v reflect.Value) fmt.Stringer {
	if !v.CanInterface() {
		return nil
	}

	stringer, ok := v.Interface().(fmt.Stringer)
	if !ok {
		return nil
	}

	return stringer
}
