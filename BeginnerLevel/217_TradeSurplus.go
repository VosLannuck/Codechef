package beginnerLevel
// https://www.codechef.com/problems/SURPLUS
import "fmt"

func main() {
  var t, a1,a2,b1,b2 int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&a1, &a2, &b1, &b2);
    var netExportA int = a1 - a2;
    var netExportB int = b1 - b2;

    var combined int = netExportA + netExportB;
    if( (combined * -1 ) >0) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }

    t--;
  } 

}
