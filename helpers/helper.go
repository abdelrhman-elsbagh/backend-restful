package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%04d", rand.Intn(10000))
}
