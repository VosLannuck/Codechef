package begineerLevel

import ("fmt"
"math")
// https://www.codechef.com/problems/EXPENSES
func main() {
  var t int;
  var n, x float64;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&n, &x);
    result := math.Pow(2, x-n);

    fmt.Println(int(result));
    t--;
  }


}
