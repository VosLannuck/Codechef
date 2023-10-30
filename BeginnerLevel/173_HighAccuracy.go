package beginnerlevel
// https://www.codechef.com/problems/ACCURACY
import "fmt"

func main () {

  var t,x int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x);
    var mod = x % 3;
    if(mod == 0) {
      fmt.Println(0);
    } else{
      fmt.Println(3 - (x % 3));
    }
    t--;
  }


}
