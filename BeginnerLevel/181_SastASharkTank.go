package beginnerlevel
 // https://www.codechef.com/problems/SST
import "fmt"
func main () {

  var t, a,  b int;

  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&a, &b);
    valuationA := a* 100 / 10 ;
    valuationB := b* 100 / 20;




    if(valuationA == valuationB) {
      fmt.Println("ANY");

    }else if( valuationA >  valuationB) {
      fmt.Println("FIRST");

    } else {

      fmt.Println("SECOND");

    }



    t--;
  }

}
