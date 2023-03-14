package go_slice_iterator

func GetIterator[T any](values []T) *Itr[T] {
	return &Itr[T]{
		Idx:    0,
		Values: values,
	}
}

type Iterator[T any] interface {
	HasNext() bool
	GetNext() *T
}

type Itr[T any] struct {
	Idx    int
	Values []T
}

var _ Iterator[any] = (*Itr[any])(nil)

func (itr *Itr[T]) GetNext() *T {
	if itr.HasNext() {
		var value *T
		value = &itr.Values[itr.Idx]
		itr.Idx++
		return value
	}
	return nil
}

func (itr *Itr[T]) HasNext() bool {
	if itr.Idx < len(itr.Values) {
		return true
	}
	return false
}
