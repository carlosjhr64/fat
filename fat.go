// Float Array Tricks
package fat

import "math"
import "fmt"
import "sort"

var VERSION string = "0.2.0"

// Float Slice should be sorted for best results.
// Replaces points within delta of eachother by it's average.
// Reiterates until no reduction.
func Cluster(scatter []float64, delta float64) []float64 {
  var cluster             []float64
  var avg, sum, n, a, b   float64
  var previous            float64   = 0.0
  var i, size             int       = 0, len(scatter)

  for {
    cluster = make([]float64, size)
    for _, a = range(scatter) {
      sum, n = 0.0, 0.0
      for _, b = range(scatter) {
        if math.Abs((a - b)/a) < delta {
          n += 1.0
          sum += b
        }
      }
      avg = sum/n
      if avg != previous {
        cluster[i] = avg
        i++
      }
      previous = avg
    }
    if i == size { break }
    scatter, avg, previous, i, size = cluster[:i], 0.0, 0.0, 0, i
  }

  return cluster
}

type by_near struct {
  slice []float64
  near float64
}
func (a by_near) Len() int {
  return len(a.slice)
}
func (a by_near) Swap(i, j int) {
  s := a.slice
  s[i], s[j] = s[j], s[i]
}
func (a by_near) Less(i, j int) bool {
  s, n := a.slice, a.near
  return math.Abs(s[i] - n) < math.Abs(s[j] - n)
}

func SortByNear(s []float64, n float64) {
  sort.Sort(by_near{s, n})
}

var sprintf = fmt.Sprintf
func Join(a []float64, format string, sep string) string {
  n := len(a)
  if n == 0 { return "" }
  s := sprintf(format, a[0])
  for i:=1; i<n; i++ {
    s += sep
    s += sprintf(format, a[i])
  }
  return s
}

func Sum(a []float64) float64 {
  var i int; n := len(a); sum := 0.0
  for i=0; i<n; i++ { sum += a[i] }
  return sum
}

func Copy(a []float64) []float64 {
  var i, n int = 0, len(a)
  b := make([]float64, n)
  for ; i< n;  i++ { b[i] = a[i] }
  return b
}
