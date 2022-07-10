package hello

import "testing"

func TestGreet(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions. davidlovezoe.club is awesome" {
		t.Errorf("Greet() = %s; Expected Hello GitHub Actions. davidlovezoe.club is awesome", result)
	}

}