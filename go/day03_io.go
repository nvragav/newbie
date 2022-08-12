package main
import  ( "fmt" 
           "time"
        )
/*
Output:
-------
$ go run day03_io.go 
Enter your YOB: 1985
Your age is  37
Enter the nos to add in array
1
2
3
4
5
--- Printing the array ---
1,2,3,4,5,
Length of slice:  1
[slice]
Enter no of strings to add : 2
is
good
Length of slice: 3
[slice is good]
--- Printing structures ---
{niju sfo 645121}
{aira  900091}
--- Printing maps ---
map[3:lion 1:dog 2:cat]
--- Traversing Map ---
1 dog
2 cat
3 lion
tiger
false

$
*/
type person  struct {
        name string
        city string
        zipcode int
}
func main() {
        t := time.Now()
        var year int
        var no int
        //NOTE: size must be given for array in case of values are not initialized.
       // Another method: a:= [...]int{1,2,3}
        a := [5]int{}
        var s string
        
        fmt.Printf("Enter your YOB: ")
        fmt.Scanln(&year)
        fmt.Println("Your age is ",t.Year()-year)
        //Arrays
        fmt.Println("Enter the nos to add in array") 
        for i := 0; i < 5 ; i++ {
                fmt.Scanln(&a[i])
        } 
        fmt.Println("--- Printing the array ---")
        for _,b := range a {
                fmt.Printf("%d,",b)
                
        }
        fmt.Println()

        // Slices 
        ss := []string{}
        
        ss = append(ss,"slice")
        fmt.Println("Length of slice: ", len(ss))
        fmt.Println(ss)
        fmt.Printf("Enter no of strings to add : ")
        fmt.Scanln(&no)
         for j := 0; j < no; j++ {
        fmt.Scanln(&s)
        ss = append(ss,s)
        }
        fmt.Println("Length of slice:", len(ss))
        fmt.Println(ss)

        //Structures
        //initialization with order
        a1 := person {"niju","sfo",645121}
        //initialization with keys
        a2 := person{name:"aira",zipcode : 900091}
        fmt.Println("--- Printing structures ---")
        fmt.Println(a1)
        fmt.Println(a2)

        //Maps
        m := map[int]string {1:"dog",2:"cat",3:"lion"}
        fmt.Println("--- Printing maps ---")
        fmt.Println(m)
        fmt.Println("--- Traversing Map ---")
        for id,p := range m {
                fmt.Println(id,p)
        }
        m[4] = "tiger"
        //Printing using keys 
        fmt.Println(m[4])
        //Checking whether the key is present or not
        p1,r1 :=  m[5]
        fmt.Println(r1)
        fmt.Println(p1)


}       
