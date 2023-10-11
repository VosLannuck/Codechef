// https://www.codechef.com/problems/SCALENE
package beginnerlevel
import "fmt"

func main() {

  var t,a,b,c int;

  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&a,&b,&c);
    if(a != b && b != c && c != a) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
    t--;
  }
}
