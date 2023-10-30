package beginnerlevel
// https://www.codechef.com/problems/TODOLIST
import "fmt"


func main() {

  var t , n int;
  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&n);
    var d, count int = 0, 0;

    for n > 0 {
      fmt.Scan(&d);
      if(d >= 1000) {
        count ++;
      }
      n--;
    }
    fmt.Println(count)
    t--;
  }
      

}
