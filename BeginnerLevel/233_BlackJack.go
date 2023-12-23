package beginnerlevel

import "fmt"

// https://www.codechef.com/problems/BLACKJACK

func main() {
  var t, a, b int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&a, &b);
    var sum int = a + b;
    var residual = 21 - sum;
    if(residual > 10) {
      fmt.Println(-1);
    } else {
      fmt.Println(residual);
    }
    t--;
  }
}
