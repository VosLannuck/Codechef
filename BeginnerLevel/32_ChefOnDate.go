package beginnerlevel
// https://www.codechef.com/problems/CHEFONDATE
import "fmt"

func main() {
  var test int;

  var x, y int;

  fmt.Scan(&test);
  for test > 0 {
    fmt.Scan(&x, &y);
    if(x >= y) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
    test --;
  }
}
