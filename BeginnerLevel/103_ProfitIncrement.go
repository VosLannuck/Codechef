// https://www.codechef.com/problems/PROINC
package beginnerlevel
import "fmt"
func main() {

  var t,x ,y int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    var newSellingPrice int = ( x * 10 / 100 );
    fmt.Println(y + newSellingPrice);
    t--;
  }
}
