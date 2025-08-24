package main

import (
	"encoding/json"
	"fmt"
)

type requiredFuncs interface {
	getRows() int
	getCols() int
	Set(i, j, val int)
	Add(other *Matrix) *Matrix
	printJson()
}

type Matrix struct {
	NumRow  int
	NumCol  int
	Matrixm [][]int
}

func NewMatrix(row, col int) *Matrix {
	data := make([][]int, row)
	for i := 0; i < row; i++ {
		data[i] = make([]int, col)

	}
	return &Matrix{row, col, data}
}
func (m *Matrix) getRows() int {
	return m.NumRow

}

func (m *Matrix) getCols() int {
	return m.NumCol
}

func (m *Matrix) Set(i, j, val int) {
	if i < 0 || i >= m.NumRow && j < 0 || j >= m.NumCol {
		fmt.Println("Index out of bounds")
		return
	}
	m.Matrixm[i][j] = val

}

func (m *Matrix) Add(other *Matrix) *Matrix {
	if m.NumRow != other.NumRow || m.NumCol != other.NumCol {
		fmt.Println("Dimensions does not match")
		return nil

	}
	newMat := NewMatrix(m.NumRow, m.NumCol)
	for i := 0; i < m.NumRow; i++ {
		for j := 0; j < m.NumCol; j++ {
			newMat.Matrixm[i][j] = m.Matrixm[i][j] + other.Matrixm[i][j]

		}
	}
	return newMat
}

func (m *Matrix) printJson() {
	jsonData, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("Error marshalling to json :", err)
		return
	}

	fmt.Println(string(jsonData))
}
func main() {
	mat1 := NewMatrix(2, 2)
	mat2 := NewMatrix(2, 2)

	mat1.Set(0, 0, 1)
	mat1.Set(0, 1, 2)
	mat1.Set(1, 0, 3)
	mat1.Set(1, 1, 4)

	mat2.Set(0, 0, 5)
	mat2.Set(0, 1, 6)
	mat2.Set(1, 0, 7)
	mat2.Set(1, 1, 8)

	newMat := mat1.Add(mat2)

	fmt.Println(newMat.Matrixm) //we do not need to derefernce as we are getting the address of the newMatrix fromt he add methid but the go has a special feature of automatic derefernce therer fore we can do that without derefernce

	if newMat != nil {
		fmt.Println(newMat.Matrixm)
		newMat.printJson()
	}

}
