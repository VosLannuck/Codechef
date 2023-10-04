package beginnerlevel;
// https://www.codechef.com/problems/TOP10
import "fmt";

func main() {

  var test int;
  fmt.Scan(&test);
  var rank int;
  for i:=0 ; i< test; i++ {
    fmt.Scan(&rank);
    if(rank <=10 ) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
  }
}   
