package main
import  ( 
        "fmt"
        "bufio"
        "os"
)

func main() {
        var fName string
        var lName string
        var fullName string
        input := bufio.NewReader(os.Stdin)        
        //Reading multiple strings that has been separated by spaces
        fmt.Printf("1. Enter your full name : ")
        fmt.Scanln(&fName,&lName) 
        fmt.Printf("\tFirstName : %s LastName : %s\n",fName,lName)       
        
        //Reading full line that has space separated strings
        fmt.Printf("\n2. Enter your full name : ")
        fullName,_  = input.ReadString('\n')
        fmt.Println("\tYour full name : ",fullName)

}
