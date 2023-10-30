package beginnerlevel
// https://www.codechef.com/problems/SUGARCANE
import "fmt"
func main() {

  var t, n int; 
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n);
    var totalIncome int = 50 * n;
    fmt.Println(totalIncome - (20 * totalIncome / 100) - (20 * totalIncome / 100) - (30 * totalIncome / 100));
    t--;
  }

}
