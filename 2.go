package main

func fibber() func() int {
   f1, f2 := 0, 1
   return func() int {
      f1, f2 = f2, f1 + f2
      return f2
   }
}

func main() {
   sum := 0
   f := fibber()
   f() // gets us to 1
   for i := f(); i < 4000000; _, _, i = f(), f(), f() {
      sum += i
   }
   println(sum)
}
