package hello

import "testing"

func TestGreet(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions. https://davidlovezoe.club/wordpress/archives/tag/golang is awesome" {
		t.Errorf("Greet() = %s; Expected Hello GitHub Actions. https://davidlovezoe.club/wordpress/archives/tag/golang is awesome", result)
	}

}