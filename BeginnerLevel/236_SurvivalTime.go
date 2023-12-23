package beginnerlevel

import "fmt"

// https://www.codechef.com/problems/FIZZBUZZ23_2
func main() {

  var t, n, x, d int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &x, &d);
    fmt.Println(n / (x * 5) + d);
    t--;
  }
}
