//https://www.codechef.com/problems/MVR

package beginnerlevel
import "fmt"

func main() {
  var test, a, b, x, y int ;
  fmt.Scan(&test);

  fmt.Scan(&a, &b, &x, &y);
  var messi int = a * 2 + b;
  var ronaldo int = x * 2 + y;
  if(messi == ronaldo) {
    fmt.Println("Equal");
  }else if(messi > ronaldo) {
    fmt.Println("Messi");
  }else {
    fmt.Println("Ronaldo");
  }

}

