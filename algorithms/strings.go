package algorithms

import (
	"bytes"
	"fmt"
	"strings"
)

// Join ...
func Join(data []string) string {

	return strings.Join(data, ":")
}

// Sprintf ...
func Sprintf(data []string) string {

	return fmt.Sprintf("%s:%s:%s:%s:%s", data[0], data[1], data[2], data[3], data[4])
}

// Concat ...
func Concat(data []string) string {
	s := data[0] + ":"
	s += data[1] + ":"
	s += data[2] + ":"
	s += data[3] + ":"
	s += data[4]
	return s
}

// ConcatOneLine ...
func ConcatOneLine(data []string) string {
	return data[0] + ":" +
		data[1] + ":" +
		data[2] + ":" +
		data[3] + ":" +
		data[4]
}

// Buffer ...
func Buffer(data []string) string {

	var b bytes.Buffer
	b.WriteString(data[0])
	b.WriteByte(':')
	b.WriteString(data[1])
	b.WriteByte(':')
	b.WriteString(data[2])
	b.WriteByte(':')
	b.WriteString(data[3])
	b.WriteByte(':')
	b.WriteString(data[4])
	return b.String()

}

// BufferWithReset ...
func BufferWithReset(data []string) string {
	var buf bytes.Buffer
	buf.Reset()

	buf.WriteString(data[0])
	buf.WriteByte(':')
	buf.WriteString(data[1])
	buf.WriteByte(':')
	buf.WriteString(data[2])
	buf.WriteByte(':')
	buf.WriteString(data[3])
	buf.WriteByte(':')
	buf.WriteString(data[4])
	return buf.String()
}

// BufferFprintf ...
func BufferFprintf(data []string) string {
	buf := &bytes.Buffer{}
	buf.Reset()

	fmt.Fprintf(buf, "%s:%s:%s:%s:%s", data[0], data[1], data[2], data[3], data[4])
	return buf.String()

}

// BufferStringBuilder ...
func BufferStringBuilder(data []string) string {
	var buf strings.Builder
	buf.Reset()

	buf.WriteString(data[0])
	buf.WriteByte(':')
	buf.WriteString(data[1])
	buf.WriteByte(':')
	buf.WriteString(data[2])
	buf.WriteByte(':')
	buf.WriteString(data[3])
	buf.WriteByte(':')
	buf.WriteString(data[4])
	return buf.String()
}
