package main 
import ( 
    "fmt"
    "math/rand"
    "time"
) 
  
func main() { 
fmt.Println(rand.Intn(200)) 
fmt.Println(rand.Intn(200)) 
fmt.Println(rand.Intn(200)) 
fmt.Println() 
fmt.Println(rand.Float64()) 
fmt.Println((rand.Float64() * 12) + 54) 
fmt.Println((rand.Float64() * 815) + 23) 
fmt.Println() 
s1 := rand.NewSource(time.Now().UnixNano()) 
p1 := rand.New(s1) 
fmt.Println(p1.Intn(200)) 
fmt.Println(p1.Intn(200)) 
fmt.Println() 
s2 := rand.NewSource(55) 
p2 := rand.New(s2) 
fmt.Println(p2.Intn(200)) 
fmt.Println(p2.Intn(200)) 
fmt.Println() 
 s3 := rand.NewSource(5) 
 p3 := rand.New(s3) 
fmt.Println(p3.Intn(200)) 
fmt.Println(p3.Intn(200)) 
} 

Output:
81
87
47

0.4377141871869802
59.09564996485519
582.7608043866942

0
128

112
164

26
36
