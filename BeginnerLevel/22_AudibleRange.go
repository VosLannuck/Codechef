// https://www.codechef.com/problems/AUDIBLE
package beginnerlevel
import "fmt"
func main() {

  var test int;
  var x int;

  fmt.Scan(&test);

  for test > 0 {
    fmt.Scan(&x);
    if(x < 67 || x > 45000) {
      fmt.Println("NO");
    }else {
      fmt.Println("YES");
    }
    test --;
  }

}
