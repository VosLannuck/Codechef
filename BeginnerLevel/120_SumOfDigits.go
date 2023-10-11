package beginnerlevel
import "fmt"
// https://www.codechef.com/problems/FLOW006
func main() {
    
  var t, n int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n);
    var sum int  = 0;
    for n > 0 {
      sum += n % 10;
      n/=10;
    }
    fmt.Println(sum);
    t--
  }

}
