// https://www.codechef.com/problems/TRUESCORE
package beginnerlevel

import "fmt"
func main () {
  var t, a,b,c,d int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&a, &b, &c, &d);
    if(c < a || d < b) {
    fmt.Println("IMPOSSIBLE");
    } else {
    fmt.Println("POSSIBLE");
    }
    t--;
  }

}
