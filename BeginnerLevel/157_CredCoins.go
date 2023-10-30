// https://www.codechef.com/problems/CREDCOINS
package beginnerlevel

import "fmt"
func main() {
  
  var t, x, y int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    fmt.Println( ( x * y)  / 100);
    t--;
  } 
  
}
