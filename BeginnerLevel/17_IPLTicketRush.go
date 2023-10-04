// https://www.codechef.com/problems/IPLTRSH
package beginnerlevel
import "fmt"

func main() {
  var test int;
  
  fmt.Scan(&test);

  var n,m int ;

  for test > 0 {
    fmt.Scan(&n, &m);
    var result int = n - m;
    if(result < 0) {
      fmt.Println(0);
    } else {
      fmt.Println(result);
    }
    test--;

  }

}
