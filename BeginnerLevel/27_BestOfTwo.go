// https://www.codechef.com/problems/BESTOFTWO
package beginnerlevel

import "fmt"

func main() {
  var test int;
  var x, y int;
  
  fmt.Scan(&test);
  for test > 0 {
    fmt.Scan(&x, &y);
    if(x > y) {
      fmt.Println(x);
    }else {
      fmt.Println(y);
    }
    test --;
  }
}
