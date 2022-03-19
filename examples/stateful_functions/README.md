# Stateful function examples using the generic iterator

inspired by Apache Flink

## Example breakdown

### Serial baseline

see [server_baseline.go](./server_baseline.go)

this version server every request in the serial mode

### Multi workers

see [server_workers.go](./server_workers.go)

this version uses multiple goroutines (NumWorkers) to process the requests hence faster than
the serial version; its performance is far less ideal due to every read and write goes through
a mutex.

### Bad stateful function implementation

see [server_bad_stateful.go](./server_bad_stateful.go)

this version tries to solve the mutex issue by introducing a local state to each worker;

each worker stores the processing result (mutating the global state - a "state delta")
in its local state and moves on to process the next request;

it is completely useless as the state deltas are not merged back to the global state;

it does show a dramatic improvement of speed due to the complete elimination of mutex;

### Ok stateful function implementation

see [server_ok_stateful.go](./server_ok_stateful.go)

this version introduces a "merging" solution that periodically merge the local state that
holds the global state delta, back to the global store.

the frequency of merging is parameterized (called the **MergeTicks**) and can be tuned;

it shows if the frequency is high - leading to better state consistency - the performance
is low; and if the frequency is low - leading to potential errors due to state inconsistency -
the performance can be high

in this very limited example:

```text
ticks(30), low freq   6x speedup
ticks(10), high freq  4x speedup
ticks(4), e-high freq 2x speedup
```

## Thought

there are a few variations of the local state and the merger algorithms:

- use append-only logs for the local state and snapshots for the global state
- use tiered merge frequency as certain properties are less important and can be inconsistent to some extent
- partition the global state and snapshot each partition instead of the entire global state
- use append-only logs for both the local state and the global state; introduce a log-view for read-access only
