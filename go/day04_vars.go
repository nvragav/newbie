package main

import ("fmt"
        )

/*
OUTPUT:
-------
$ go run day04_vars.go 
Local variable:  10
Global variable:  20
---------------------------------------------
Original structure:  {alex MarinaBay 900091}
Original integer:  10
Modified structure:  {joe MarinaBay 900091} true
Modified integer:  99
---------------------------------------------
Original slice: [slice]
Original  map[1:dog 2:cat 3:lion]
Slice Inside function  [slice example1]
Modified slice: [slice example2]
Modified  map[1:dog 2:cat 3:lion 4:tiger]
---------------------------------------------
Source slice :  [101 102 103 104 105]
Extracted slice :  [102 103]
---------------------------------------------
$
*/
type person struct {
        name string
        city string
        zipcode int
        }
//NOTE: g := 20 - Compilation error, expects declaration 
var g int = 20
func printVars() {
        g := 10
        fmt.Println("Local variable: ",g)
}

func modifyVars1(p *person, ip *int) bool {
        p.name = "joe"
        *ip = 99
        return true
}
func modifyVars2(m map[int]string,ss []string, ssp *[]string) {
        m[4] = "tiger"
        //ss[len(ss)]  = "by ref"
        //NOTE: Slice cant be modified 
        ss = append(ss,"example1")
        fmt.Println("Slice Inside function ", ss)
        //NOTE: To reflect the changes in calling function, we have to use pointer variable to update it.
        *ssp = append(*ssp,"example2")
}
func main() {
        var b bool
        a := person {"alex", "MarinaBay", 900091} 
        i := 10
        
        printVars()
        fmt.Println("Global variable: ",g)
        fmt.Println("---------------------------------------------")
        
        fmt.Println("Original structure: ",a)
        fmt.Println("Original integer: ",i)
        b = modifyVars1(&a,&i)
        fmt.Println("Modified structure: ",a,b)
        fmt.Println("Modified integer: ",i)
        fmt.Println("---------------------------------------------")
        
        m := map[int]string {1:"dog",2:"cat",3:"lion"}
        ss := []string{}
        ss = append(ss,"slice")
        fmt.Println("Original slice:",ss)
        fmt.Println("Original ",m)
        modifyVars2(m,ss,&ss)
        fmt.Println("Modified slice:",ss)
        fmt.Println("Modified ",m)
        

        fmt.Println("---------------------------------------------")
        
        si1 := []int{101,102,103,104,105}
        //Extracting the elements excluding the end index
        si2  :=  si1[1:3] 
        fmt.Println("Source slice : ",si1)
        fmt.Println("Extracted slice : ",si2)
        fmt.Println("---------------------------------------------")

}
