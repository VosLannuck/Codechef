package beginnerlevel
// https://www.codechef.com/problems/SEMCOURSES
import "fmt"
func main() {

  var t,x,y,z int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y, &z);
    fmt.Println(x * y *z);
    t--;
  } 

}
