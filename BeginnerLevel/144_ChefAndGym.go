// https://www.codechef.com/problems/CGYM

package beginnerlevel

import "fmt"

func main() {
  
  var t,x,y,z int;

  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y, &z);
    if(x + y <= z) {
      fmt.Println(2);
    } else if(x <= z) {
      fmt.Println(1);
    } else {
      fmt.Println(0);
    }
    t--;
  }
  
}
