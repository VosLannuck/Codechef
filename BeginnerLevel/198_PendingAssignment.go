package beginnerlevel
// https://www.codechef.com/problems/ASSIGNMNT
import "fmt"
func main () {

  var t, x , y, z int;
  fmt.Scan(&t);

  for t > 0 {

    fmt.Scan(&x, &y, &z);
    var totalHours int= ( x * y )  ;
    if(totalHours <= ((z * 24 ) * 60)) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }

    t--;
  }


}
