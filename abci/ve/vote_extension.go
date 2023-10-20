package ve

import (
	"context"
	"time"

	"cosmossdk.io/log"
	cometabci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/skip-mev/slinky/abci/ve/types"
	"github.com/skip-mev/slinky/service"
)

// VoteExtensionHandler is a handler that extends a vote with the oracle's
// current price feed. In the case where oracle data is unable to be fetched
// or correctly marshalled, the handler will return an empty vote extension to
// ensure liveliness.
type VoteExtensionHandler struct {
	logger log.Logger

	// oracleClient is the oracle client (remote or local) that is responsible for fetching prices
	//
	// TODO: Add a separate interface just for the client.
	oracleClient service.OracleService

	// timeout is the maximum amount of time to wait for the oracle to respond
	// to a price request.
	timeout time.Duration
}

// NewVoteExtensionHandler returns a new VoteExtensionHandler.
func NewVoteExtensionHandler(logger log.Logger, oracleClient service.OracleService, timeout time.Duration) *VoteExtensionHandler {
	return &VoteExtensionHandler{
		logger:       logger,
		oracleClient: oracleClient,
		timeout:      timeout,
	}
}

// ExtendVoteHandler returns a handler that extends a vote with the oracle's
// current price feed. In the case where oracle data is unable to be fetched
// or correctly marshalled, the handler will return an empty vote extension to
// ensure liveliness.
func (h *VoteExtensionHandler) ExtendVoteHandler() sdk.ExtendVoteHandler {
	return func(ctx sdk.Context, req *cometabci.RequestExtendVote) (resp *cometabci.ResponseExtendVote, err error) {
		// Catch any panic that occurs in the oracle request.
		defer func() {
			if r := recover(); r != nil {
				h.logger.Error(
					"recovered from panic in ExtendVoteHandler",
					"err", r,
				)

				resp, err = &cometabci.ResponseExtendVote{VoteExtension: []byte{}}, nil
			}
		}()

		// Create a context with a timeout to ensure we do not wait forever for the oracle
		// to respond.
		reqCtx, cancel := context.WithTimeout(ctx.Context(), h.timeout)
		defer cancel()

		// To ensure liveliness, we return a vote even if the oracle is not running
		// or if the oracle returns a bad response.
		oracleResp, err := h.oracleClient.Prices(reqCtx, &service.QueryPricesRequest{})
		if err != nil {
			h.logger.Error(
				"failed to retrieve oracle prices for vote extension; returning empty vote extension",
				"height", req.Height,
				"ctx_err", reqCtx.Err(),
				"err", err,
			)

			return &cometabci.ResponseExtendVote{VoteExtension: []byte{}}, nil
		}

		// If we get no response, we return an empty vote extension.
		if oracleResp == nil {
			h.logger.Error(
				"oracle returned nil prices for vote extension; returning empty vote extension",
				"height", req.Height,
			)

			return &cometabci.ResponseExtendVote{VoteExtension: []byte{}}, nil
		}

		voteExt := abcitypes.OracleVoteExtension{
			Height:    req.Height,
			Prices:    oracleResp.Prices,
			Timestamp: oracleResp.Timestamp,
		}

		bz, err := voteExt.Marshal()
		if err != nil {
			h.logger.Error(
				"failed to marshal vote extension; returning empty vote extension",
				"height", req.Height,
				"err", err,
			)

			return &cometabci.ResponseExtendVote{VoteExtension: []byte{}}, nil
		}

		h.logger.Info(
			"extending vote with oracle prices",
			"prices", oracleResp.Prices,
			"vote_extension_height", voteExt.Height,
			"vote_extension_timestamp", voteExt.Timestamp,
			"req_height", req.Height,
		)

		return &cometabci.ResponseExtendVote{VoteExtension: bz}, nil
	}
}

// VerifyVoteExtensionHandler returns a handler that verifies the vote extension provided by
// a validator is valid. In the case when the vote extension is empty, we return ACCEPT. This means
// that the validator may have been unable to fetch prices from the oracle and is voting an empty vote extension.
// We reject any vote extensions that are not empty and fail to unmarshal or contain invalid prices.
func (h *VoteExtensionHandler) VerifyVoteExtensionHandler() sdk.VerifyVoteExtensionHandler {
	return func(ctx sdk.Context, req *cometabci.RequestVerifyVoteExtension) (*cometabci.ResponseVerifyVoteExtension, error) {
		voteExtension := req.VoteExtension

		if err := ValidateOracleVoteExtension(voteExtension, req.Height); err != nil {
			h.logger.Error(
				"failed to validate vote extension",
				"height", req.Height,
				"err", err,
			)

			return &cometabci.ResponseVerifyVoteExtension{Status: cometabci.ResponseVerifyVoteExtension_REJECT}, err
		}

		h.logger.Info(
			"validated vote extension",
			"height", req.Height,
		)

		return &cometabci.ResponseVerifyVoteExtension{Status: cometabci.ResponseVerifyVoteExtension_ACCEPT}, nil
	}
}