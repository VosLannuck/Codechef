// https://www.codechef.com/problems/AGEING
package beginnerlevel
import "fmt"

func main() {

  var test int;
  fmt.Scan(&test);
  var x int;
  for test > 0 {
    fmt.Scan(&x);

    fmt.Println(x - 10);
    test --;
  }
  
}
