package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	list [][]bool
}

func generateRandom() bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float32() > 0.5
}

func (b *Board) refreshBoard() {
	for i := range b.list {
		for j := range b.list[i] {
			b.list[i][j] = generateRandom()
		}
	}
}

func (b Board) inspectNeighbor(row, column int) int {
	if b.valid(row, column) {
		if b.list[row][column] {
			return 1
		}
	}
	return 0
}

func (b Board) checkNeighbors(row, column int) int {
	count := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			count += b.inspectNeighbor(row+i, column+j)
		}
	}
	return count
}

func (b Board) valid(row, col int) bool {
	if row < 0 || col < 0 || len(b.list) < row || len(b.list[0]) < col {
		return false
	}
	return true
}

func (b *Board) step() {
	for i := range b.list {
		for j := range b.list[i] {
			count := b.checkNeighbors(i, j)
		}
	}
}

func (b *Board) replaceValue(count, row, col int) {
	switch count {
	case 3:
		b.list[row][col] = true
	}
}

func (b Board) print() {

}

func main() {
	fmt.Println("hello,world!")
}
