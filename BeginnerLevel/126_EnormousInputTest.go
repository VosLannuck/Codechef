package beginnerlevel
// https://www.codechef.com/problems/INTEST
import "fmt"

func main() {

  var n,k, a int ;
  fmt.Scan(&n, &k);
  var counter int = 0 ;
  for n > 0 {

    fmt.Scan(&a);
    if(a % k == 0 ) {
      counter++;
    }
    n--;
  }
  fmt.Println(counter);

}
