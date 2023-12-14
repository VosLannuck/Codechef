package beginnerLevel
// https://www.codechef.com/problems/WTRMIXING
import "fmt"

func main() {
  var t, a, b, x , y int;
  fmt.Scan(&t);
  
  for t > 0 {
    fmt.Scan(&a, &b, &x, &y);

    var difference int = b - a;
    if(difference == 0) {
      fmt.Println("YES"); 
    } else if (difference > 0) {

      if(difference <= x ) {
        fmt.Println("YES")
      }else {
        fmt.Println("NO")
      }

    } else if(difference < 0) {
      if( y >= difference * -1) {
        fmt.Println("YES");
      }else {
        fmt.Println("NO");
      }
    }
    t--;
  } 
}
