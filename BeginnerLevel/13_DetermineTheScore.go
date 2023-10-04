// https://www.codechef.com/problems/DETSCORE
package beginnerlevel

import "fmt"

func main() {
  
  var test int;
  const TESTCASES int = 10;
  fmt.Scan(&test);
  var x, n int; 
  for test > 0 {
    fmt.Scan(&x, &n);
    x /= TESTCASES;
    fmt.Println(x * n);
    test --;
  }

}
