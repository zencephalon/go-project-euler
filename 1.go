package main

func divs (a, b int) bool {
   return a % b == 0
}

func main() {
   sum := 0
   for i := 1; i < 1000; i++ {
      if divs(i, 3) || divs(i, 5) {
         sum += i
      }
   }
   print(sum)
}
