package beginnerlevel
// https://www.codechef.com/problems/HIGHSCORE
import "fmt"

func main() {

  var t, n int;
  fmt.Scan(&t);
  for t > 0 {
    var i = 4;
    var curMax int = 0;
    var inp int = 0;
    fmt.Scan(&n);
    for i > 0 {
      fmt.Scan(&inp);
      if(inp > curMax) {
        curMax = inp;
      }
      i--;
    }

    fmt.Println(curMax);
    t--;
  } 


}
