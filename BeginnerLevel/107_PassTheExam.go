package beginnerlevel
// https://www.codechef.com/problems/FLOW002
import "fmt"

func main() {
  
  var t,a,b,c int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&a, &b, &c);
    var totalScore int = a + b + c;
    if(totalScore >= 100 && a >= 10 && b >= 10 && c >= 10) {
      fmt.Println("PASS");
    } else {
      fmt.Println("FAIL");
    }
    t--;
  }
}
