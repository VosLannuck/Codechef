package beginnerlevel
// https://www.codechef.com/problems/EXAMCHEF
import("fmt")

func main() {
  
  var t, x, y, z int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y, &z);
    allStudents := x * y;
    halfStudents := allStudents / 2;

    if (z >  halfStudents) {
      fmt.Println("YES");
    } else {
      fmt.Println("NO");
    }

    t--;
  }
  
  
}
