# Notes

## Concurency

**Executes foo in a new go routine**
```go
go foo()
```
**Sleep for x `time.Duration`**

```go
time.sleep(x)
```

**Changing number of logical processors**
```go
XX
```

**Wait groups**
```go
wg sync.waitGroup
wg.Add(n) // increment number of parties
wg.Done() // decrement
wg.Wait() // block until counter is 0
```

**Mutex/mutual exclusion**
```go
var mutex sync.Mutex;
mutex.Lock()
mutex.Unlock()
```

**Channels**
```go
c:= make(chan c); // create unbuffered/blocking a channel
c <- 3 // send data to a channel
x := <- c // receive data from a channel
<- c // throwing received data can be used as a mode of synchronization

c:= make(chan c, 3); // create unbuffered a channel of capacity 3

for i := range c {
    fmt.PrintLg(i)
}
close(c) // used by the sender to end the loop

select { // First-come first-served
    case a = <- c1:
        fmt.Println(a)
    case b = <- c2:
        fmt.Println(b)
}

select { // select either send or receive
    case a = <- c1:
        fmt.Println("received a")
    case c2 <- b:
        fmt.Println("sent b")
}

// abort Channel
for {
    select {
        case a = <- c1:
            fmt.Println(a)
        case  <- abort:
            break // out of the loop
    }
}

// default case: don't block at all
select { // First-come first-served
    case a = <- c1:
        fmt.Println(a)
    case b = <- c2:
        fmt.Println(b)
    default:
        fmt.Println("nop")
}
```

**Execute a function once even if called multiple times**
```go
func init() {
    fmt.Println("Init")
}

var oneTime sync.Once
oneTime.Do(init) // only 1 go routine will execute the function setup. The other call will block
```

