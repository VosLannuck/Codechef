package beginnerlevel
// https://www.codechef.com/problems/DOMINANT
import "fmt"

func main() {

  var t, na,nb,nc int;
  fmt.Scan(&t);
  for t > 0 {

    fmt.Scan(&na, &nb, &nc);
    if(na > nb + nc) {
      fmt.Println("Yes");
    }else if(nb > na + nc) {
      fmt.Println("Yes");
    }else if(nc > na + nb) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }

    t--;
  }
}
