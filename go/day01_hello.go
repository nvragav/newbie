package main
import "fmt"

func main() {
        myName := "nvragav"
        // This will be printed based on the format
        address :=  "No:9,\n east main road,\n chennai-32"
        //this will be printed as it is
        address1 := `No:1,
                     east main road, chennai-32`
        fmt.Println(myName)
        fmt.Printf(address)
        fmt.Println(address1)
}


