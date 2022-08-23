package main

import ( "fmt"
        )

type dllNode struct {
        prev *dllNode
        data int
        next *dllNode
}
type student struct {
        name string
        yoj  int
        yog  int
        grade float64
}
func changeVal(p *int) {
        switch(*p) {
            case 10:
                    *p += 50
            break;
            case 60 :
                    *p += 40
            break;
        }
}

func createNewStudent(n string,yj int, yg int, g float64 )(*student) {
        stu := new(student)
        stu.name = n
        stu.yoj = yj
        stu.yog = yg
        stu.grade = g
        
        return(stu)        
}
var root *dllNode
func createDllNode(data int) (*dllNode){
        node := new(dllNode)
        node.data  = data
        node.prev = nil
        node.next = nil
        return(node)
}

func insertDllNode(r *dllNode, n *dllNode)(*dllNode) {

        var t *dllNode
        if r == nil {
                root = n
                r = root
        } else  {
                for t = r ; t.next != nil ; t = t.next {
                } 
                t.next = n
                n.prev = t
       }
        return r
}

func traverseDll(r *dllNode) {

    fmt.Println("--- DLL Traversal ---")
    for t := r ; t != nil ; t = t.next {
            fmt.Println(t.data)
    } 

}
func deleteDllNode(data int) {
        
        var t *dllNode
        for t = root ; t != nil ; t = t.next {
                if t.data == data {
                    fmt.Println("Deleting node : ", t.data)
                    t.prev.next = t.next
                    t.next.prev = t.prev
                    t = nil
                    break
                }
        }
        
}
/*
OUTPUT:
-------
02c101 &{jay 2002 2006 7.5}
Deleting the entry from map
Length of Map:  0
New Value 1 -  60
New Value 2 -  100
--- DLL Traversal ---
5
4
3
Deleting node :  4
--- DLL Traversal ---
5
3

*/
func main() {

       var a int = 10
       var sdb map[string]*student 
        //NOTE: Access violation error will pop up if map is not initialized using 'make'        
        sdb = make(map[string]*student)

        sdb["02c101"] = createNewStudent("jay",2002,2006,7.5)        

        
        for r,stu := range sdb {
                fmt.Println(r,stu)
                //NOTE: Setting nil to indicate that this memory is no longer be used.
                //     Garbage Collector can free the memory
                stu = nil
        }
        fmt.Println("Deleting the entry from map") 
        delete(sdb,"02c101")

        fmt.Println("Length of Map: ", len(sdb))
        changeVal(&a)
        fmt.Println("New Value 1 - ", a)
        changeVal(&a)
        fmt.Println("New Value 2 - ", a)
        
        insertDllNode(root, createDllNode(5))
        insertDllNode(root, createDllNode(4))
        insertDllNode(root, createDllNode(3))         
        traverseDll(root)                
        deleteDllNode(4)
        traverseDll(root)
}
