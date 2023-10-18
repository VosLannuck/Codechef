// https://www.codechef.com/problems/NETFLIX

package beginnerlevel

import "fmt"
func main() {
  var t,a,b,c,x int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&a, &b , &c , &x ) ;
    if(a + b >= x || b +c >= x || c + a >= x) {
      fmt.Println("Yes");
    } else {
      fmt.Println("NO");
    }
    t--;
  }

}
