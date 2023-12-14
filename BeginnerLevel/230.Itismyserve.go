package beginnerlevel
// https://www.codechef.com/problems/MYSERVE
import  (
  "fmt"

)

func main () {  
  var t, p, q int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&p, &q);
    sum := p + q; 
    mod := sum % 4 ;
    if( mod == 0 || mod == 1) {
      fmt.Println("Alice");
    } else {
      fmt.Println("Bob");
    }
    t--;
  }

}
