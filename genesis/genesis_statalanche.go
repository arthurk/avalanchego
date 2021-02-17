// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/ava-labs/avalanchego/utils/units"
)

var (
	statalancheGenesisConfigJSON = `{
		"networkID": 115110116,
		"allocations": [
			{
				"avaxAddr": "X-statalanche15faat649drfma0x5v3jrglswjyt0d043echgy8",
				"ethAddr": "0xe027688a57c4A6Fb2708343cF330aaeB8fe594bb",
				"initialAmount": 40000000000000000,
				"unlockSchedule": [
					{
						"amount": 40000000000000000
					}
				]
			},
			{
				"ethAddr": "0xe027688a57c4A6Fb2708343cF330aaeB8fe594bb",
				"avaxAddr": "X-statalanche1v2j6lcmqlrqanpks6x54pptpyaqgdkcd6gkewv",
				"initialAmount": 40000000000000000,
				"unlockSchedule": [
					{
						"amount": 40000000000000000
					}
				]
			},
			{
				"ethAddr": "0xe027688a57c4A6Fb2708343cF330aaeB8fe594bb",
				"avaxAddr": "X-statalanche1xwhych03n8h3hmtcxqvkz9l2uaaa3dw20zpc9r",
				"initialAmount": 40000000000000000,
				"unlockSchedule": [
					{
						"amount": 40000000000000000
					}
				]
			},
			{
				"ethAddr": "0xe027688a57c4A6Fb2708343cF330aaeB8fe594bb",
				"avaxAddr": "X-statalanche1t3vej6wn3v9jn8f9ke0z2vjsa38heqxhwlfqsj",
				"initialAmount": 100000000000000000
			},
			{
				"ethAddr": "0xe027688a57c4A6Fb2708343cF330aaeB8fe594bb",
				"avaxAddr": "X-statalanche1vw8w9lxp3t3dz92vm9rvutqwk0zcvsx7jnxvdq",
				"initialAmount": 100000000000000000
			}
		],
		"startTime": 1613590094,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [
			"X-statalanche15faat649drfma0x5v3jrglswjyt0d043echgy8",
			"X-statalanche1v2j6lcmqlrqanpks6x54pptpyaqgdkcd6gkewv",
			"X-statalanche1xwhych03n8h3hmtcxqvkz9l2uaaa3dw20zpc9r"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-E349V27puxyemQTc6QDDGaLD1kEF1GJfa",
				"rewardAddress": "X-statalanche15faat649drfma0x5v3jrglswjyt0d043echgy8",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-Pvfe5vYMUMecgXWcbtmSjqVxcisQJuHty",
				"rewardAddress": "X-statalanche1v2j6lcmqlrqanpks6x54pptpyaqgdkcd6gkewv",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-PCe2PQb5wwXT7JUsFQjrqb8bVhRinnKoW",
				"rewardAddress": "X-statalanche1xwhych03n8h3hmtcxqvkz9l2uaaa3dw20zpc9r",
				"delegationFee": 200000
			}
		],
		"cChainGenesis": "{\"config\":{\"chainId\":13375,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"0100000000000000000000000000000000000000\":{\"code\":\"0x7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80631e010439146042578063b6510bb314606e575b600080fd5b605c60048036036020811015605657600080fd5b503560b1565b60408051918252519081900360200190f35b818015607957600080fd5b5060af60048036036080811015608e57600080fd5b506001600160a01b03813516906020810135906040810135906060013560b6565b005b30cd90565b836001600160a01b031681836108fc8690811502906040516000604051808303818888878c8acf9550505050505015801560f4573d6000803e3d6000fd5b505050505056fea26469706673582212201eebce970fe3f5cb96bf8ac6ba5f5c133fc2908ae3dcd51082cfee8f583429d064736f6c634300060a0033\",\"balance\":\"0x0\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "It is all about the State of Us"
	}`

	// StatalancheParams are the params used for local networks
	StatalancheParams = Params{
		TxFee:                units.MilliAvax,
		CreationTxFee:        10 * units.MilliAvax,
		UptimeRequirement:    .6, // 60%
		MinValidatorStake:    1 * units.Avax,
		MaxValidatorStake:    3 * units.MegaAvax,
		MinDelegatorStake:    1 * units.Avax,
		MinDelegationFee:     20000, // 2%
		MinStakeDuration:     24 * time.Hour,
		MaxStakeDuration:     365 * 24 * time.Hour,
		StakeMintingPeriod:   365 * 24 * time.Hour,
		EpochFirstTransition: time.Unix(1607626800, 0),
		EpochDuration:        30 * time.Minute,
		ApricotPhase0Time:    time.Date(2020, 12, 5, 5, 00, 0, 0, time.UTC),
	}
)
