package storage

import "testing"

func TestStorage(t *testing.T) {

	s := Local()

	s.SetItem("bgcolor", "red")
	s.SetItem("font", "Helvetica")
	s.SetItem("image", "miGato.png")

	if s.Length() != 3 {
		t.Errorf("Length failure")
	}

	v, ok := s.GetItem("bgcolor")
	if !ok || v != "red" {
		t.Errorf("GetItem failure")
	}

	s.RemoveItem("bgcolor")
	v, ok = s.GetItem("bgcolor")
	if ok || v != "" {
		t.Errorf("RemoveItem failure")
	}

	s.Clear()

	if s.Length() != 0 {
		t.Errorf("Clear Failure")
	}
}
