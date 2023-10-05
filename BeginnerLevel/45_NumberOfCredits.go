package beginnerlevel
// https://www.codechef.com/problems/CREDS
import "fmt"


func main() {
  var test int;
  fmt.Scan(&test);
  
  for test > 0 {

    var x,y,z int;
    fmt.Scan(&x, &y, &z);
    fmt.Println(x * 4 + y * 2);
    test --;
  }


}
