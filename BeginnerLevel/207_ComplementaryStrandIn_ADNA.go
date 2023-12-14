package beginnerlevel
// https://www.codechef.com/problems/DNASTRAND
import ("fmt"
        "strings"
)
func main() {
  
  var t, n int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&n);
    var charInp string;  
    var c []string;
    fmt.Scan(&charInp);
    for _, v := range charInp {
      val := string(v);
      if(val == "A") {
        c = append(c,"T");
      } else if(val == "T") {
        c = append(c,"A");
      } else if(val =="C") {
        c = append(c,"G");
      } else  {
        c = append(c,"C");
      }
    }
    fmt.Println(strings.Join(c, ""));
  
    t--;
  }


}
