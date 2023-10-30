package beginnerlevel

import "fmt"

func main() {

  var t, n int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n);

    if(n % 4 == 0 ) {
      fmt.Println(n / 4);
    } else {
      fmt.Println( (n / 4) +1 );
    }
    t--;
  }

}
