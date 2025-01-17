// VulcanizeDB
// Copyright © 2018 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package integration

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/geth"
	"github.com/vulcanize/vulcanizedb/pkg/geth/client"
	vRpc "github.com/vulcanize/vulcanizedb/pkg/geth/converters/rpc"
	"github.com/vulcanize/vulcanizedb/pkg/geth/node"
	"github.com/vulcanize/vulcanizedb/test_config"
)

var _ = Describe("Rewards calculations", func() {

	It("calculates a block reward for a real block", func() {
		rawRpcClient, err := rpc.Dial(test_config.InfuraClient.IPCPath)
		Expect(err).NotTo(HaveOccurred())
		rpcClient := client.NewRpcClient(rawRpcClient, test_config.InfuraClient.IPCPath)
		ethClient := ethclient.NewClient(rawRpcClient)
		blockChainClient := client.NewEthClient(ethClient)
		node := node.MakeNode(rpcClient)
		transactionConverter := vRpc.NewRpcTransactionConverter(ethClient)
		blockChain := geth.NewBlockChain(blockChainClient, rpcClient, node, transactionConverter)
		block, err := blockChain.GetBlockByNumber(1071819)
		Expect(err).ToNot(HaveOccurred())
		Expect(block.Reward).To(Equal(5.31355))
	})

	It("calculates an uncle reward for a real block", func() {
		rawRpcClient, err := rpc.Dial(test_config.InfuraClient.IPCPath)
		Expect(err).NotTo(HaveOccurred())
		rpcClient := client.NewRpcClient(rawRpcClient, test_config.InfuraClient.IPCPath)
		ethClient := ethclient.NewClient(rawRpcClient)
		blockChainClient := client.NewEthClient(ethClient)
		node := node.MakeNode(rpcClient)
		transactionConverter := vRpc.NewRpcTransactionConverter(ethClient)
		blockChain := geth.NewBlockChain(blockChainClient, rpcClient, node, transactionConverter)
		block, err := blockChain.GetBlockByNumber(1071819)
		Expect(err).ToNot(HaveOccurred())
		Expect(block.UnclesReward).To(Equal(6.875))
	})

})
