package beginnerlevel;
// https://www.codechef.com/problems/BIRYANI
import "fmt";

func main() {
  var test int;
  
  fmt.Scan(&test);
  var a,b int;
  for test > 0 {
    fmt.Scan(&a, &b);
    fmt.Println(a * b);

    test--;
  }
  
}
