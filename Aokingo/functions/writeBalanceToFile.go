package functions

import (
	"os"
	"strconv"
)

func WriteBalanceToFile(filename string, balance int) {
	os.WriteFile(filename, []byte(strconv.Itoa(balance)), 0o644)
}
