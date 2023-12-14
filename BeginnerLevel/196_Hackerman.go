package beginnerlevel
// https://www.codechef.com/problems/PRIMEDICE
import "fmt"
func main() {
    
  var t, a, b int;
  fmt.Scan(&t);
  for t > 0 {
    
    fmt.Scan(&a, &b);
    var sum int = a + b;

    if(sum >= 2 && sum < 12 && (sum == 3 || sum == 2 || sum == 5 || sum == 7 || sum == 11)) {
      fmt.Println("ALICE");
    } else {
      fmt.Println("BOB");
    }
    
    t--;
  }
}
