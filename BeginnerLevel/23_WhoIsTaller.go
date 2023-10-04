// https://www.codechef.com/problems/TALLER
package beginnerlevel
import "fmt"
func main() {

  var test int;
  var x, y int;
  fmt.Scan(&test);
  for test > 0 {
    fmt.Scan(&x, &y);
    if(x > y ) {
      fmt.Println("A");
    } else { 
      fmt.Println("B");
    }
    test --;
  }
}
