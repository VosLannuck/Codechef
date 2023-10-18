package beginnerlevel
import "fmt"
//  https://www.codechef.com/problems/BOBBANK
func main() {

  var t,w, x, y ,z int;

  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&w, &x, &y, &z);
    var totalDeposited  int = x * z ;
    var totalCharged int = y * z;
    fmt.Println( w + totalDeposited - totalCharged ) ;
    t--;
  }

}
