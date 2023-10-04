// https://www.codechef.com/problems/BTRYHLTH
package beginnerlevel

import "fmt"

func main() {
  
  var test int;
  var x int;
  fmt.Scan(&test);
  for test > 0 {
    fmt.Scan(&x);
    if(x >= 80) {
      fmt.Println("YES");
    }else {
      fmt.Println("NO");
    }

    test --;
  }

}
