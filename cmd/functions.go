package main

import (
	"encoding/csv"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	//"log"
	"net/http"
)

// Returns a simple test string.
//	Output:
// "Backend Challenge."
func About() string {
	return "Backend Challenge."
}

// Returns the matrix as a string in matrix format.
//	Output:
// 	- Given
// 		[1,2,3]
// 		[4,5,6]
// 		[7,8,9]
// 	- Expected
// 		"1,2,3
// 		 4,5,6
// 		 7,8,9"
func Echo(matrix [][]string) (string, error) {
	if canUseGivenMatrix, err := isValidMatrix(matrix); !canUseGivenMatrix {
		return "", err
	}

	var result string
	for _, row := range matrix {
		result = fmt.Sprintf("%s%s\n", result, strings.Join(row, ","))
	}

	return result, nil
}

// Returns the matrix as a string in matrix format where the columns and rows are inverted.
//	Output:
// 	- Given
// 		[1,2,3]
// 		[4,5,6]
// 		[7,8,9]
// 	- Expected
// 		"1,4,7
// 		 2,5,8
// 		 3,6,9"
func Invert(matrix [][]string) (string, error) {
	if canUseGivenMatrix, err := isValidMatrix(matrix); !canUseGivenMatrix {
		return "", err
	}

	rows := len(matrix)
	cols := len(matrix[0])
	transpose := make([][]string, cols)

	for rowIndex := range transpose {
		transpose[rowIndex] = make([]string, rows)
	}
	for row := 0; row < cols; row++ {
		for col := 0; col < rows; col++ {
			transpose[row][col] = matrix[col][row]
		}
	}

	var result string
	for _, row := range transpose {
		result = fmt.Sprintf("%s%s\n", result, strings.Join(row, ","))
	}

	return result, nil
}

// Returns the matrix as a 1 line string, with values separated by commas.
//	Output:
// 	- Given
// 		[1,2,3]
// 		[4,5,6]
// 		[7,8,9]
// 	- Expected
// 		"1,2,3,4,5,6,7,8,9"
func Flatten(matrix [][]string) (string, error) {
	if canUseGivenMatrix, err := isValidMatrix(matrix); !canUseGivenMatrix {
		return "", err
	}

	var result string
	for _, row := range matrix {
		for _, col := range row {
			if len(result) > 0 {
				result = result + ","
			}
			result = result + col
		}
	}
	return result, nil
}

// Returns the sum of the integers in the matrix.
//	Output:
// 	- Given
// 		[1,2,3]
// 		[4,5,6]
// 		[7,8,9]
// 	- Expected
// 		"45"
func Sum(matrix [][]string) (string, error) {
	if canUseGivenMatrix, err := isValidMatrix(matrix); !canUseGivenMatrix {
		return "", err
	}

	result := big.NewInt(0)
	for _, row := range matrix {
		for _, col := range row {
			if numericCol, err := new(big.Int).SetString(col, 0); err {
				result = result.Add(result, numericCol)
			} else {
				return "", fmt.Errorf("value \"%v\" is not an integer", col)
			}
		}
	}
	return result.Text(10), nil
}

// Returns the product of the integers in the matrix.
// Output:
// 	- Given
// 		[1,2,3]
// 		[4,5,6]
// 		[7,8,9]
// 	- Expected
// 		"362880"
func Multiply(matrix [][]string) (string, error) {
	if canUseGivenMatrix, err := isValidMatrix(matrix); !canUseGivenMatrix {
		return "", err
	}

	result := big.NewInt(1)
	for _, row := range matrix {
		for _, col := range row {
			if numericCol, err := new(big.Int).SetString(col, 0); err {
				result = result.Mul(result, numericCol)
			} else {
				return "", fmt.Errorf("value \"%v\" is not an integer", col)
			}
		}
	}
	return result.Text(10), nil
}

// Returns whether given matrix is valid
func isValidMatrix(matrix [][]string) (bool, error) {

	// Check whether given matrix...

	// is nil
	if matrix == nil {
		return false, fmt.Errorf("matrix is nil")
	}
	// has elements
	if len(matrix) == 0 {
		return false, fmt.Errorf("matrix has no elements")
	}
	// is a square matrix or not
	if canUseGivenMatrix, err := isSquareMatrix(matrix); !canUseGivenMatrix {
		return false, err
	}
	// has only integer elements
	if canUseGivenMatrix, err := isIntegerMatrix(matrix); !canUseGivenMatrix {
		return false, err
	}

	return true, nil
}

// Returns whether given matrix is a square matrix or not
func isSquareMatrix(matrix [][]string) (bool, error) {
	rows := len(matrix)
	for _, row := range matrix {
		cols := len(row)
		if rows != cols {
			return false, fmt.Errorf("matrix must be a square matrix")
		}
	}
	return true, nil
}

// Returns whether given matrix has only integer elements
func isIntegerMatrix(matrix [][]string) (bool, error) {
	for _, row := range matrix {
		for _, col := range row {
			if _, err := strconv.Atoi(col); err != nil {
				return false, fmt.Errorf("matrix must have only integer elements")
			}
		}
	}

	return true, nil
}

// Reads the uploaded csv file and extracts it's matrix.
func GetFileMatrix(r *http.Request) ([][]string, error) {
	var matrix [][]string = nil

	file, _, err := r.FormFile("file")
	if err != nil {
		return matrix, fmt.Errorf("could not find file \"%v\". error %v", file, err.Error())
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return matrix, fmt.Errorf("could not read file \"%v\". error %v", file, err.Error())
	}

	return records, nil
}

func ErrorHandler(w http.ResponseWriter, err error, message string) bool {
	if err != nil {
		w.Write([]byte(fmt.Sprintf(message+" error %s", err.Error())))
		// TODO add log
		return false
	}
	return true
}
