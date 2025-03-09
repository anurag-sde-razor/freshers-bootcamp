package main

import (
	"encoding/json"
	"fmt"
	"log"
)


type Matrix struct {
	rows    int
	cols    int
	elements [][]int
}

func NewMatrix(rows, cols int) *Matrix {
	elements := make([][]int, rows)
	for i := range elements {
		elements[i] = make([]int, cols)
	}
	return &Matrix{
		rows:     rows,
		cols:     cols,
		elements: elements,
	}
}

func (m *Matrix) GetRows() int {
	return m.rows
}


func (m *Matrix) GetCols() int {
	return m.cols
}


func (m *Matrix) SetElement(i, j, value int) {
	if i >= 0 && i < m.rows && j >= 0 && j < m.cols {
		m.elements[i][j] = value
	}
}

func (m *Matrix) Add(other *Matrix) (*Matrix, error) {
	if m.rows != other.rows || m.cols != other.cols {
		return nil, fmt.Errorf("matrix dimensions must match for addition")
	}
	result := NewMatrix(m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result.elements[i][j] = m.elements[i][j] + other.elements[i][j]
		}
	}
	return result, nil
}

func (m *Matrix) PrintJSON() {
	data := struct {
		Rows     int       `json:"rows"`
		Cols     int       `json:"cols"`
		Elements [][]int  `json:"elements"`
	}{
		Rows:     m.rows,
		Cols:     m.cols,
		Elements: m.elements,
	}


	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("Error marshalling matrix to JSON:", err)
	}
	fmt.Println(string(jsonData))
}

func
	m1 := NewMatrix(2, 2)
	m1.SetElement(0, 0, 1)
	m1.SetElement(0, 1, 2)
	m1.SetElement(1, 0, 3)
	m1.SetElement(1, 1, 4)

	m2 := NewMatrix(2, 2)
	m2.SetElement(0, 0, 5)
	m2.SetElement(0, 1, 6)
	m2.SetElement(1, 0, 7)
	m2.SetElement(1, 1, 8)


	result, err := m1.Add(m2)
	if err != nil {
		log.Fatal(err)
	}

	
	result.PrintJSON()
}
