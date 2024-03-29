package localip_test

import (
	"net"

	"code.cloudfoundry.org/localip"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Localip", func() {
	Describe("LocalIP", func() {
		It("returns a local IP", func() {
			ip, err := localip.LocalIP()
			Expect(err).NotTo(HaveOccurred())

			// http://golang.org/pkg/net/#ParseIP
			// If s is not a valid textual representation of an IP address, ParseIP returns nil.
			Expect(net.ParseIP(ip)).NotTo(BeNil())
		})
	})

	Describe("LocalPort", func() {
		It("returns a local port", func() {
			port, err := localip.LocalPort()
			Expect(err).NotTo(HaveOccurred())
			Expect(port).To(BeNumerically(">", 0))
		})
	})
})
