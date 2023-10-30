package beginnerlevel
import "fmt"
//https://www.codechef.com/problems/BIRDFARM

func main() {
    
  var t,x,y,z int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&x, &y, &z);
    if(z%x == 0 && z % y == 0 ) {
      fmt.Println("Any");
    } else if (z % x == 0 && z % y != 0) {
      fmt.Println("Chicken");
    } else if (z % x != 0 && z % y == 0 ) {
      fmt.Println("Duck");
    } else if ( z % x != 0 && z % y != 0) {
      fmt.Println("NONE");
    }
    t--;
  }

}
