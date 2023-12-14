package beginnerlevel

import ("fmt"
      "math")
// https://www.codechef.com/problems/FAIRSHARE
func main() {

  var t int;
  var n , k float64;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &k);
    fmt.Println(n - math.Floor(n / (k + 1)) * k);
    t--;
  }

}
