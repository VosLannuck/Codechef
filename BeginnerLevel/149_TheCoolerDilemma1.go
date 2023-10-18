// https://www.codechef.com/problems/WATERCOOLER1

package beginnerlevel
import "fmt"

func main() {

  var t,x,m,y int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&x, &y, &m);
    var totalCosts = m * x ; 
    if(totalCosts >= y) {
      fmt.Println("NO");
    } else {
      fmt.Println("YES");
    }
    t--;
  }

}


