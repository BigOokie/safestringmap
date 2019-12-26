// Copyright © 2019 BigOokie
//
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package safestringmap

import (
	"testing"

	"github.com/go-test/deep"
)

func Test_NewSafeStringMap(t *testing.T) {
	expect := &SafeStringMap{
		internal: make(map[string]string),
	}

	actual := NewSafeStringMap()

	if diff := deep.Equal(expect, actual); diff != nil {
		t.Error(diff)
	}
}

func Test_Count(t *testing.T) {
	ssm := NewSafeStringMap()

	if ssm.Count() != 0 {
		t.Errorf("NewSageStringMap Count = %d, expected 0", ssm.Count())
	}

	ssm.Store("1", "a")
	if ssm.Count() != 1 {
		t.Errorf("NewSageStringMap Count = %d, expected 1", ssm.Count())
	}

	ssm.Store("2", "b")
	if ssm.Count() != 2 {
		t.Errorf("NewSageStringMap Count = %d, expected 2", ssm.Count())
	}

	ssm.Store("3", "c")
	if ssm.Count() != 3 {
		t.Errorf("NewSageStringMap Count = %d, expected 3", ssm.Count())
	}

	ssm.Store("4", "d")
	ssm.Store("5", "e")
	ssm.Store("6", "f")
	ssm.Store("7", "g")
	ssm.Store("8", "h")
	if ssm.Count() != 8 {
		t.Errorf("NewSageStringMap Count = %d, expected 8", ssm.Count())
	}

	ssm.Store("1", "a2")
	if ssm.Count() != 8 {
		t.Errorf("NewSageStringMap Count = %d, expected 8", ssm.Count())
	}

	ssm.Delete("1")
	if ssm.Count() != 7 {
		t.Errorf("NewSageStringMap Count = %d, expected 7", ssm.Count())
	}

	ssm.Delete("im not here")
	if ssm.Count() != 7 {
		t.Errorf("NewSageStringMap Count = %d, expected 7", ssm.Count())
	}
}

func Test_Load(t *testing.T) {
	ssm := NewSafeStringMap()

	ssm.Store("1", "a")
	if ssm.Count() != 1 {
		t.Error("Expected NewSageStringMap Count to be 1")
	}

	value, ok := ssm.Load("1")
	if !ok {
		t.Error("Expected NewSageStringMap Load to find key '1'")
	} else {
		if value != "a" {
			t.Errorf("Expected NewSageStringMap Load to return value 'a' for key '1'. Got: %s", value)
		}
	}

	value, ok = ssm.Load("im not here")
	if ok {
		t.Error("Did not expect NewSageStringMap Load to find key 'im not here'")
	}
}
