package beginnerlevel
// https://www.codechef.com/problems/CRICUP
import (
  "fmt"
  "math"
)
func main() {
  var t int;
  var x,y,d float64;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y, &d);
    if (math.Abs(x - y) <= d) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
    t--;
  }
}

