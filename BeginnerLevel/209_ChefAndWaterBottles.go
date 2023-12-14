package beginnerlevel
// https://www.codechef.com/problems/CHEFBOTTLE
import "fmt"
func main() {

  var t, n , x , k int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n, &x, &k);
    var totalLiters int = n * x;

    if(totalLiters > k) {
      fmt.Println(k / x);
    } else {
      fmt.Println(n);
    }
      

    t--;
  }

}
