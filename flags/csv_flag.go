package flags

import (
	"fmt"
	"strconv"
	"strings"
)

// CSIntFlagValue holds a slice of integers parsed from a CSV string.
type CSIntFlagValue []int

// String returns the CSV representation of the integers.
func (csi *CSIntFlagValue) String() string {
	sb := strings.Builder{}
	for n, v := range *csi {
		if n != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprint(v))
	}
	return sb.String()
}

// Set parses a CSV string into integers and stores them.
func (csi *CSIntFlagValue) Set(value string) error {
	parts := strings.Split(value, ",")
	for _, value := range parts {
		v, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("illegal value: %s [err=%w]", value, err)
		}
		*csi = append(*csi, v)
	}
	return nil
}
