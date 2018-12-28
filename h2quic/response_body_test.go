package h2quic

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response Body", func() {
	var stream *mockStream

	BeforeEach(func() {
		stream = newMockStream(42)
	})

})
