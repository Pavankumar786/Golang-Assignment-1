1)
package main
import (
	"fmt"
	"time"
)

func main() {
c := make(chan int, 5)
a := 20
go myfunc(c, 1)
go myfunc(c, 2)
go myfunc(c, 3)
go myfunc(c, 4)
go myfunc(c, 5)
go generateInts(c, a)
time.Sleep(1 * time.Second)
}
func myfunc(c chan int, id int) {
for i := range c {
fmt.Printf("myfunc %d received value %d\n", id, i)
time.Sleep(1 * time.Second)
}
}

func generateInts(c chan int, a int) {
for i := 0; i < a; i++ {
c <- i
fmt.Printf("Sent value %d\n", i)
time.Sleep(1 * time.Millisecond)
fmt.Println("Goroutines Finished")
}
close(c)
}

Output:
Sent value 0
myfunc 3 received value 0
Sent value 1
myfunc 1 received value 1
Sent value 2
myfunc 2 received value 2
Sent value 3
myfunc 4 received value 3
Sent value 4
myfunc 5 received value 4
Sent value 5
Sent value 6
Sent value 7
Sent value 8
Sent value 9
Goroutines Finished
