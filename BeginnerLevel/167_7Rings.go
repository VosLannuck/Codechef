package beginnerlevel
// https://www.codechef.com/problems/SEVENRINGS
import "fmt"
import "strconv"
func main()  {

  var t , n , x int;
  fmt.Scan(&t);
  for t > 0 {
    var isValidPhoneNumber bool = false;

    var sum int = 0;
    fmt.Scan(&n);
    fmt.Scan(&x);
    sum = n * x;
  
    var sumStr string = strconv.Itoa(sum);

    if(sumStr[0] != '0' && len(sumStr) == 5) {
      isValidPhoneNumber = true;
    }
    
    if(isValidPhoneNumber) {
      fmt.Println("YES");
    } else  {
      fmt.Println("NO");
    }


    t--;

  } 




}
