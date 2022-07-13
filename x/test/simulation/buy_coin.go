package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sujine/test/x/test/keeper"
	"github.com/sujine/test/x/test/types"
)

func SimulateMsgBuyCoin(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		_ = &types.MsgMintCoin{
			Creator: simAccount.Address.String(),
			Denom:   "ttt",
			Amount:  100000,
		}

		_ = &types.MsgRegisterCoin{
			Creator: simAccount.Address.String(),
			Symbol:  "ttt",
			Price:   200,
			Amount:  500,
		}

		msg := &types.MsgBuyCoin{
			Creator: simAccount.Address.String(),
			Symbol:  "ttt",
			Amount:  200,
		}

		// TODO: Handling the BuyCoin simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "BuyCoin simulation not implemented"), nil, nil
	}
}
