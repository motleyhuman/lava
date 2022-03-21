package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/lavanet/lava/testutil/keeper"
	"github.com/lavanet/lava/testutil/nullify"
	"github.com/lavanet/lava/x/user/types"
)

func TestBlockDeadlineForCallbackQuery(t *testing.T) {
	keeper, ctx := keepertest.UserKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestBlockDeadlineForCallback(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetBlockDeadlineForCallbackRequest
		response *types.QueryGetBlockDeadlineForCallbackResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetBlockDeadlineForCallbackRequest{},
			response: &types.QueryGetBlockDeadlineForCallbackResponse{BlockDeadlineForCallback: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.BlockDeadlineForCallback(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
