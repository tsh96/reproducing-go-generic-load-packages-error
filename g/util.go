package g

type A[T any] struct{}

func NewA[T any]() *A[T] { return new(A[T]) }

func NewIntA() *A[int] { return NewA[int]() }
