package acceptance

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os/exec"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("credentials command", func() {
	var (
		server *httptest.Server
	)

	const tableOutput = `+--------------------------------+
|      some-credential-key       |
+--------------------------------+
| some-credential-value          |
| newline                        |
| another-line                   |
| another-line                   |
+--------------------------------+
`
	BeforeEach(func() {
		server = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			switch req.URL.Path {
			case "/uaa/oauth/token":
				w.Write([]byte(`{
				"access_token": "some-opsman-token",
				"token_type": "bearer",
				"expires_in": 3600
			}`))
			case "/api/v0/deployed/products":
				w.Write([]byte(`[
					{"installation_name":"p-bosh","guid":"p-bosh-guid","type":"p-bosh","product_version":"1.10.0.0"},
					{"installation_name":"cf","guid":"cf-guid","type":"cf","product_version":"1.10.0-build.177"},
					{"installation_name":"some-product","guid":"some-product-guid","type":"some-product","product_version":"1.0.0"},
					{"installation_name":"p-isolation-segment","guid":"p-isolation-segment-guid","type":"p-isolation-segment","product_version":"1.10.0-build.31"}
				]`))
			case "/api/v0/deployed/products/some-product-guid/credentials/some-credential":
				w.Write([]byte(`{
					"credential": {
						"type": "some-credential-type",
						"value": {
							"some-credential-key": "some-credential-value\nnewline\nanother-line\nanother-line"
						}
					}
				}`))
			default:
				out, err := httputil.DumpRequest(req, true)
				Expect(err).NotTo(HaveOccurred())
				Fail(fmt.Sprintf("unexpected request: %s", out))
			}
		}))
	})

	It("fetches a credential of a deployed product", func() {
		command := exec.Command(pathToMain,
			"--target", server.URL,
			"--username", "some-username",
			"--password", "some-password",
			"--skip-ssl-validation",
			"credentials",
			"--product-name", "some-product",
			"--credential-reference", "some-credential")

		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))

		Expect(string(session.Out.Contents())).To(Equal(tableOutput))
	})

	It("outputs a specific credential value of a deployed product", func() {
		command := exec.Command(pathToMain,
			"--target", server.URL,
			"--username", "some-username",
			"--password", "some-password",
			"--skip-ssl-validation",
			"credentials",
			"--product-name", "some-product",
			"--credential-reference", "some-credential",
			"--credential-field", "some-credential-key")

		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))

		Expect(string(session.Out.Contents())).To(Equal("some-credential-value\nnewline\nanother-line\nanother-line\n"))
	})

})
