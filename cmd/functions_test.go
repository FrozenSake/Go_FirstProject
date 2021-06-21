package main

import (
	"errors"
	"math"
	"strconv"
	"testing"
)

var strMaxInt string = strconv.Itoa(math.MaxInt32) //"2147483647"

// Returns correct test matrix
func getTestMatrix_Ok() [][]string {
	return [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
}

// Returns nil test matrix
func getTestMatrix_Nil() [][]string {
	var m [][]string = nil
	return m
}

// Returns empty test matrix
func getTestMatrix_Empty() [][]string {
	var m [][]string = make([][]string, 0)
	return m
}

// Returns non square test matrix
func getTestMatrix_NonSquare() [][]string {
	return [][]string{{"1", "2", "3"}, {"4", "5"}, {"7", "8", "9"}}
}

// Returns non integer test matrix
func getTestMatrix_NonInteger() [][]string {
	return [][]string{{"1", "2", "3"}, {"a", "b", "c"}, {"7", "8", "9"}}
}

// Returns big values test matrix
func getTestMatrix_Big() [][]string {
	return [][]string{{strMaxInt, strMaxInt, strMaxInt}, {strMaxInt, strMaxInt, strMaxInt}, {strMaxInt, strMaxInt, strMaxInt}}
}

// Test for About func
func TestAbout(t *testing.T) {
	got := About()
	want := "Backend Challenge."

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Test for Echo func
func TestEcho(t *testing.T) {
	t.Run("correct matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Ok()
		got, _ := Echo(testMatrix)
		want := "1,2,3\n4,5,6\n7,8,9"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("nil matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Nil()
		_, got := Echo(testMatrix)
		want := errors.New("matrix is nil")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("empty matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Empty()
		_, got := Echo(testMatrix)
		want := errors.New("matrix has no elements")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("non square matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_NonSquare()
		_, got := Echo(testMatrix)
		want := errors.New("matrix must be a square matrix")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("non integer matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_NonInteger()
		_, got := Echo(testMatrix)
		want := errors.New("matrix must have only integer elements")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("big matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Big()
		got, _ := Echo(testMatrix)
		want := strMaxInt + "," + strMaxInt + "," + strMaxInt + "\n" + strMaxInt + "," + strMaxInt + "," + strMaxInt + "\n" + strMaxInt + "," + strMaxInt + "," + strMaxInt

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})
}

// Test for Invert func
func TestInvert(t *testing.T) {
	t.Run("correct matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Ok()
		got, _ := Invert(testMatrix)
		want := "1,4,7\n2,5,8\n3,6,9"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("nil matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Nil()
		_, got := Invert(testMatrix)
		want := errors.New("matrix is nil")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("empty matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Empty()
		_, got := Invert(testMatrix)
		want := errors.New("matrix has no elements")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("non square matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_NonSquare()
		_, got := Invert(testMatrix)
		want := errors.New("matrix must be a square matrix")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("non integer matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_NonInteger()
		_, got := Invert(testMatrix)
		want := errors.New("matrix must have only integer elements")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("big matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Big()
		got, _ := Invert(testMatrix)
		want := strMaxInt + "," + strMaxInt + "," + strMaxInt + "\n" + strMaxInt + "," + strMaxInt + "," + strMaxInt + "\n" + strMaxInt + "," + strMaxInt + "," + strMaxInt

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})
}

// Test for Flatten func
func TestFlatten(t *testing.T) {
	t.Run("correct matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Ok()
		got, _ := Flatten(testMatrix)
		want := "1,2,3,4,5,6,7,8,9"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("nil matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Nil()
		_, got := Flatten(testMatrix)
		want := errors.New("matrix is nil")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("empty matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Empty()
		_, got := Flatten(testMatrix)
		want := errors.New("matrix has no elements")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("non square matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_NonSquare()
		_, got := Flatten(testMatrix)
		want := errors.New("matrix must be a square matrix")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("non integer matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_NonInteger()
		_, got := Flatten(testMatrix)
		want := errors.New("matrix must have only integer elements")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("big matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Big()
		got, _ := Flatten(testMatrix)
		want := strMaxInt + "," + strMaxInt + "," + strMaxInt + "," + strMaxInt + "," + strMaxInt + "," + strMaxInt + "," + strMaxInt + "," + strMaxInt + "," + strMaxInt

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})
}

// Test for Sum func
func TestSum(t *testing.T) {
	t.Run("correct matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Ok()
		got, _ := Sum(testMatrix)
		want := "45"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("nil matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Nil()
		_, got := Sum(testMatrix)
		want := errors.New("matrix is nil")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("empty matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Empty()
		_, got := Sum(testMatrix)
		want := errors.New("matrix has no elements")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("non square matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_NonSquare()
		_, got := Sum(testMatrix)
		want := errors.New("matrix must be a square matrix")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("non integer matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_NonInteger()
		_, got := Sum(testMatrix)
		want := errors.New("matrix must have only integer elements")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("big matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Big()
		got, _ := Sum(testMatrix)
		want := "19327352824"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})
}

// Test for Multiply func
func TestMultiply(t *testing.T) {
	t.Run("correct matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Ok()
		got, _ := Multiply(testMatrix)
		want := "362880"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("nil matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Nil()
		_, got := Multiply(testMatrix)
		want := errors.New("matrix is nil")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("empty matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Empty()
		_, got := Multiply(testMatrix)
		want := errors.New("matrix has no elements")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("non square matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_NonSquare()
		_, got := Multiply(testMatrix)
		want := errors.New("matrix must be a square matrix")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("non integer matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_NonInteger()
		_, got := Multiply(testMatrix)
		want := errors.New("matrix must have only integer elements")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})

	t.Run("big matrix", func(t *testing.T) {
		testMatrix := getTestMatrix_Big()
		got, _ := Multiply(testMatrix)
		want := "971334442042048905792818449949263431919261344060050972981379158342510148733965959167"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, testMatrix)
		}
	})
}
