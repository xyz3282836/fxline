package other

import (
	"log"

	"go.uber.org/zap"
)

type OtherOne struct {
	log *zap.Logger
}
type OtherTwo struct {
	name string
}

func NewOtherTwo() *OtherTwo {
	return &OtherTwo{}
}

func NewOtherOne(log *zap.Logger) *OtherOne {
	return &OtherOne{log: log}
}

func Hello1(two *OtherTwo) {
	two.name = "hello1"
}

func Hello2(two *OtherTwo) {
	two.name = "hello2"
}

func EchoName(two *OtherTwo) {
	log.Println("11111111111111111 1111echo name is ", two.name)
}
