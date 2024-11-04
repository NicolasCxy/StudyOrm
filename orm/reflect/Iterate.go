package reflect

import "reflect"

func IterateArrOrSlice(entity any) ([]any, error) {
	val := reflect.ValueOf(entity)
	res := make([]any, 0, val.Len())
	for i := 0; i < val.Len(); i++ {
		els := val.Index(i)
		res = append(res, els.Interface())
	}

	return res, nil
}

func IterateMap(entity any) ([]any, []any, error) {
	val := reflect.ValueOf(entity)
	//mapKeys := val.MapKeys()
	//keys := make([]any, 0, len(mapKeys))
	//vals := make([]any, 0, len(mapKeys))
	//for i := 0; i < len(mapKeys); i++ {
	//	key := mapKeys[i]
	//	value := val.MapIndex(key)
	//	keys = append(keys, key.Interface())
	//	vals = append(vals, value.Interface())
	//}

	keys := make([]any, 0, val.Len())
	vals := make([]any, 0, val.Len())
	mapRange := val.MapRange()
	for mapRange.Next() {
		keys = append(keys, mapRange.Key().Interface())
		vals = append(vals, mapRange.Value().Interface())
	}

	return keys, vals, nil
}
