package beginnerlevel
// https://www.codechef.com/problems/ELECTN
import "fmt"

func main() {
  
  var t, n , x int;
  fmt.Scan(&t); 

  for t > 0 {
  
    fmt.Scan(&n, &x);
    var counter int = 0;
    for n > 0 {
      var i int;
      fmt.Scan(&i);
      if( i >= x) {
        counter ++;
      }
      n--;
    }

    fmt.Println(counter);
    t--;
  }
    

}
