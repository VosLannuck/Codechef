// https://www.codechef.com/problems/CANDIVIDE
package beginnerlevel

import "fmt"

func main() {
  var test int;
  fmt.Scan(&test);

  var n int;
  for test > 0 {
    fmt.Scan(&n);
    if(n % 3 != 0) {
      fmt.Println("NO");
    } else {

      fmt.Println("YES");
    }
    test --;
  }
}
