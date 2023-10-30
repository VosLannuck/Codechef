package beginnerlevel
// https://www.codechef.com/problems/MONOPOLY2
import "fmt"

func main () {

  var t, p,q,r,s int;
  fmt.Scan(&t);

  for t > 0  {
    fmt.Scan(&p, &q, &r, &s);

    if(p > q + r + s) {
      fmt.Println("YES");
    } else if(q > p + r + s) {
      fmt.Println("YES");
    } else  if(r > p + q + s) {
      fmt.Println("YES");
    } else if(s > p + q+ r) {
      fmt.Println("YES");
    } else {
      fmt.Println("No");
    }

    t--;
  }

}

