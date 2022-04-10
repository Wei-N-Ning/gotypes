# Task Queue

## Summary

The task queue is an abstraction that combines an FIFO queue and concurrent execution.

Its core building materials are:

- a couple of go channels named `fi` and `fo` (standing for first-in and first-out)
- a size that controls the buffer size of both channels
- a context-aware function `f` that takes a `T` and produces `P, error`

Here is an overview of its lifecycle:

Upon creation, the task queue will poll the `fi` channel for incoming values of type `T`;

`Q.Enqueue(T)` will push a value of `T` into `fi`, giving the task queue something to work on; the task queue will
immediately call the context-aware function `f` with the given new value.

The result of this computation (or any error coming out of it), `(P, error)`, are stored in a thin structure
called `TaskOutput[P]`; both the return value and the error are accessible via the public fields `.Value` and `.Err`;

the `TaskOutput[P]` then is pushed into the output channel `fo` for consumption.

The task queue exhibits its concurrency when there are many input values of type `T` in the input channel `fi`; The task
queue will apply `f` to all the input values, creating many future-like computation;

The caller has two way to retrieve the value from such "futures" from the Task Queue:

call `Q.Dequeue()`

- this will block if the first pending computation in the output channel `fo` is still working;
- this returns the plain value representation `(P, error)` instead of `TaskOutput[P]` (the task queue unwraps it).

call `Q.OutputChannel()`

- this will "flatten()" all the "futures" in the output channel `fo`;
- resulting in a go channel of values of type `TaskOutput[P]`;
- it will also block if the first pending computation in the output channel `fo` is still working.

To express this concurrency in plain words:
**While we are (potentially) blocked, waiting for the first output value, we let the task queue work on the rest of the
inputs and make them ready for consumption, so that we are less likely blocked when we read these values.**

The client should call `Q.Seal()` to inform the task queue that there is no more input value to come (i.e., it
causes the input channel to close).

## Context-awareness

It is the developers' responsibility to ensure the encapsulated computation `f` obeys the given context.

More practically, it should poll `ctx.Done()` and exit whenever this becomes true. See the unit tests for examples.

## Failure handling

The task queue will not stop when a certain computation returns an error. It is the developers' responsibility
to cancel the remaining computation gracefully. See the unit tests for examples.

