package extractor

import "io"

type Extractor interface {
	Extract(src, dest string) error
	ExtractStream(dest string, input io.Reader) error
}
