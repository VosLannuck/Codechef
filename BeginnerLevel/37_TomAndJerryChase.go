// https://www.codechef.com/problems/JERRYCHASE
package beginnerlevel

import "fmt"

func main() {
  
  var test int;
  var x, y int;
  fmt.Scan(&test);

  for test > 0 {
    fmt.Scan(&x, &y);
    if(y > x) {
      fmt.Println("YES");
    }else {
      fmt.Println("NO");
    }
    test--;
  }

}

