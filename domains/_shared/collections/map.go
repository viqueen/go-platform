package collections

func Map[TYPE interface{}, RETURN interface{}](items []TYPE, transform func(TYPE) RETURN) []RETURN {
	result := make([]RETURN, len(items))
	for index, value := range items {
		result[index] = transform(value)
	}
	return result
}
