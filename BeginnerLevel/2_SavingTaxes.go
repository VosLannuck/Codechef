package beginnerlevel

import "fmt"
import "math"

func main() {
  
  var t int;
  var a,b  float64;
  fmt.Scan(&t);

  for i := 0 ; i < t; i++ {
      fmt.Scan(&a, &b);
      fmt.Println(math.Abs( a - b));
  }
  
}
