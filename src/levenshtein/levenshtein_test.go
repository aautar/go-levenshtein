package levenshtein

import "testing"

func TestDistanceInsertion(t *testing.T) {
	t.Log("Testing calculation with insertion of 1 character")

	distInsertEnd := Distance("abc", "abc0")
	distInsertBeginning := Distance("abc", "0abc")
	distInsertMiddle := Distance("abc", "a0bc")

	if distInsertEnd != 1 || distInsertBeginning != 1 || distInsertMiddle != 1 {
		t.Errorf("Expected distance of 1 not found")
		t.Errorf("distInsertEnd = %d", distInsertEnd)
		t.Errorf("distInsertBeginning = %d", distInsertBeginning)
		t.Errorf("distInsertMiddle = %d", distInsertMiddle)
	}
}

func TestDistanceRemoval(t *testing.T) {
	t.Log("Testing calculation with removal of 1 character")

	distRemoveEnd := Distance("abc", "ab")
	distRemoveBeginning := Distance("abc", "bc")
	distRemoveMiddle := Distance("abc", "ac")

	if distRemoveEnd != 1 || distRemoveBeginning != 1 || distRemoveMiddle != 1 {
		t.Errorf("Expected distance of 1 not found")
		t.Errorf("distRemoveEnd = %d", distRemoveEnd)
		t.Errorf("distRemoveBeginning = %d", distRemoveBeginning)
		t.Errorf("distRemoveMiddle = %d", distRemoveMiddle)
	}
}

func TestDistanceChange(t *testing.T) {
	t.Log("Testing calculation with change of 1 character")

	distChangeEnd := Distance("abc", "ab0")
	distChangeBeginning := Distance("abc", "0bc")
	distChangeMiddle := Distance("abc", "a0c")

	if distChangeEnd != 1 || distChangeBeginning != 1 || distChangeMiddle != 1 {
		t.Errorf("Expected distance of 1 not found")
		t.Errorf("distChangeEnd = %d", distChangeEnd)
		t.Errorf("distChangeBeginning = %d", distChangeBeginning)
		t.Errorf("distChangeMiddle = %d", distChangeMiddle)
	}
}
