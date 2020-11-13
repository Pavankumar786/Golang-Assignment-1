8) Error Handling
package main
import "fmt"
import "errors"

//Dividing two numbers
func Divide(a int, b int) (int, errors) {
//cannot divide by '0'
if b == 0 {
return 0, errors.New("cannot divide by Zero!")
} else {
return (a / b), nil
}
}
func main() {

//divide 4 by 0
if result, err : = Divide(4,0); err ! =nil {
fmt.Println("Error occured :",  err)
} else
 {
fmt.Println("4/0 is", result)
 
}

}

Output:

Error occured: cannot divide bt Zero!