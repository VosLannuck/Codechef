package beginnerlevel
// https://www.codechef.com/problems/BETDEAL
import "fmt"
func main() {
   
  var t, a ,b int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&a, &b);
    var aNewPrice = 100 - (100 * a / 100 );
    var bNewPrice = 200 - (200 * b / 100);

    if(aNewPrice == bNewPrice ) {
      fmt.Println("BOTH");
    } else  if(aNewPrice > bNewPrice) {
      fmt.Println("SECOND");
    } else {
      fmt.Println("FIRST");
    }
    t--;
  }

}
