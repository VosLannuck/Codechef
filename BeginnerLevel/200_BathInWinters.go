package beginnerlevel

// 
import "fmt"
// https://www.codechef.com/problems/BATH
func main() {

  var t, x , y int;
  fmt.Scan(&t);
  for t > 0  {
    fmt.Scan(&x,&y);
    total := y * 2;
    fmt.Println(x / total);

    t--;
  }

}
