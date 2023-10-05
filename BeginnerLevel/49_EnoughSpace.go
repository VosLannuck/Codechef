package beginnerlevel
import "fmt"
// https://www.codechef.com/problems/ENSPACE
func main() {

  var t,n,x,y int;
  fmt.Scan(&t);

  for t > 0 {

    fmt.Scan(&n, &x, &y);
    var total = x + y * 2;
    if(total <= n) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
    t--;
  }

}
