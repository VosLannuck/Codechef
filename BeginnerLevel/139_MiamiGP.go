// https://www.codechef.com/problems/F1RULE
package beginnerlevel
import "fmt"

func main() {

  var t int
  var x,y float64;
  fmt.Scan(&t);
  for t > 0 {

    fmt.Scan(&x, &y);
    var _107Rule = x * 107/100;

    if( y > _107Rule) {
      fmt.Println("NO");
    } else {
      fmt.Println("YES");
    }
    t--;

  }




}
