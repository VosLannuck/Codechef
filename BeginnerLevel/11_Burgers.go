// https://www.codechef.com/problems/BURGERS

package beginnerlevel

import "fmt"

func main() {

  var test int;
  var a, b int;
  fmt.Scan(&test);
  for test > 0 {
    fmt.Scan(&a, &b);
    if (a == b) {
      fmt.Println(a);
    } else if(a < b) {
      fmt.Println(a);
    } else{
      fmt.Println(b);
    }
    test --;
  }

}
