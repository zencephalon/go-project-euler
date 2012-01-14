package util

func Divs (a, b int) bool {
   return a % b == 0
}

func PrintIntArray (a []int) {
   for _, v := range a {
      println(v)
   }
}

func Max(a []int) int {
   m := a[0]
   for _, v := range a {
      if v > m {
         m = v
      }
   }
   return m
}
