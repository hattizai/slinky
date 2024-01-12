package handlers_test

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	wserrors "github.com/skip-mev/slinky/providers/base/websocket/errors"
	"github.com/skip-mev/slinky/providers/base/websocket/handlers"
	handlermocks "github.com/skip-mev/slinky/providers/base/websocket/handlers/mocks"
	providertypes "github.com/skip-mev/slinky/providers/types"
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var (
	logger  = zap.NewExample()
	btcusd  = oracletypes.NewCurrencyPair("BTC", "USD")
	ethusd  = oracletypes.NewCurrencyPair("ETH", "USD")
	atomusd = oracletypes.NewCurrencyPair("ATOM", "USD")

	websocketURL = "ws://localhost:8080"
	name         = "sirmoggintonwebsocket"
)

func TestWebSocketQueryHandler(t *testing.T) {
	testCases := []struct {
		name        string
		connHandler func() handlers.WebSocketConnHandler
		dataHandler func() handlers.WebSocketDataHandler[oracletypes.CurrencyPair, *big.Int]
		ids         []oracletypes.CurrencyPair
		responses   providertypes.GetResponse[oracletypes.CurrencyPair, *big.Int]
	}{
		{
			name: "fails to dial the websocket",
			connHandler: func() handlers.WebSocketConnHandler {
				connHandler := handlermocks.NewWebSocketConnHandler(t)

				connHandler.On("Dial", mock.Anything).Return(fmt.Errorf("no rizz alert")).Once()

				return connHandler
			},
			dataHandler: func() handlers.WebSocketDataHandler[oracletypes.CurrencyPair, *big.Int] {
				dataHandler := handlermocks.NewWebSocketDataHandler[oracletypes.CurrencyPair, *big.Int](t)

				dataHandler.On("Name").Return(name).Once()
				dataHandler.On("URL").Return(websocketURL).Once()

				return dataHandler
			},
			ids: []oracletypes.CurrencyPair{btcusd},
			responses: providertypes.GetResponse[oracletypes.CurrencyPair, *big.Int]{
				UnResolved: map[oracletypes.CurrencyPair]error{
					btcusd: wserrors.ErrDial,
				},
			},
		},
		{
			name: "fails to create subscriptions",
			connHandler: func() handlers.WebSocketConnHandler {
				connHandler := handlermocks.NewWebSocketConnHandler(t)

				connHandler.On("Dial", mock.Anything).Return(nil).Once()

				return connHandler
			},
			dataHandler: func() handlers.WebSocketDataHandler[oracletypes.CurrencyPair, *big.Int] {
				dataHandler := handlermocks.NewWebSocketDataHandler[oracletypes.CurrencyPair, *big.Int](t)

				dataHandler.On("Name").Return(name).Once()
				dataHandler.On("URL").Return(websocketURL).Once()
				dataHandler.On("CreateMessage", mock.Anything).Return(nil, fmt.Errorf("no rizz alert")).Once()

				return dataHandler
			},
			ids: []oracletypes.CurrencyPair{btcusd},
			responses: providertypes.GetResponse[oracletypes.CurrencyPair, *big.Int]{
				UnResolved: map[oracletypes.CurrencyPair]error{
					btcusd: wserrors.ErrCreateMessage,
				},
			},
		},
		{
			name: "fails to write to the websocket on initial subscription",
			connHandler: func() handlers.WebSocketConnHandler {
				connHandler := handlermocks.NewWebSocketConnHandler(t)

				connHandler.On("Dial", mock.Anything).Return(nil).Once()
				connHandler.On("Write", mock.Anything).Return(fmt.Errorf("no rizz alert")).Once()

				return connHandler
			},
			dataHandler: func() handlers.WebSocketDataHandler[oracletypes.CurrencyPair, *big.Int] {
				dataHandler := handlermocks.NewWebSocketDataHandler[oracletypes.CurrencyPair, *big.Int](t)

				dataHandler.On("Name").Return(name).Once()
				dataHandler.On("URL").Return(websocketURL).Once()
				dataHandler.On("CreateMessage", mock.Anything).Return([]byte("gib me money"), nil).Once()

				return dataHandler
			},
			ids: []oracletypes.CurrencyPair{btcusd},
			responses: providertypes.GetResponse[oracletypes.CurrencyPair, *big.Int]{
				UnResolved: map[oracletypes.CurrencyPair]error{
					btcusd: wserrors.ErrWrite,
				},
			},
		},
		{
			name: "fails to read from the websocket",
			connHandler: func() handlers.WebSocketConnHandler {
				connHandler := handlermocks.NewWebSocketConnHandler(t)

				connHandler.On("Dial", mock.Anything).Return(nil).Once()
				connHandler.On("Write", mock.Anything).Return(nil).Once()
				connHandler.On("Read").Return(nil, fmt.Errorf("no rizz alert")).Maybe().After(time.Second)
				connHandler.On("Close").Return(nil).Once()

				return connHandler
			},
			dataHandler: func() handlers.WebSocketDataHandler[oracletypes.CurrencyPair, *big.Int] {
				dataHandler := handlermocks.NewWebSocketDataHandler[oracletypes.CurrencyPair, *big.Int](t)

				dataHandler.On("Name").Return(name).Once()
				dataHandler.On("URL").Return(websocketURL).Once()
				dataHandler.On("CreateMessage", mock.Anything).Return([]byte("gib me money"), nil).Once()

				return dataHandler
			},

			ids:       []oracletypes.CurrencyPair{btcusd},
			responses: providertypes.GetResponse[oracletypes.CurrencyPair, *big.Int]{},
		},
		{
			name: "fails to parse the response from the websocket",
			connHandler: func() handlers.WebSocketConnHandler {
				connHandler := handlermocks.NewWebSocketConnHandler(t)

				connHandler.On("Dial", mock.Anything).Return(nil).Once()
				connHandler.On("Write", mock.Anything).Return(nil).Once()
				connHandler.On("Read").Return([]byte("gib me money"), nil).Maybe().After(time.Second)
				connHandler.On("Close").Return(nil).Once()

				return connHandler
			},
			dataHandler: func() handlers.WebSocketDataHandler[oracletypes.CurrencyPair, *big.Int] {
				dataHandler := handlermocks.NewWebSocketDataHandler[oracletypes.CurrencyPair, *big.Int](t)

				dataHandler.On("Name").Return(name).Once()
				dataHandler.On("URL").Return(websocketURL).Once()
				dataHandler.On("CreateMessage", mock.Anything).Return([]byte("gib me money"), nil).Once()
				dataHandler.On("HandleMessage", mock.Anything).Return(
					providertypes.NewGetResponse[oracletypes.CurrencyPair, *big.Int](nil, nil),
					nil,
					fmt.Errorf("no rizz alert"),
				).Maybe()

				return dataHandler
			},

			ids:       []oracletypes.CurrencyPair{btcusd},
			responses: providertypes.GetResponse[oracletypes.CurrencyPair, *big.Int]{},
		},
		{
			name: "pseudo heart beat update message with no response",
			connHandler: func() handlers.WebSocketConnHandler {
				connHandler := handlermocks.NewWebSocketConnHandler(t)

				connHandler.On("Dial", mock.Anything).Return(nil).Once()
				connHandler.On("Write", mock.Anything).Return(nil).Maybe()
				connHandler.On("Read").Return([]byte("gib me money"), nil).Maybe().After(time.Second)
				connHandler.On("Close").Return(nil).Once()

				return connHandler
			},
			dataHandler: func() handlers.WebSocketDataHandler[oracletypes.CurrencyPair, *big.Int] {
				dataHandler := handlermocks.NewWebSocketDataHandler[oracletypes.CurrencyPair, *big.Int](t)

				dataHandler.On("Name").Return(name).Once()
				dataHandler.On("URL").Return(websocketURL).Once()
				dataHandler.On("CreateMessage", mock.Anything).Return([]byte("gib me money"), nil).Once()
				dataHandler.On("HandleMessage", mock.Anything).Return(
					providertypes.NewGetResponse[oracletypes.CurrencyPair, *big.Int](nil, nil),
					[]byte("hearb eat"),
					nil,
				).Maybe()

				return dataHandler
			},

			ids:       []oracletypes.CurrencyPair{btcusd},
			responses: providertypes.GetResponse[oracletypes.CurrencyPair, *big.Int]{},
		},
		{
			name: "fails to send the update message to the websocket",
			connHandler: func() handlers.WebSocketConnHandler {
				connHandler := handlermocks.NewWebSocketConnHandler(t)

				connHandler.On("Dial", mock.Anything).Return(nil).Once()
				connHandler.On("Read").Return([]byte("gib me money"), nil).Maybe().After(time.Second)
				connHandler.On("Close").Return(nil).Once()
				connHandler.On("Write", mock.Anything).Return(nil).Once()
				connHandler.On("Write", mock.Anything).Return(fmt.Errorf("no rizz alert")).Maybe()

				return connHandler
			},
			dataHandler: func() handlers.WebSocketDataHandler[oracletypes.CurrencyPair, *big.Int] {
				dataHandler := handlermocks.NewWebSocketDataHandler[oracletypes.CurrencyPair, *big.Int](t)

				dataHandler.On("Name").Return(name).Once()
				dataHandler.On("URL").Return(websocketURL).Once()
				dataHandler.On("CreateMessage", mock.Anything).Return([]byte("gib me money"), nil).Once()
				dataHandler.On("HandleMessage", mock.Anything).Return(
					providertypes.NewGetResponse[oracletypes.CurrencyPair, *big.Int](nil, nil),
					[]byte("hearb eat"),
					nil,
				).Maybe()

				return dataHandler
			},
			ids:       []oracletypes.CurrencyPair{btcusd},
			responses: providertypes.GetResponse[oracletypes.CurrencyPair, *big.Int]{},
		},
		{
			name: "fails to close the websocket",
			connHandler: func() handlers.WebSocketConnHandler {
				connHandler := handlermocks.NewWebSocketConnHandler(t)

				connHandler.On("Dial", mock.Anything).Return(nil).Once()
				connHandler.On("Read").Return([]byte("gib me money"), nil).Maybe().After(time.Second)
				connHandler.On("Close").Return(fmt.Errorf("no rizz alert")).Once()
				connHandler.On("Write", mock.Anything).Return(nil).Maybe()

				return connHandler
			},
			dataHandler: func() handlers.WebSocketDataHandler[oracletypes.CurrencyPair, *big.Int] {
				dataHandler := handlermocks.NewWebSocketDataHandler[oracletypes.CurrencyPair, *big.Int](t)

				dataHandler.On("Name").Return(name).Once()
				dataHandler.On("URL").Return(websocketURL).Once()
				dataHandler.On("CreateMessage", mock.Anything).Return([]byte("gib me money"), nil).Once()
				dataHandler.On("HandleMessage", mock.Anything).Return(
					providertypes.NewGetResponse[oracletypes.CurrencyPair, *big.Int](nil, nil),
					[]byte("hearb eat"),
					nil,
				).Maybe()

				return dataHandler
			},
			ids:       []oracletypes.CurrencyPair{btcusd},
			responses: providertypes.GetResponse[oracletypes.CurrencyPair, *big.Int]{},
		},
		{
			name: "returns a single response with no update message",
			connHandler: func() handlers.WebSocketConnHandler {
				connHandler := handlermocks.NewWebSocketConnHandler(t)

				connHandler.On("Dial", mock.Anything).Return(nil).Once()
				connHandler.On("Read").Return([]byte("gib me money"), nil).Maybe().After(time.Second)
				connHandler.On("Close").Return(nil).Once()
				connHandler.On("Write", mock.Anything).Return(nil).Once()

				return connHandler
			},
			dataHandler: func() handlers.WebSocketDataHandler[oracletypes.CurrencyPair, *big.Int] {
				dataHandler := handlermocks.NewWebSocketDataHandler[oracletypes.CurrencyPair, *big.Int](t)

				dataHandler.On("Name").Return(name).Once()
				dataHandler.On("URL").Return(websocketURL).Once()
				dataHandler.On("CreateMessage", mock.Anything).Return([]byte("gib me money"), nil).Once()

				resolved := map[oracletypes.CurrencyPair]providertypes.Result[*big.Int]{
					btcusd: {
						Value: big.NewInt(100),
					},
				}
				response := providertypes.NewGetResponse[oracletypes.CurrencyPair, *big.Int](resolved, nil)
				dataHandler.On("HandleMessage", mock.Anything).Return(
					response,
					nil,
					nil,
				).Maybe()

				return dataHandler
			},
			ids: []oracletypes.CurrencyPair{btcusd},
			responses: providertypes.GetResponse[oracletypes.CurrencyPair, *big.Int]{
				Resolved: map[oracletypes.CurrencyPair]providertypes.Result[*big.Int]{
					btcusd: {
						Value: big.NewInt(100),
					},
				},
			},
		},
		{
			name: "returns a single response with an update message",
			connHandler: func() handlers.WebSocketConnHandler {
				connHandler := handlermocks.NewWebSocketConnHandler(t)

				connHandler.On("Dial", mock.Anything).Return(nil).Once()
				connHandler.On("Read").Return([]byte("gib me money"), nil).Maybe().After(time.Second)
				connHandler.On("Close").Return(nil).Once()
				connHandler.On("Write", mock.Anything).Return(nil).Maybe()

				return connHandler
			},
			dataHandler: func() handlers.WebSocketDataHandler[oracletypes.CurrencyPair, *big.Int] {
				dataHandler := handlermocks.NewWebSocketDataHandler[oracletypes.CurrencyPair, *big.Int](t)
				dataHandler.On("Name").Return(name).Once()
				dataHandler.On("URL").Return(websocketURL).Once()
				dataHandler.On("CreateMessage", mock.Anything).Return([]byte("gib me money"), nil).Once()

				resolved := map[oracletypes.CurrencyPair]providertypes.Result[*big.Int]{
					btcusd: {
						Value: big.NewInt(100),
					},
				}
				response := providertypes.NewGetResponse[oracletypes.CurrencyPair, *big.Int](resolved, nil)
				dataHandler.On("HandleMessage", mock.Anything).Return(
					response,
					[]byte("hearb eat"),
					nil,
				).Maybe()

				return dataHandler
			},
			ids: []oracletypes.CurrencyPair{btcusd},
			responses: providertypes.GetResponse[oracletypes.CurrencyPair, *big.Int]{
				Resolved: map[oracletypes.CurrencyPair]providertypes.Result[*big.Int]{
					btcusd: {
						Value: big.NewInt(100),
					},
				},
			},
		},
		{
			name: "returns multiple responses with no update message",
			connHandler: func() handlers.WebSocketConnHandler {
				connHandler := handlermocks.NewWebSocketConnHandler(t)

				connHandler.On("Dial", mock.Anything).Return(nil).Once()
				connHandler.On("Read").Return([]byte("gib me money"), nil).Maybe().After(time.Second)
				connHandler.On("Close").Return(nil).Once()
				connHandler.On("Write", mock.Anything).Return(nil).Once()

				return connHandler
			},
			dataHandler: func() handlers.WebSocketDataHandler[oracletypes.CurrencyPair, *big.Int] {
				dataHandler := handlermocks.NewWebSocketDataHandler[oracletypes.CurrencyPair, *big.Int](t)

				dataHandler.On("Name").Return(name).Once()
				dataHandler.On("URL").Return(websocketURL).Once()
				dataHandler.On("CreateMessage", mock.Anything).Return([]byte("gib me money"), nil).Once()

				resolved := map[oracletypes.CurrencyPair]providertypes.Result[*big.Int]{
					btcusd: {
						Value: big.NewInt(100),
					},
				}
				resolved2 := map[oracletypes.CurrencyPair]providertypes.Result[*big.Int]{
					ethusd: {
						Value: big.NewInt(200),
					},
				}
				unresolved := map[oracletypes.CurrencyPair]error{
					atomusd: wserrors.ErrHandleMessage,
				}

				response1 := providertypes.NewGetResponse[oracletypes.CurrencyPair, *big.Int](resolved, nil)
				dataHandler.On("HandleMessage", mock.Anything).Return(
					response1,
					nil,
					nil,
				).Once()

				response2 := providertypes.NewGetResponse[oracletypes.CurrencyPair, *big.Int](resolved2, nil)
				dataHandler.On("HandleMessage", mock.Anything).Return(
					response2,
					nil,
					nil,
				).Once()

				response3 := providertypes.NewGetResponse[oracletypes.CurrencyPair, *big.Int](nil, unresolved)
				dataHandler.On("HandleMessage", mock.Anything).Return(
					response3,
					nil,
					nil,
				).Maybe()

				return dataHandler
			},
			ids: []oracletypes.CurrencyPair{btcusd, ethusd},
			responses: providertypes.GetResponse[oracletypes.CurrencyPair, *big.Int]{
				Resolved: map[oracletypes.CurrencyPair]providertypes.Result[*big.Int]{
					btcusd: {
						Value: big.NewInt(100),
					},
					ethusd: {
						Value: big.NewInt(200),
					},
				},
				UnResolved: map[oracletypes.CurrencyPair]error{
					atomusd: wserrors.ErrHandleMessage,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			handler, err := handlers.NewWebSocketQueryHandler[oracletypes.CurrencyPair, *big.Int](
				logger,
				tc.dataHandler(),
				tc.connHandler(),
			)
			require.NoError(t, err)

			responseCh := make(chan providertypes.GetResponse[oracletypes.CurrencyPair, *big.Int], 20)
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				handler.Start(ctx, tc.ids, responseCh)
				cancel()
				close(responseCh)
			}()

			expectedResponses := tc.responses
			seenResponses := make(map[oracletypes.CurrencyPair]bool)
			for resp := range responseCh {
				for id, result := range resp.Resolved {
					if _, ok := seenResponses[id]; ok {
						continue
					}

					require.Equal(t, expectedResponses.Resolved[id], result)
					delete(expectedResponses.Resolved, id)
					seenResponses[id] = true
				}

				for id, err := range resp.UnResolved {
					if _, ok := seenResponses[id]; ok {
						continue
					}

					require.True(t, errors.Is(err, expectedResponses.UnResolved[id]))
					delete(expectedResponses.UnResolved, id)
					seenResponses[id] = true
				}
			}

			// Ensure all responses are account for.
			require.Empty(t, expectedResponses.Resolved)
			require.Empty(t, expectedResponses.UnResolved)
		})
	}
}
