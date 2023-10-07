// https://www.codechef.com/problems/MARKSTW
package beginnerlevel

import "fmt"

func main() {
  
  var x,y int;
  fmt.Scan(&x, &y);
  if(x >= y * 2) {
    fmt.Println("Yes");
  }  else {
    fmt.Println("No");
  }

}
