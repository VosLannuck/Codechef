// https://www.codechef.com/problems/ODDSUMPAIR

package beginnerlevel

import "fmt"

func main() {

  var t, a, b , c int ;
  fmt.Scan(&t);
  for t> 0 {
    fmt.Scan(&a, &b, &c);
    if( (a + b) % 2 == 1  || (a + c ) % 2 == 1 || (b+c) % 2 == 1) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
    t--;
  } 
}
