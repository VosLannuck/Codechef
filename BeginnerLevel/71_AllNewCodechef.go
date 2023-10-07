package beginnerlevel
// https://www.codechef.com/problems/NEWCC
import "fmt"

func main() {
  
  var x, y int;
  fmt.Scan(&x, &y);
  if(x == y) {
    fmt.Println("Same");
  } else if (x > y) {
    fmt.Println("New");
  } else {
    fmt.Println("Old");
  }

}

