package collections

func Partition[TYPE interface{}](items []TYPE, size int) (chunks [][]TYPE) {
	for size < len(items) {
		items, chunks = items[size:], append(chunks, items[0:size:size])
	}
	return append(chunks, items)
}
