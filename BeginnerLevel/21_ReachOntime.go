//https://www.codechef.com/problems/TIMELY
package beginnerlevel

import "fmt"

func main() {
  var test int;
  fmt.Scan(&test);
  var x int;
  for test > 0 {
    fmt.Scan(&x);
    if(x >= 30) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO")
    }
    test --;
  }

}
