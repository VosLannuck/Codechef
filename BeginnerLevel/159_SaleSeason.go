package beginnerlevel
 // https://www.codechef.com/problems/SALESEASON
import "fmt"

func main () {

  var x , t int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&x);
    if(x <= 100) {
      fmt.Println(x);
    } else if (x >100 && x <=1000) {
      fmt.Println(x -  25 );
    } else if (x > 1000 && x <= 5000) {
      fmt.Println(x - 100 );
    } else {
      fmt.Println(x - 500 );
    }
    t--;
  }

  

}

