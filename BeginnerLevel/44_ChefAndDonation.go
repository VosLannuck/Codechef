package beginnerlevel
// https://www.codechef.com/problems/DNATION
import "fmt"

func main() {
  var test int;
  fmt.Scan(&test);
  for test > 0 {
    var x, y int ;
    fmt.Scan(&x, &y);
    fmt.Println(y - x);
    test --;
  }
}
