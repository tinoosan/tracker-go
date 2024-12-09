package tracker

import "testing"


func TestAddCategory(t *testing.T) {
  category := "bills!"
  AddCategory(category)
  _, ok := c[category]
  if !ok {
    t.Error("Category could not be added")
  }
  for k := range c {
    delete(c, k)
  }


}
