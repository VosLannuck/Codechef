package beginnerlevel

import "fmt"
// https://www.codechef.com/problems/FIZZBUZZ2301

func main() {

  var a, b, c int;

  fmt.Scan(&a, &b , &c ) ;

  if(a > b && a > c) {
    fmt.Println("YES")
  } else {
    fmt.Println("NO")
  }

}


