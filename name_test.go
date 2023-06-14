package go_name_generate

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	log.Println("boy:")
	log.Println(GetNames(10, SexBoy))
	log.Println("girl:")
	log.Println(GetNames(10, SexGirl))
	log.Println("all:")
	log.Println(GetNames(10, SexAll))
}
