package concurrentmap

import (
	"testing"
)

func TestPutGet(t *testing.T) {
	m := New()

	structValue := struct {
		key   int32
		value string
	}{
		42, "answer",
	}

	cases := []struct {
		in   string
		want interface{}
	}{
		{"Integer value", 432},
		{"String value", "string key"},
		{"String value", "string key 2"},
		{"Array value", [3]string{"Val01", "Val02", "Val03"}},
		{"Struct value", structValue},
		{"Struct reference value", &structValue},
	}
	for _, c := range cases {
		m.Put(c.in, c.want)

		got, found := m.Get(c.in)
		if !found {
			t.Errorf("m.Get(%q) not found", c.in)
		}

		if got != c.want {
			t.Errorf("m.Get(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestPutTwice(t *testing.T) {
	m := New()

	m.Put(1, 1)
	m.Put(1, 2)

	if item, _ := m.Get(1); item != 2 {
		t.Errorf("m.Get(1) == %v, want %v", item, 2)
	}

	if m.Size() != 1 {
		t.Errorf("m.Len() == %v, want %v", m.Size(), 1)
	}
}

func TestNotFound(t *testing.T) {
	m := New()
	got, found := m.Get("key")

	if found {
		t.Error("m.Get(key) returned true")
	}

	if got != nil {
		t.Errorf("m.Get(key) == %v, want %v", got, nil)
	}

}

func TestRemove(t *testing.T) {
	m := New()

	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)

	if found := m.Remove(2); !found {
		t.Errorf("m.Remove(2) == false")
	}

	if _, found := m.Get(2); found {
		t.Errorf("m.Get(2) == found")
	}

	if found := m.Remove(20); found {
		t.Errorf("m.Remove(20) == true")
	}
}
