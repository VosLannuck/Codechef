// https://www.codechef.com/problems/TVDISC
package beginnerlevel
import "fmt"
func main() {
  var t, a , b , c ,d int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&a, &b, &c, &d);
    aFinalPrice := a - (c);
    bFinalPrice := b - (d);

    if(aFinalPrice == bFinalPrice) {
      fmt.Println("any");
    } else if(aFinalPrice > bFinalPrice) {
      fmt.Println("Second");
    } else {
      fmt.Println("First");
    }

    t--;
  }

}
