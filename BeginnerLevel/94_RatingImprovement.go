package beginnerlevel
import "fmt"
// https://www.codechef.com/problems/ADVANCE
func main() {
  
  var t,x,y int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&x, &y);
    if(y >= x && y <= x + 200) {
      fmt.Println("Yes");
    } else {
      fmt.Println("No");
    }
    t--;
  }
}
