package main

import (
	"testing"
)

//Normal test to the newFooter function
func newFooterTest(t *testing.T) {

	result := newFooter(1, 15, 1, 1).List
	expected := []string{"1", "2", "...", "15"}
	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Error()
		}
	}

}

//Test the newFooter function if the currentpage is negative
func currentPageNegativeTest(t *testing.T) {

	result := newFooter(-1, 15, 1, 1).List
	expected := []string{}
	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Error()
		}
	}

}

//Test the newFooter function if the around is negative
func aroundNegativeTest(t *testing.T) {

	result := newFooter(1, 15, 0, 1).List
	expected := []string{}
	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Error()
		}
	}

}

//Test the newFooter function if the totalpages is equal to 0
func totalPagesNegativeTest(t *testing.T) {

	result := newFooter(2, -1, 1, 1).List
	expected := []string{}
	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Error()
		}
	}

}

//Test the newFooter function if the totalpages is equal to 0
func boundariesNegativeTest(t *testing.T) {

	result := newFooter(1, 15, 1, -1).List
	expected := []string{}
	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Error()
		}
	}

}
