package beginnerlevel
// https://www.codechef.com/problems/KITCHENSPICE
import "fmt"

func main() {

  var t, x int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&x);
    if(x < 4) {
      fmt.Println("MILD");
    } else if(x >=4 && x < 7) {
      fmt.Println("MEDIUM");
    }else {
      fmt.Println("HOT");
    }
    t--;
  }

}
