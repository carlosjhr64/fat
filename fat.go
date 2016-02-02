// Float Array Tricks
package fat

import "math"
import "fmt"
import "sort"

var VERSION string = "1.0.0"
var Rounder float64 = 10000.0

// Float Slice should be sorted for best results.
// Replaces points within delta of eachother by it's average.
// Reiterates until no reduction.
func Cluster(scatter []float64, delta float64) []float64 {
  size := len(scatter)
  cluster := make([]float64, size)
  var index, i, j int
  var previous, avg, n, a, b float64
  for {
    index, previous = 0, 0.0
    for i=0; i<size; i++ {
      a = scatter[i]
      avg, n = 0.0, 0.0
      for j=0; j<size; j++ {
        if i==j {continue}
        b = scatter[j]
        if math.Abs(a - b) < delta {
          avg += b
          n += 1.0
        }
      }
      if n == 0.0 {
        avg = a
      } else {
        avg += a
        n += 1.0
        avg = float64(int(0.5 + Rounder*(avg/n)))/Rounder
      }
      if avg != previous {
        cluster[index] = avg
        index += 1
      }
      previous = avg
    }
    if index == size { break }
    for i=0; i<index; i++ { scatter[i] = cluster[i] }
    size = index
  }
  return scatter[:size]
}

func Agglomerate(scatter []float64, delta float64) []float64 {
  // Very sensitive to rounding errors...
  // Ruby's implementation of the exact same algorithm
  // yielded different lists until rounding was introduced.
  size := len(scatter)
  cluster := make([]float64, size)
  var index, i, j int
  var previous, avg, n, a, b float64
  for {
    index, previous = 0, 0.0
    for i=0; i<size; i++ {
      a = scatter[i]
      avg, n = 0.0, 0.0
      for j=0; j<size; j++ {
        if i==j {continue}
        b = scatter[j]
        if math.Abs(math.Log(a/b)) < delta {
          avg += math.Log(b)
          n += 1.0
        }
      }
      if n == 0.0 {
        avg = a
      } else {
        avg += math.Log(a)
        n += 1.0
        avg = float64(int(0.5 + Rounder*math.Exp(avg/n)))/Rounder
      }
      if avg != previous {
        cluster[index] = avg
        index += 1
      }
      previous = avg
    }
    if index == size { break }
    for i=0; i<index; i++ { scatter[i] = cluster[i] }
    size = index
  }
  return scatter[:size]
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

func Join(a []float64, format string, sep string) string {
  n := len(a)
  if n == 0 { return "" }
  s := fmt.Sprintf(format, a[0])
  for i:=1; i<n; i++ {
    s += sep
    s += fmt.Sprintf(format, a[i])
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
