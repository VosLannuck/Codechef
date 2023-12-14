package beginnerlevel
// https://www.codechef.com/problems/BRUNCH
import "fmt"
func main() {

  var t, x, y int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    var maxNeighbor int = x / y;
    if(maxNeighbor > 20 ) {
      fmt.Println(20);
    } else {
      fmt.Println(maxNeighbor);
    }
    t--;
  }
  

}
