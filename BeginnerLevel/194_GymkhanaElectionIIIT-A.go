package beginnerlevel
import "fmt"
//https://www.codechef.com/problems/CS2023_GYM
func main (){


  var t, n ,m int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &m);
    fmt.Println( (m / 2) + 1); // You only need half + 1 to be chosen.
    t--
  }


}
