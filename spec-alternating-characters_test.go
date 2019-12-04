package main

import "testing"

func TestAlternatingCharacters1(t *testing.T) {
	s := "AAAA"

	deletions := alternatingCharacters(s)

	if deletions != 3 {
		t.Errorf("got %d instead of 3", deletions)
	}
}

func TestAlternatingCharacters2(t *testing.T) {
	s := "BBBBB"

	deletions := alternatingCharacters(s)

	if deletions != 4 {
		t.Errorf("got %d instead of 4", deletions)
	}
}

func TestAlternatingCharacters3(t *testing.T) {
	s := "ABABABAB"

	deletions := alternatingCharacters(s)

	if deletions != 0 {
		t.Errorf("got %d instead of 0", deletions)
	}
}

func TestAlternatingCharacters4(t *testing.T) {
	s := "BABABA"

	deletions := alternatingCharacters(s)

	if deletions != 0 {
		t.Errorf("got %d instead of 0", deletions)
	}
}

func TestAlternatingCharacters5(t *testing.T) {
	s := "AAABBB"

	deletions := alternatingCharacters(s)

	if deletions != 4 {
		t.Errorf("got %d instead of 4", deletions)
	}
}
