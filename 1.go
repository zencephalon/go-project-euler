package main

import . "./util"

func main() {
   sum := 0
   for i := 1; i < 1000; i++ {
      if Divs(i, 3) || Divs(i, 5) {
         sum += i
      }
   }
   print(sum)
}
