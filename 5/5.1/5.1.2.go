package main

import "time"

type status int

type BetterTask struct {
	title  string
	status status
	due    *time.Time
}

const UNKNOWN status = 0
const TODO status = 1
const DONE status = 2

// it is better
// const UNKNOWN  	status = 0
// const TODO 		status = 1
// const DONE 		status = 2

//it is much better ( 굳이 힘들게 0,1,2 쓰지말자 )
// const (
// 	UNKNOWN status = iota
// 	TODO
// 	DONE
// )

//활용하기
type ByteSize int

const (
	KB ByteSize = 1 << (10 * (1 + iota))
	MB
	GB
	TB
	PB
)
