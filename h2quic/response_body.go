package h2quic

import (
	"io"
	"io/ioutil"

	quic "github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/internal/utils"
)

type responseBody struct {
	eofRead utils.AtomicBool

	dataStream quic.Stream
}

var _ io.ReadCloser = &responseBody{}

func (rb *responseBody) Read(b []byte) (int, error) {
	n, err := rb.dataStream.Read(b)
	if err == io.EOF {
		rb.eofRead.Set(true)
	}
	return n, err
}

func (rb *responseBody) Close() error {
	if !rb.eofRead.Get() {
		rb.dataStream.CancelRead(0)
	}
	// Drain the stream, in order to release all flow control credit.
	// TODO(#1693): this currently fails, if we called CancelRead before
	io.Copy(ioutil.Discard, rb.dataStream)
	// don't call dataStream.Close().
	// The stream might still be used to send data to the server.
	return nil
}
