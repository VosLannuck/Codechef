// https://www.codechef.com/problems/KITCHENTIME
package beginnerlevel
import "fmt"
func main() {
  
  var test int;
  fmt.Scan(&test);
  var x, y int ;
  for test > 0{  
    fmt.Scan(&x, &y);
    fmt.Println( y - x);
    test --;
  }

}
