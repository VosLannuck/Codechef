package beginnerlevel
// https://www.codechef.com/problems/FLOW007
import ( "fmt"
        "strconv")
func main() {

  var t, n int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n);
    var str string = ""
    for n != 0 {
      mod := n%10;
      str += strconv.Itoa(mod);
      n/=10;
    }
    value,_ := strconv.Atoi(str)
    fmt.Println(value);

    t--;
  }

}
