package functions

import (
	"os"
	"strconv"
)

func ReadBalanceFromFile(filename string) (int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	balance, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}
	return balance, nil
}
