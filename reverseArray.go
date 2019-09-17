package main
import "fmt"


func main(){
    var a [5]int
    fmt.Println("enter array elements")
    for i:=0;i<len(a);i++ {
        fmt.Scanln(&a[i])
    }
    fmt.Println(a)
    fmt.Println("print reverse array")
    reverseArray(a)
}
func reverseArray(arr [5]int){
    var arrRev [5]int
    j := 0
    for i:=len(arr)-1;i>=0;i--{
    arrRev[i]=arr[j]
    j=j+1   
    }
    fmt.Println(arrRev)
}