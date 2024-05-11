package ds

import "testing"

func TestNewLinkedList(t *testing.T) {
	l := NewLinkedList()
	if l.First != nil || l.Last != nil || l.Size != 0 {
		t.Errorf("Linked List should be created empty, first=%v last=%v size=%d", l.First, l.Last, l.Size)
	}
}

func TestInsertNewNode(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		expected_last  any
		expected_first any
		expected_size  int
	}{
		{"InsertOneElement", []any{2}, 2, 2, 1},
		{"InsertMultipleElements", []any{4, 3.8}, 3.8, 4, 2},
		{"InsertMultipleElements", []any{1.1, 2, 10, 100, 5.4}, 5.4, 1.1, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := NewLinkedList()
			for _, element := range test.elements {
				list.Insert(element)
			}

			firstNode := list.First.value
			lastNode := list.Last.value

			if firstNode != test.expected_first || lastNode != test.expected_last || list.Size != test.expected_size {
				t.Errorf("Error while inserting, expected first=%v expected_last=%v expected_size=%d", test.expected_first, test.expected_last, test.expected_size)
			}
		})
	}
}

func TestSearchNode(t *testing.T) {
	tests := []struct {
		name         string
		elements     []any
		target       any
		expected_pos int
	}{
		{"SearchFirstNode", []any{2, 4, 6, 8}, 2, 0},
		{"SearchLastNode", []any{1, 3, 5, 7}, 7, 3},
		{"SearchRandomNode", []any{2, 5, 7, 17, 1, 4}, 1, 4},
		{"SearchTargetDontExist", []any{2.4, -2, 5.2, 10, 17, 89}, 100, -1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := NewLinkedList()
			for _, element := range test.elements {
				list.Insert(element)
			}
			got := list.Search(test.target)

			if got != test.expected_pos {
				t.Errorf("Error while searching, expected_pos=%d, got=%d", test.expected_pos, got)
			}
		})
	}
}
