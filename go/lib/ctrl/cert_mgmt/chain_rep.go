// Copyright 2017 ETH Zurich
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cert_mgmt

import (
	"fmt"

	"github.com/netsec-ethz/scion/go/lib/common"
	"github.com/netsec-ethz/scion/go/lib/crypto/cert"
	"github.com/netsec-ethz/scion/go/proto"
)

var _ proto.Cerealizable = (*ChainRep)(nil)

type ChainRep struct {
	RawChain common.RawBytes `capnp:"chain"`
}

func (c *ChainRep) Chain() (*cert.Chain, error) {
	return cert.ChainFromRaw(c.RawChain, true)
}

func (c *ChainRep) ProtoId() proto.ProtoIdType {
	return proto.CertChainRep_TypeID
}

func (c *ChainRep) String() string {
	chain, err := c.Chain()
	if err != nil {
		return fmt.Sprintf("Invalid CertificateChain: %v", err)
	}
	return chain.String()
}
