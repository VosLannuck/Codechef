// https://www.codechef.com/problems/PUZHUNT
package beginnerlevel

import "fmt"

func main() {
  var n int;
  fmt.Scan(&n);

  if(n < 6 || n > 8) {
    fmt.Println("NO");
  }else {
    fmt.Println("YES");
  }
}


