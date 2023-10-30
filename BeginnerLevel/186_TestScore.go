package beginnerlevel
// https://www.codechef.com/problems/CHEFSCORE
import "fmt"
func main() {


  var t, n, x , y int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&n, &x, &y);

    if( ( y % x == 0 && y > x && n *x >=y)) {
      fmt.Println("YES");
    } else {
      if(y % x == 0) {
        fmt.Println("YES");

      } else {

        fmt.Println("NO");
      }
    }




    t--;
  }
}
