package beginnerlevel

import "fmt"
import "math"


func main() {
  
  var t int;
  var n, x float64;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &x);
    var minimumPizza float64 = math.Ceil((n * x) / 4);
    fmt.Println(minimumPizza);
    t--;
  }


}
