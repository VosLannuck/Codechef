package beginnerlevel
import "fmt"
// https://www.codechef.com/problems/TAXES
func main() {
  
  var test int;
  fmt.Scan(&test);
  var x int;
  for test > 0{
    fmt.Scan(&x);
    if(x > 100) {
      fmt.Println(x - 10);
    }else {
      fmt.Println(x);
    }
    test --;
  }
  
}
