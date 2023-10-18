package beginnerlevel

import "fmt"

func main () {
  

  var t , x, n int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &x);
    fmt.Println(n / (x * 3 ));
    t--;
  }
}

