package fat

import "testing"

func gimea() []float64 {
  a := make([]float64, 5)
  a[0] = 10.0
  a[1] = 12.0
  a[2] = 20.0
  a[3] = 29.0
  a[4] = 31.0
  return a
}

func TestCluster(test *testing.T) {
  bad := test.Error
  a := gimea()
  // Want (12,10) and (29,31) to be considered a cluster
  delta := (12.1 - 9.9) / 10.0
  b := Cluster(a, delta)
  if len(b) != 3 { bad("Cluster len.") }
  if b[0] != 11.0 || b[1] != 20.0 || b[2] != 30.0 { bad("Cluster values.") }
}

func TestSortByNear(test *testing.T) {
  bad := test.Error
  a := gimea()
  SortByNear(a, 19.0)
  if a[0]!=20.0 || a[1]!=12.0 || a[2]!=10.0 || a[3]!=29.0 || a[4]!=31.0 {
    bad("SortByNear")
  }
}

func TestJoin(test *testing.T) {
  bad := test.Error
  a := gimea()
  s := Join(a, "$%.2f", ", ")
  if s != "$10.00, $12.00, $20.00, $29.00, $31.00" {
    bad(Join)
  }
}

func TestSum(test *testing.T) {
  bad := test.Error
  a := gimea()
  s := Sum(a)
  if s != 102.0 { bad("Sum") }
}
