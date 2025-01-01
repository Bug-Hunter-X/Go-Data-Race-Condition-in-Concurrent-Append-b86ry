# Go Data Race Condition in Concurrent Append

This repository demonstrates a data race condition that can occur in Go when multiple goroutines concurrently append to a shared slice. The program attempts to populate a slice with numbers from 0 to 999 using goroutines, but due to the non-atomic nature of append, the final length of the slice might be less than 1000.

## How to Reproduce

1. Clone this repository.
2. Run `go run bug.go`.
3. Observe that the output might not always be 1000, indicating a data race.

## Solution

The solution involves using a `sync.Mutex` to protect the `append` operation. This ensures that only one goroutine can access the slice at a time, preventing the race condition.

## License

MIT