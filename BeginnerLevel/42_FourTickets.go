
package beginnerlevel
// https://www.codechef.com/problems/FOURTICKETS
import "fmt"

func main() {
  var test int;

  fmt.Scan(&test);

  for test > 0 {
    var x int;
    fmt.Scan(&x);
    if(x * 4 <= 1000) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
    test--;
  }

}
