// https://www.codechef.com/problems/CRICMATCH

package beginnerlevel

import "fmt"

func main() {

  var t, n, m int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&n, &m);
    var potentialRuns int = 6 * 6 * m ;
    if(potentialRuns >= n) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");

    }
    t--;
  }


}
