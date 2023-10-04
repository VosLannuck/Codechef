package beginnerlevel

//https://www.codechef.com/problems/LUDO
import "fmt";
func main() {
  var test int;
  fmt.Scan(&test);
  for test > 0 {
    var inputX int;
    fmt.Scan(&inputX);
    if(inputX != 6) {
      fmt.Println("NO");
    } else {
      fmt.Println("YES")
    }
    test --;
  }

}
