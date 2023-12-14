package datastructure

// https://www.codechef.com/problems/DOLL
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  var t int;
  scanner := bufio.NewReader(os.Stdin)
  fmt.Fscan(scanner, &t);
  for t > 0 {
    var n, k, ans int;
    fmt.Fscan(scanner, &n, &k);
    var i = make([]int, n);
    for n > 0 {
      fmt.Fscan(scanner, &i[n-1]);
      if (i[n-1]> k ) {
        ans += 1;
      }
      n --;
    }

    fmt.Println(ans);

    t--;
  }
}
