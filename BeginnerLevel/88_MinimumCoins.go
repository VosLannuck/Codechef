// https://www.codechef.com/problems/MINCOINSREQ
package beginnerlevel

import "fmt"
func main() {
  
  var t, x int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x);
    fmt.Println(x % 10);
    t--;
  }

}
