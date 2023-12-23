package beginnerlevel

import "fmt"

// https://www.codechef.com/problems/XJUMP

func main() {
  var t, x, y int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    step := (x / y);
    fmt.Println(step + (x % y));
    t--;
  }
}
