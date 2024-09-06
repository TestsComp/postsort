package postsort_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"postsort/postsort"
)

func TestCommonSyntax(t *testing.T) {
	assert.Nil(t, nil, "Common syntax error")
}

type testPackage struct {
	width, height, length uint
	mass                  float32
	result                postsort.SortResult
	err                   error
}

func Test_Sort(t *testing.T) {
	packages := []testPackage{
		// wrong data
		{width: 0, height: 50, length: 50, mass: 10, result: postsort.REJECTED, err: postsort.InvalidPackageError},
		{width: 50, height: 0, length: 50, mass: 10, result: postsort.REJECTED, err: postsort.InvalidPackageError},
		{width: 50, height: 50, length: 0, mass: 10, result: postsort.REJECTED, err: postsort.InvalidPackageError},
		{width: 10, height: 10, length: 10, mass: -1, result: postsort.REJECTED, err: postsort.InvalidPackageError},

		// REJECTED, dims: 100 * 100 * 100 = 1_000_000, 20 kg
		{width: 100, height: 100, length: 100, mass: 20, result: postsort.REJECTED, err: nil},
		{width: 100, height: 100, length: 100, mass: 20.001, result: postsort.REJECTED, err: nil},
		{width: 100, height: 100, length: 101, mass: 20, result: postsort.REJECTED, err: nil},

		// SPECIAL by dimensions
		{width: 150, height: 10, length: 10, mass: 10, result: postsort.SPECIAL, err: nil},
		{width: 10, height: 150, length: 10, mass: 10, result: postsort.SPECIAL, err: nil},
		{width: 10, height: 10, length: 150, mass: 10, result: postsort.SPECIAL, err: nil},
		{width: 100, height: 101, length: 100, mass: 10, result: postsort.SPECIAL, err: nil},

		// SPECIAL by mass
		{width: 10, height: 10, length: 10, mass: 21.01, result: postsort.SPECIAL, err: nil},

		// STANDARD
		{width: 10, height: 10, length: 10, mass: 10, result: postsort.STANDARD, err: nil},
		{width: 1, height: 1, length: 1, mass: 1, result: postsort.STANDARD, err: nil},
	}

	for _, p := range packages {
		res, err := postsort.Sort(p.width, p.height, p.length, p.mass)
		str := fmt.Sprintf("width: %d, height: %d, length: %d, mass: %f", p.width, p.height, p.length, p.mass)

		// t TestingT, expected, actual interface{}, msgAndArgs ...interface{})
		assert.Equal(t, p.result, res, "result for %s"+str)

		if p.err != nil {
			assert.NotNil(t, err, "error for %s", str)
		} else {
			assert.Nil(t, err, "no error for %s", str)
		}
	}

}
