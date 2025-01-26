package valueobjects

import "fmt"

type Ratio struct {
  Value float64
}

func NewRatio(value float64) (*Ratio, error) {
  if value <=0 {
    return &Ratio{}, fmt.Errorf("ratio must be positive")
  }
  return &Ratio{Value: value}, nil
}

func (r *Ratio) Apply(amount int) int {
  return int(float64(amount)*r.Value)
}
