package beginnerlevel

import "fmt"
func main() {

  var t, s int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&s);
    fmt.Println(s * 24 * 1000);
    t--;
  }


}
