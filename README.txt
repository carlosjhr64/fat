package fat // import "github.com/carlosjhr64/fat"

Float Array Tricks
var Rounder float64 = 10000.0
var VERSION string = "1.0.0"
func Agglomerate(scatter []float64, delta float64) []float64
func Cluster(scatter []float64, delta float64) []float64
func Copy(a []float64) []float64
func Join(a []float64, format string, sep string) string
func SortByNear(s []float64, n float64)
func Sum(a []float64) float64
