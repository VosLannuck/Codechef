package beginnerlevel
// https://www.codechef.com/problems/SNDMAX
import "fmt"
func main() {
  var test int;
  fmt.Scan(&test);
  for test > 0 {
    
    var a, b, c int;
    fmt.Scan(&a, &b, &c);

    if(a > b && a < c ) {
      fmt.Println(a);
    } else if(a > c && a < b) {
      fmt.Println(a);
    } else if(c > a && c < b) {
      fmt.Println(c);
    } else if(c > b && c < a) {
      fmt.Println(c);
    } else if(b > a && b < c) {
      fmt.Println(b);
    } else if(b > c && b < a) {
      fmt.Println(b);
    }
    
    test --;
  }
}

