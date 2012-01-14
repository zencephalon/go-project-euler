package prime

import "math"
import . "./util"

func Prime(n int) bool {
   if n == 1 || n % 2 == 0 {
      return false
   }

   root := int(math.Sqrt(float64(n)))

   for i := 3; i < root; i += 2 {
      if Divs(n, i) {
         return false
      }
   }

   return true
}
