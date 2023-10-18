// https://www.codechef.com/problems/SUBSCRIBE_

package beginnerlevel

import "fmt"
import "math"
func main() {

  var t int ;
  var n,x float64;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &x);
    var totalGroups float64;
    if(n <= 6) {
      totalGroups = 1;
    } else {
      totalGroups = math.Ceil(n / 6);
    }
    fmt.Println(totalGroups * x);
    t--;
  }

}
