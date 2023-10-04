// https://www.codechef.com/problems/REACHTARGET
package beginnerlevel

import "fmt"

func main() {
  var test int;
  var y, x int;

  fmt.Scan(&test);

  for test > 0 {
    fmt.Scan(&y, &x);
    fmt.Println(y - x);
    test--;
  }
}
