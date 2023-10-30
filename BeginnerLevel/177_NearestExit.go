package beginnerlevel
// https://www.codechef.com/problems/NEARESTEXIT
import "fmt"

func main() {


  var t , x int;

  fmt.Scan(&t);
  for t > 0 {

    fmt.Scan(&x);
    if(x <= 50) {
      fmt.Println("LEFT");
    }else {
      fmt.Println("RIGHT");
    }

    t--;
  }

}
