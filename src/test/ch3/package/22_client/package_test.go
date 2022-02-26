package _2_client

import (
	"go-learn/src/test/ch3/package/22_series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(_2_series.GetFibonacciSerie(5))

	t.Log(_2_series.Square(2))
}
