package beginnerlevel
 import "fmt"
// https://www.codechef.com/problems/DOUBLERENT

func main () {

 
  var t, n int; 
  fmt.Scan(&t);
  
  for t > 0 {
    fmt.Scan(&n);
    fmt.Println(n * 10 );
    t--;
  } 

}
