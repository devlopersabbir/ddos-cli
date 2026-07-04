// Package configs is a package to store configurations
package configs

type Metrics struct {
	total   int64
	success int64
	failed  int64

	totalLetency int64
}
