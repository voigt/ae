/* -------------------------------------------------------------------------- *\
 *             Apache 2.0 License Copyright © 2022 The Aurae Authors          *
 *                                                                            *
 *                +--------------------------------------------+              *
 *                |   █████╗ ██╗   ██╗██████╗  █████╗ ███████╗ |              *
 *                |  ██╔══██╗██║   ██║██╔══██╗██╔══██╗██╔════╝ |              *
 *                |  ███████║██║   ██║██████╔╝███████║█████╗   |              *
 *                |  ██╔══██║██║   ██║██╔══██╗██╔══██║██╔══╝   |              *
 *                |  ██║  ██║╚██████╔╝██║  ██║██║  ██║███████╗ |              *
 *                |  ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝ |              *
 *                +--------------------------------------------+              *
 *                                                                            *
 *                         Distributed Systems Runtime                        *
 *                                                                            *
 * -------------------------------------------------------------------------- *
 *                                                                            *
 *   Licensed under the Apache License, Version 2.0 (the "License");          *
 *   you may not use this file except in compliance with the License.         *
 *   You may obtain a copy of the License at                                  *
 *                                                                            *
 *       http://www.apache.org/licenses/LICENSE-2.0                           *
 *                                                                            *
 *   Unless required by applicable law or agreed to in writing, software      *
 *   distributed under the License is distributed on an "AS IS" BASIS,        *
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. *
 *   See the License for the specific language governing permissions and      *
 *   limitations under the License.                                           *
 *                                                                            *
\* -------------------------------------------------------------------------- */

package pki

import (
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestCreateAuraeRootCA(t *testing.T) {
	t.Run("createAuraeCA", func(t *testing.T) {
		path := ""
		domainName := "my.domain.com"

		auraeCa, err := CreateAuraeRootCA(path, domainName)
		if err != nil {
			t.Errorf("could create auraeCA")
		}

		// load ca.cert
		cert, _ := pem.Decode([]byte(auraeCa.Certificate))

		crt, err := x509.ParseCertificate(cert.Bytes)
		if err != nil {
			t.Errorf("could parse certificate")
		}
		if crt.Subject.CommonName != "my.domain.com" {
			t.Errorf("certificate does not contain common name")
		}
	})
}