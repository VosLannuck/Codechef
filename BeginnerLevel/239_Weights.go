package begineerLevel
// https://www.codechef.com/problems/WGHTS
import ("fmt")

func main() {
  var t, w, x, y, z int;

  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&w, &x, &y, &z);
    if (w == x + y || w  == x + z || w == y + z || w == x || w == y || w == z || w == x + y + z) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }
    t--;
  }
}
