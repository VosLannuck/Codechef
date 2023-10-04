package beginnerlevel;
//https://www.codechef.com/problems/FIT;
import "fmt";

func main() {
  var test int;
  fmt.Scan(&test);
  var xKm int;
  for test > 0 {
    fmt.Scan(&xKm);
    xKm *= 2;
    fmt.Println(xKm * 5);
    test--;
  }

}





