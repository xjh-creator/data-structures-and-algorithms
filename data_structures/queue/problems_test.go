package queue

import (
	"log"
	"testing"
)

func TestMyStack(t *testing.T) {
	q := NewMyStack(5)
	q.Push(1)
	log.Println(q.Top())
	log.Println(q.Empty())
	log.Println(q.Pop())
	log.Println(q.Empty())
}

