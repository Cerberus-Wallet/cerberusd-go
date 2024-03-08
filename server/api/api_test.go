package api

import (
	"testing"
)

// Test the origin validation
func TestOriginValidator(t *testing.T) {
	testcases := []struct {
		origin string
		allow  bool
	}{
		// `null` should be denied
		{"null", false},
		// HTTPS for cerberus.uraanai.com should be allowed
		{"https://cerberus.uraanai.com", true},
		{"https://foo.cerberus.uraanai.com", true},
		{"https://bar.foo.cerberus.uraanai.com", true},
		// but HTTP for cerberus.uraanai.com should be denied
		{"http://cerberus.uraanai.com", false},
		{"http://foo.cerberus.uraanai.com", false},
		{"http://bar.foo.cerberus.uraanai.com", false},
		// Fakes should be denied
		{"https://fakecerberus.io", false},
		{"https://foo.fakecerberus.io", false},
		{"http://fakecerberus.io", false},
		{"http://foo.fakecerberus.io", false},
		{"https://foo.cerberus.uraanai.comm", false},
		{"http://foo.cerberus.uraanai.comm", false},
		// Cerberus onion should be allowed
		{"http://cerberusiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", true},
		{"https://cerberusiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", true},
		{"http://foo.cerberusiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", true},
		{"https://foo.cerberusiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", true},
		{"http://bar.foo.cerberusiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", true},
		{"https://bar.foo.cerberusiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", true},
		// Fake cerberus onions should be denied
		{"http://cerberusiovpjcahpzkrewelclulmszwbqpzmzgub48gbcjlvluxtruqad.onion", false},
		{"https://cerberusiovpjcahpzkrewelclulmszwbqpzmzgub48gbcjlvluxtruqad.onion", false},
		{"http://cerberusiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqbd.onion", false},
		{"https://cerberusiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqbd.onion", false},
		{"http://cerberusiowpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		{"https://cerberusiowpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		{"http://foo.cerberusiovpjcahpzkrewelclulmszwbqpzmzgub48gbcjlvluxtruqad.onion", false},
		{"https://bar.foo.cerberusiovpjcahpzkrewelclulmszwbqpzmzgub48gbcjlvluxtruqad.onion", false},
		{"http://fakecerberusiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		{"https://fakecerberusiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		{"http://foo.fakecerberusiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		{"https://foo.fakecerberusiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		{"http://bar.foo.fakecerberusiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		{"https://bar.foo.fakecerberusiovpjcahpzkrewelclulmszwbqpzmzgub37gbcjlvluxtruqad.onion", false},
		// Localhost 8xxx and 5xxx should be allowed for local development
		{"https://localhost:8000", true},
		{"http://localhost:8000", true},
		{"http://localhost:8999", true},
		{"https://localhost:5000", true},
		{"http://localhost:5000", true},
		{"http://localhost:5999", true},
		// SatoshiLabs dev servers should be allowed
		{"https://sldev.cz", true},
		{"https://foo.sldev.cz", true},
		{"https://bar.foo.sldev.cz", true},
		// Fake SatoshiLabs dev servers should be denied
		{"https://fakesldev.cz", false},
		{"https://foo.fakesldev.cz", false},
		{"https://foo.sldev.czz", false},
		{"http://foo.cerberus.sldev.cz", false},
		// Other ports should be denied
		{"http://localhost", false},
		{"http://localhost:1234", false},
	}
	validator := corsValidator()
	for _, tc := range testcases {
		allow := validator(tc.origin)
		if allow != tc.allow {
			t.Errorf("Origin %q: expected %v, got %v", tc.origin, tc.allow, allow)
		}
	}
}
