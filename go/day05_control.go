package main
import ( "fmt"
        "math"
        "math/rand"
        )
/*
OUTPUT:
--------
$ go run day05_control.go
Enter the no to check for prime or not(1 to exit) : 5
5 - Prime
Enter the no to check for prime or not(1 to exit) : 99
99 - Not Prime
Enter the no to check for prime or not(1 to exit) : 10000003
10000003 - Not Prime
Enter the no to check for prime or not(1 to exit) : 1
Ready to roll(r/E) : r
2
Ready to roll(r/E) : r
3
Ready to roll(r/E) : r
3
Ready to roll(r/E) : r
5
Ready to roll(r/E) : r
2
Ready to roll(r/E) : r
4
Ready to roll(r/E) : r
1
Ready to roll(r/E) : E
See you, next time
$
*/
func main() {
        no := 0
        flag := false
        ip := "E"
        min := 1
        max := 6
 
       for {
            fmt.Printf("Enter the no to check for prime or not(1 to exit) : ")
            fmt.Scanln(&no)
            if no > 1 {
                for i := 2; i <= int(math.Sqrt(float64(no))); i++ {
                        if no % i == 0 {
                                flag = true
                                break;
                        }       
                }
                if flag == true {
                        fmt.Printf("%d - Not Prime\n",no)
                } else {
                        fmt.Printf("%d - Prime\n",no)
                }
            } else if no == 1 {
                    break
            } else {
                fmt.Println("Invalid input")
            }            
        }
        for {
            fmt.Printf("Ready to roll(r/E) : ")
            fmt.Scanln(&ip)
            if ip == "r" {
                // NOTE: without '+ min', it does create integers with 0. Hence added '+ min'
               fmt.Println(rand.Intn(max-min) + min) 
            } else if ip == "E" {
                    fmt.Println("See you, next time")
                    break
            } else {
                fmt.Println("Invalid input")
            }            
        }       
        
}
