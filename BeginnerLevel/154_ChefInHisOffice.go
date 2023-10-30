package beginnerlevel

import "fmt"
// https://www.codechef.com/problems/OFFICE
func main () {

  var t, x , y int;

  fmt.Scan(&t);
  for t > 0 {
    var total int;
    fmt.Scan(&x, &y);
    total+= x * 4 + y;
    fmt.Println(total);
    t--;
  }
  
}
