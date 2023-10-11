//https://www.codechef.com/problems/CANDYSTORE

package beginnerlevel

import "fmt"
func main() {
  
  var t, x, y int;

  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    if(y > x) {
      y = y-x;
      fmt.Println(x + (y * 2));
    }else {
      fmt.Println(y);
    }
    t--;
  }

}
