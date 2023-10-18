// https://www.codechef.com/problems/FAVOURITENUM
package beginnerlevel
import "fmt"

func main() {
  
  var t, a int;
  fmt.Scan(&t );
  for t > 0 {
    fmt.Scan(&a);
    if( a % 2 == 0 && a % 7 == 0) {

      fmt.Println("Alice");

    } else if( a % 2 == 1 && a % 9 == 0 ) {

      fmt.Println("BOb");
    } else {
      fmt.Println("Charlie");
    }
    t--;
  }

}
