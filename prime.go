package prime

import "math"
import . "./util"
import "big"

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

func NextPrime(n int) int {
   i := n + 1
   for !Prime(i) {
      i++
   }
   return i
}


func PrimeFactors(n int) []int {
   factors := []int{}

   for i := 2; i < n; i = NextPrime(i) {
      for ; Divs(n, i); n /= i {
         factors = append(factors, i)
      }
   }

   return factors
}

func PrimeFactorsBig(s string) []int {
   n := big.NewInt(0)
   n.SetString(s, 10)
   factors := []int{}

   for i, big_i := 2, big.NewInt(2); n.Cmp(big_i) == 1; i = NextPrime(i) {
      big_i.SetInt64(int64(i))
      for ; BigDivs(n, big_i); n.Quo(n, big_i) {
         factors = append(factors, i)
      }
   }

   return factors
}

func BigDivs(a, b *big.Int) bool {
   r := big.NewInt(0)
   return r.Rem(a, b).Int64() == 0
}
   
