package beginnerlevel
// https://www.codechef.com/problems/MAXIMUMSUBS

import "fmt"

func main() {

  var t, x int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x);
    fmt.Println( (x * 60 ) / 30 );
    t--;
  }

}
