package beginnerlevel

import "fmt"

func main() {
  var t,x,y int ;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    if(3 * x < 2 * y) {
      fmt.Println(3 *x);
    } else {
      fmt.Println( 2 * y);
    }
    t--;
  }
  
}
