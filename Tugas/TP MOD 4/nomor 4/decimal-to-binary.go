package main
import (
   "fmt"
)

func decimalToBinary(angka int){
   var binary []int

   for angka !=0 {
      binary = append(binary, angka%2)
      angka = angka / 2
   }
   if len(binary)==0{
      fmt.Printf("%d\n", 0)
   } else {
      for i:=len(binary)-1; i>=0; i--{
         fmt.Printf("%d", binary[i])
      }
      fmt.Println()
   }
}

func main() {
   decimalToBinary(1)
   decimalToBinary(12)
   decimalToBinary(29)
   decimalToBinary(156)
   decimalToBinary(12345)
}