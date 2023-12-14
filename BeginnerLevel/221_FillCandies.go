package beginnerlevel
// https://www.codechef.com/problems/FILLCANDIES
import "fmt"

func main() {
  var t, n, k ,m int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &k, &m);
    numBags := k * m ;
    if(n % numBags == 0) {
      fmt.Println(n / numBags );
      
    }else {
      
      fmt.Println(n / numBags + 1);
    }
    t--;
  }

}
