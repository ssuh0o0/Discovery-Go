// 포인터 리시버 : 자료형이 포인터인 리시버
package main

import "io"

type Graph [][]int

func WriteTo(w io.Writer, adjList [][]int) error

// ReadFrom은 포인터 자료형이므로 포인트 리시버가 필요하다.
func ReadFrom(r io.Reader, adjList *[][]int) error

// func (adjList Graph) WriteTo(w io.Writer) error
// func (adjList *Graph) ReadFrom(r io.Reader) error
