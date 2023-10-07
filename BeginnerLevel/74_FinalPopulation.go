package beginnerlevel

import "fmt"
// https://www.codechef.com/problems/POPULATION
func main() {
  
  var t,x,y,z int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y, &z);
    fmt.Println(x - y + z);
    t--;
  }

}
