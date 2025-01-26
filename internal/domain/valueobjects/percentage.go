package valueobjects

import "fmt"


type Percentage struct {
	Value float64
}


func NewPercentage(value float64) (*Percentage, error) {
 if value < 0 || value > 1 {
    return &Percentage{}, fmt.Errorf("percent must be between 0 and 1")
  }

  return &Percentage{Value: value}, nil

}
func (p *Percentage) Apply(amount int) int {
  return int(float64(amount)*p.Value)
}
