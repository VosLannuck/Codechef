package beginnerlevel
// https://www.codechef.com/problems/READPAGES
import "fmt"

func main() {

  var t,x,y,n int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &x, &y);
    var total int = x * y; 
    if(total >= n) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }

    t--;
  }



}
