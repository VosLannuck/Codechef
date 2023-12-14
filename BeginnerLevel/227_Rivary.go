package beginnerlevel
import(
  "fmt"
)
// https://www.codechef.com/problems/CPRIVAL
func main() {
  var r1, r2, d1, d2 int;
    fmt.Scan(&r1, &r2 , &d1, &d2);
    if(r1 + d1 >= r2 + d2) {
      fmt.Println("Dominater");
    } else {
      fmt.Println("Everule");
    }
}
