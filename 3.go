package main

import prime "./prime"
import util "./util"

func main() {
   print(util.Max(prime.PrimeFactorsBig("600851475143")))
}
