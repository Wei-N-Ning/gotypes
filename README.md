# Golang Monadic Types

To stress test Golang's generic type system by implementing the basic monadic types (maybe/option, either, vector, etc.)

## Limitation

- cannot produce a new type parameter in a receiver function (a method)
    - e.g. this won't work: `func (iter Iterator[T]) ChunkSlice(size int) Iterator[[]T]` (`[]T` is a new type parameter)
- type inference sometimes is broken (see the unit tests where I have to add the variable type hint)
- cannot define type alias inside function (will treat the type parameter as index) - it is not a critical issue but
  makes the code cumbersome
- Go doesn't first class tuple and pair types: `Tuple[T...], or Pair[T, P]`
- no variadic type parameter support (or I don't know?)
