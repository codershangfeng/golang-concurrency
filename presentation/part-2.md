# Part 2: Concurrency Toolkit in Golang

## Goroutine


## `sync`包

- `WaitGroup`

    ```golang
    var wg sync.WaitGroup

    // Add counter in main/parent goroutine
    wg.Add()

    go func() {
        // Defer execute Done() in sub goroutine
        defer wg.Done()
    }

    // Block main/parent goroutine
    wg.Wait()
    ```

- `Mutex` & `RWMutex`

- `cond`

- `once`

- `pool`

- `channel`

- `select`

- `GOMAXPROCS`