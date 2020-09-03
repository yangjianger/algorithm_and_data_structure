package cum_list

import "testing"

const (
	_TEST_CAPACITY = 10
)

func TestArrayList(t *testing.T) {
	arratList := NewArrayList(_TEST_CAPACITY)
	for i := 0; i < _TEST_CAPACITY; i++ {
		arratList.Add(i)
	}

	if arratList.Size() != _TEST_CAPACITY{
		t.Error("arratList Size error")
	}

	if arratList.Get(9) != 9{
		t.Error("arratList Get error")
	}
}

