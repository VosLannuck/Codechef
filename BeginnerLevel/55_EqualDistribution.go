// https://www.codechef.com/problems/EQUALDIST
package beginnerlevel

import "fmt"

func main() {

  var t,a,b int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&a, &b);
    if( (a + b) % 2 == 0) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
    t--;
  }
}
