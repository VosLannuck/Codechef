package beginnerlevel
import "fmt"
  // https://www.codechef.com/problems/AUCTION
func main() {
  var t, a, b, c int;

  fmt.Scan(&t);
  
  for t > 0 {
    fmt.Scan(&a, &b, &c);
    if(a > b && a > c ) {
      fmt.Println("Alice");
    } else if (b > a && b > c) {
      fmt.Println("Bob");
    } else if ( c > a && c > b) {
      fmt.Println("Charlie");
    }
    t--;
  }
}
