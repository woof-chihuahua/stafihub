package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	rvotetypes "github.com/stafihub/stafihub/x/rvote/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetEraUnbondLimit{}, "ledger/SetEraUnbondLimit", nil)
	cdc.RegisterConcrete(&MsgSetPoolDetail{}, "ledger/SetPoolDetail", nil)
	cdc.RegisterConcrete(&MsgSetLeastBond{}, "ledger/SetLeastBond", nil)
	cdc.RegisterConcrete(&MsgClearCurrentEraSnapShots{}, "ledger/ClearCurrentEraSnapShots", nil)
	cdc.RegisterConcrete(&MsgSetStakingRewardCommission{}, "ledger/SetStakingRewardCommission", nil)
	cdc.RegisterConcrete(&MsgSetProtocolFeeReceiver{}, "ledger/SetProtocolFeeReceiver", nil)
	cdc.RegisterConcrete(&MsgSetUnbondRelayFee{}, "ledger/SetUnbondRelayFee", nil)
	cdc.RegisterConcrete(&MsgLiquidityUnbond{}, "ledger/LiquidityUnbond", nil)
	cdc.RegisterConcrete(&MsgSetUnbondCommission{}, "ledger/SetUnbondCommission", nil)

	cdc.RegisterConcrete(&SetChainEraProposal{}, "ledger/SetChainEraProposal", nil)
	cdc.RegisterConcrete(&BondReportProposal{}, "ledger/BondReportProposal", nil)
	cdc.RegisterConcrete(&BondAndReportActiveProposal{}, "ledger/BondAndReportActiveProposal", nil)
	cdc.RegisterConcrete(&ActiveReportProposal{}, "ledger/ActiveReportProposal", nil)
	cdc.RegisterConcrete(&WithdrawReportProposal{}, "ledger/WithdrawReportProposal", nil)
	cdc.RegisterConcrete(&TransferReportProposal{}, "ledger/TransferReportProposal", nil)
	cdc.RegisterConcrete(&ExecuteBondProposal{}, "ledger/ExecuteBondProposal", nil)

	cdc.RegisterConcrete(&MsgSubmitSignature{}, "ledger/SubmitSignature", nil)
	cdc.RegisterConcrete(&MsgSetRParams{}, "ledger/SetRParams", nil)
	cdc.RegisterConcrete(&MsgOnboard{}, "ledger/Onboard", nil)
	cdc.RegisterConcrete(&MsgSetRValidatorIndicator{}, "ledger/SetRValidatorIndicator", nil)
	cdc.RegisterConcrete(&MsgSetRelayFeeReceiver{}, "ledger/SetRelayFeeReceiver", nil)
	cdc.RegisterConcrete(&MsgSetRelayGasPrice{}, "ledger/SetRelayGasPrice", nil)
	cdc.RegisterConcrete(&MsgSetEraSeconds{}, "ledger/SetEraSeconds", nil)
	cdc.RegisterConcrete(&MsgRmBondedPool{}, "ledger/RmBondedPool", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetEraUnbondLimit{},
		&MsgSetPoolDetail{},
		&MsgSetLeastBond{},
		&MsgClearCurrentEraSnapShots{},
		&MsgSetStakingRewardCommission{},
		&MsgSetProtocolFeeReceiver{},
	)

	registry.RegisterImplementations(
		(*rvotetypes.Content)(nil),
		&SetChainEraProposal{},
		&BondReportProposal{},
		&BondAndReportActiveProposal{},
		&ActiveReportProposal{},
		&WithdrawReportProposal{},
		&TransferReportProposal{},
		&ExecuteBondProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitSignature{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetRParams{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOnboard{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetRValidatorIndicator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetRelayFeeReceiver{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetRelayGasPrice{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetEraSeconds{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmBondedPool{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
