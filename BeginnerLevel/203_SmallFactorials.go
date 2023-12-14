package beginnerlevel
// https://www.codechef.com/problems/FCTRL2
import (
  "fmt"
  "math/big"
)

func main() {
  
  var t int;
  fmt.Scan(&t);
  for t > 0 {
    var n int64;
    fmt.Scan(&n);
    var big = new(big.Int);
    big.MulRange(1, n);
    fmt.Println(big);
    t--;
  }

}
