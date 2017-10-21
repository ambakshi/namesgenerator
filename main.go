package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/docker/docker/pkg/namesgenerator"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := namesgenerator.GetRandomName(0)
	fmt.Println(strings.Replace(s, "_", "-", 1))
}
