// https://www.codechef.com/problems/COURSEREG
package beginnerlevel

import "fmt"

func main() {
  var t, m, n , k int;
  fmt.Scan(&t);
  for t> 0 {
    fmt.Scan(&n, &m, &k);
    m = m - k;
    if(n > m) {

      fmt.Println("no");

    }else {

      fmt.Println("yes");
    }
    t--;

  }

}
