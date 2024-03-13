package utils

type Supplier[T any] func() T

type ErrorSupplier[T any] func() (T, error)

type Consumer[T any] func(T)

type Function[T any, R any] func(T) R
