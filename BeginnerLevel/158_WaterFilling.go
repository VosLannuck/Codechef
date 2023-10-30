// https://www.codechef.com/problems/WATERFILLING
package beginnerlevel
import "fmt"

func main () {

  var t, b0, b1, b2 int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&b0, &b1, &b2);
    if(b0 + b1 + b2 >=2) {
      fmt.Println("Not Now");
    } else {
      fmt.Println("Water filling time")
    }
    t--;
  }

}
