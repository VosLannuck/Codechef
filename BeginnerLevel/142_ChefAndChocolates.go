package beginnerlevel
//https://www.codechef.com/problems/CHEFCHOCO 
import "fmt"

func main() {

  var t, c, x ,y int;

  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&c, &x , &y);
    fmt.Println(y * (c -x));
    t--;
  }
}
