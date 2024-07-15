package itx

import "github.com/BooleanCat/go-functional/v2/it"

// Drop is a convenience method for chaining [it.Drop] on [Iterator]s.
func (iterator Iterator[V]) Drop(count int) Iterator[V] {
	return Iterator[V](it.Drop(iterator, count))
}

// Drop is a convenience method for chaining [it.Drop2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Drop(count int) Iterator2[V, W] {
	return Iterator2[V, W](it.Drop2(iterator, count))
}
