package beginnerlevel;
import "fmt";


func main() {
  var test int ;
  fmt.Scan(&test);

  var inp  int;
  for i:=0 ; i < test; i++ {
    fmt.Scan(&inp);

    if(inp >= 2000) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
  }
  
}
// https://www.codechef.com/problems/WATERCONS
