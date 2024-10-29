package urlValues

import (
	"testing"
)

func getValues() Values {
	values := Values{}
	values.Add("xyz", "abc")
	values.Add("abc", "xyz")
	values.Add("foo", "bar")
	values.Add("foo", "baa")
	values.Add("foo", "baz")
	return values
}

func TestValues_Encode(t *testing.T) {
	values := getValues()

	encoded := values.Encode()

	if encoded != "abc=xyz&foo=bar&foo=baa&foo=baz&xyz=abc" {
		t.Fatalf("failed to encode values. got: %s", encoded)
	}

	values.Del("xyz")

	encoded = values.Encode()

	if encoded != "abc=xyz&foo=bar&foo=baa&foo=baz" {
		t.Fatalf("failed to delete item. got: %s", encoded)
	}
}

func TestValues_EncodeWithOrder(t *testing.T) {
	values := getValues()

	encoded := values.EncodeWithOrder()

	if encoded != "xyz=abc&abc=xyz&foo=bar&foo=baa&foo=baz" {
		t.Fatalf("failed to encode values with order. got: %s", encoded)
	}

	values[OrderKey] = []string{"abc", "xyz", "foo"}

	encoded = values.EncodeWithOrder()

	if encoded != "abc=xyz&xyz=abc&foo=bar&foo=baa&foo=baz" {
		t.Fatalf("failed to encode values with new order. got: %s", encoded)
	}

	values.Del("abc")

	encoded = values.EncodeWithOrder()

	if encoded != "xyz=abc&foo=bar&foo=baa&foo=baz" {
		t.Fatalf("failed to delete item. got: %s", encoded)
	}
}
