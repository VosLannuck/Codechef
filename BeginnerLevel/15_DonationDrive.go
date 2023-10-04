package beginnerlevel

import "fmt"
// https://www.codechef.com/problems/DONDRIVE
func main() {

  var test int;
  var x, n int;
  fmt.Scan(&test);
  for test > 0 {
    fmt.Scan(&n, &x);
    fmt.Println(n - x);
    test -- ;
  }


}
