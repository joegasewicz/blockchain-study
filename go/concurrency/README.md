# Concurrency

## Goroutines
There are 3 main ways to handle Concurrency in Go
1. Wait Groups
2. [Mutexes & Race Conditions]()
3. ...


- Every Go programme runs in a goroutine - the `main` function is a goroutine.
- Goroutines are lightweight green threads managed by the go scheduler.

### Scheduling
If you spawn a new goroutine, the goroutine may never run if there is not sufficient time to complete execution.
```go
func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	go printSomething("This is the first thing to be printed")
	printSomething("This is the second thing to be printed")
}
// Output:
// This is the second thing to be printed
```
But with `time.Sleep`:
```go
func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	go printSomething("This is the first thing to be printed")
	time.Sleep(time.Second)
	printSomething("This is the second thing to be printed")
}
// Output:
// This is the first thing to be printed
// This is the second thing to be printed

```
#### Deadlocks
Deadlocks can be caused by adding too many goroutines to the delta argument of the Add method: `wg.Add(100)`
```
fatal error: all goroutines are asleep - deadlock!

```
Always use a dynamic value e.g. `wg.Add(len(my_slices))`

## Mutexes & Race Conditions
Mutex **(or mutual exclusion)** allows us to deal with race conditions.
- Relatively easy to use
- Allows us to deal with shared resources **(something that can be access by at least 2 goroutines at the same time)**.
- Able to Lock/Unlock a resource.
- Test for race conditions when running code either with a cli cmd or unit tests.

### Race Conditions
Race conditions occur when multiple goroutines try to access shared memory
```
    |--------> [ Goroutine #1 ] -------|
[ main() ]                          [ some data ]
    |--------> [ Goroutine #2 ] -------|
```

### Channels
- Channels let a goroutine share data within another goroutine.
- In effect, channels let goroutines talk to each other.
- Channels can go in both directions, unidirectional & bidirectional.
- You can have as many channels as you require.
- Go's philosophy is "Share memory by communicating, rather than communicating by sharing memory"