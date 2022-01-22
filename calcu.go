package main
import "fmt"

func main(){
   fmt.Println("===================")
   fmt.Println("  Go Calculator   ")
   fmt.Println("===================")
   fmt.Println(" [1] Addition ")
   fmt.Println(" [2] Subtraction")
   fmt.Println(" [3] Multiplication")
   fmt.Println(" [4] Divition")
    fmt.Println("===================")
   var ch int
   fmt.Println(" ")
   fmt.Print("[•] Enter your choice :: ")
   fmt.Scanln(&ch)
   if ch==1{
      add()
   }else if ch==2{
      sub()
   }else if ch==3{
      multi()
   }else if ch==4{
     div()
   }   
}

func add(){
   var x int
   var y int
   fmt.Println("ADDITION")
   fmt.Print(" Enter A : ")
   fmt.Scanln(&x)
   fmt.Print(" Enter B : ")
   fmt.Scanln(&y)
   sum := x+y
   fmt.Println(" Result = ",sum)   
}
func sub(){
   var x int
   var y int
   fmt.Println(">> SUBTRACTION")
   fmt.Print(" Enter A : ")
   fmt.Scanln(&x)
   fmt.Print(" Enter B : ")
   fmt.Scanln(&y)
   sum := x-y
   fmt.Println(" Result = ",sum)   
}
func multi(){
   var x int
   var y int
   fmt.Println(">> MULTIPLICATION")
   fmt.Print(" Enter A : ")
   fmt.Scanln(&x)
   fmt.Print(" Enter B : ")
   fmt.Scanln(&y)
   sum := x*y
   fmt.Println(" Result = ",sum)   
}
func div(){
   var x int
   var y int
   fmt.Println(">> DIVITION")
   fmt.Print(" Enter A : ")
   fmt.Scanln(&x)
   fmt.Print(" Enter B : ")
   fmt.Scanln(&y)
   sum := x/y
   fmt.Println(" Result = ",sum)   
}