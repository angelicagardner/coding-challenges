package leetcode

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	var tests = []struct {
		name     string
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{"Both lists empty", nil, nil, nil},
		{"One list empty, one non-empty",
			&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			nil,
			&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		},
		{"Both lists have elements",
			&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
		{"l1 has smaller elements",
			&ListNode{1, &ListNode{3, &ListNode{5, nil}}},
			&ListNode{2, &ListNode{4, &ListNode{6, nil}}},
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}},
		},
		{"l2 has smaller elements",
			&ListNode{2, &ListNode{4, &ListNode{6, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{5, nil}}},
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}},
		},
		{"Lists with duplicate elements",
			&ListNode{1, &ListNode{1, &ListNode{1, nil}}},
			&ListNode{1, &ListNode{1, &ListNode{1, nil}}},
			&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, nil}}}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merged := MergeTwoLists(tt.l1, tt.l2)

			currentMerged := merged
			currentExpected := tt.expected

			for currentMerged != nil && currentExpected != nil {
				if currentMerged.Val != currentExpected.Val {
					t.Errorf("MergeTwoLists() = %v, expected %v", currentMerged.Val, currentExpected.Val)
					return
				}
				currentMerged = currentMerged.Next
				currentExpected = currentExpected.Next
			}

			if currentMerged != nil || currentExpected != nil {
				t.Errorf("MergeTwoLists() lists have different lengths")
			}
		})
	}
}
