package itx

import (
	"iter"

	"github.com/BooleanCat/go-functional/v2/it"
)

// Unzip is a convenience method for chaining [it.Unzip] on [Iterator2]s.
func (iterator Iterator2[V, W]) Unzip() (Iterator[V], Iterator[W]) {
	left, right := it.Unzip(iter.Seq2[V, W](iterator))
	return Iterator[V](left), Iterator[W](right)
}

// Left is a convenience method that unzips an [Iterator2] and returns the left
// iterator, closing the right iterator.
func (iterator Iterator2[V, W]) Left() Iterator[V] {
	return Iterator[V](it.Left(iter.Seq2[V, W](iterator)))
}

// Right is a convenience method that unzips an [Iterator2] and returns the
// right iterator, closing the left iterator.
func (iterator Iterator2[V, W]) Right() Iterator[W] {
	return Iterator[W](it.Right(iter.Seq2[V, W](iterator)))
}
