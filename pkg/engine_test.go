package pkg

import (
	"context"
	"errors"
	"testing"
)

var (
	AllTrades    []Trade
	EthUSDTrades []Trade
	BtcUSDTrades []Trade
)

func init() {
	fakeTrades := []Trade{
		{Id: 1, Volume: "11.15505557", Price: "3801.13", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 2, Volume: "2.105034", Price: "3801.24", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 3, Volume: "0.02778985", Price: "3801.33", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 4, Volume: "1.11722945", Price: "3801.42", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 5, Volume: "0.50999715", Price: "3801.68", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 6, Volume: "0.001", Price: "46140.63", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 7, Volume: "0.00195483", Price: "46140.63", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 8, Volume: "0.0021002", Price: "46142.19", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 9, Volume: "0.17887", Price: "3802.1", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 10, Volume: "0.001", Price: "46144.06", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 11, Volume: "0.004415", Price: "46144.06", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 12, Volume: "0.0109", Price: "46144.07", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 13, Volume: "0.00368336", Price: "46144.39", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 14, Volume: "0.00677159", Price: "0.08239", TradeSymbol: "ETH-BTC", ProviderName: "mock", Currency: "USD"},
		{Id: 15, Volume: "0.5", Price: "3801.59", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 16, Volume: "1.13342666", Price: "3801.61", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 17, Volume: "3", Price: "3801.65", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 18, Volume: "0.38", Price: "3801.68", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 19, Volume: "0.00101725", Price: "3784.03", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 20, Volume: "0.01599", Price: "69889.02", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 21, Volume: "0.01548", Price: "69891.71", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 22, Volume: "0.00117477", Price: "69891.81", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 23, Volume: "0.00131033", Price: "69892.10", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 24, Volume: "0.00050798", Price: "3784.18", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 25, Volume: "0.01323", Price: "69893.38", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 26, Volume: "0.0000143", Price: "69893.38", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 27, Volume: "0.07900209", Price: "3784.36", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 28, Volume: "0.01543728", Price: "69894.11", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 29, Volume: "0.11510333", Price: "3784.39", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 30, Volume: "0.00005285", Price: "3784.39", TradeSymbol: "ETH-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 31, Volume: "0.0715337", Price: "69897.48", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		{Id: 32, Volume: "0..0715337", Price: "69897.48", TradeSymbol: "BTC-USD", ProviderName: "mock", Currency: "USD"},
		//{Id: 32, Volume: "0.0715337", Price: "69897.48", TradeSymbol: "BTC-USD", ProviderName: "error", Currency: "USD"},
	}

	AllTrades = fakeTrades
	// TODO: Populate EthUSDTrades and BtcUSDTrades
	for _, t := range AllTrades {
		switch t.TradeSymbol {
		case "BTC-USD":
			BtcUSDTrades = append(BtcUSDTrades, t)
			break
		case "ETH-USD":
			EthUSDTrades = append(EthUSDTrades, t)
			break
		default:
			break
		}
	}
}

func TestInit(t *testing.T) {

}

func TestVWAPComputer_Listen(t *testing.T) {
	stopCtx, cancel := context.WithCancel(context.Background())
	coinbaseFeedMock := &CoinbaseMock{stopListener: cancel}

	type args struct {
		ctx *context.Context
		wsf WebSocketFeed
	}
	tests := []struct {
		name               string
		args               args
		wantConsumedNumber int
		wantWindowSize     int
	}{
		{
			name: "test listener function",
			args: args{
				ctx: &stopCtx,
				wsf: coinbaseFeedMock,
			},
			wantConsumedNumber: 32,
			wantWindowSize:     10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := NewComputerVWAP(tt.wantWindowSize)
			engine.Listen(stopCtx, cancel, tt.args.wsf)

			// TODO: implement me; assert the number of trades consumed and the window size
		})
	}
}

func TestVWAPComputer_Engine(t *testing.T) {
	type fields struct {
		wantWindowSize       int
		wantPriceTimesVolume float64
		wantSumVolume        float64
		wantVWAP             float64
		wantErr              bool
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		trades []Trade
	}{
		{
			name:   "test VWAP computation of BTC-USD",
			args:   args{symbol: "BTC-USD"},
			trades: BtcUSDTrades,
			fields: fields{
				wantVWAP:             66157.70506683017,
				wantSumVolume:        0.15922377,
				wantPriceTimesVolume: 10533.8792152888,
				wantWindowSize:       15,
				wantErr:              false,
			},
		},
		{
			name:   "test VWAP computation of ETH-USD",
			args:   args{symbol: "ETH-USD"},
			trades: EthUSDTrades,
			fields: fields{
				wantVWAP:             3801.143769859362,
				wantSumVolume:        20.303086179999998,
				wantPriceTimesVolume: 77174.9495420247,
				wantWindowSize:       15,
				wantErr:              false,
			},
		},
		{
			name:   "test VWAP computation of ETH-USD window size 5",
			args:   args{symbol: "ETH-USD"},
			trades: EthUSDTrades,
			fields: fields{
				wantVWAP:             3784.375471700442,
				wantSumVolume:        0.19568349999999968,
				wantPriceTimesVolume: 740.5398376164923,
				wantWindowSize:       5,
				wantErr:              false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			v := NewComputerVWAP(tt.fields.wantWindowSize)
			for i := 0; i < len(tt.trades); i++ {
				v.Compute(&tt.trades[i])
			}

			gotPriceTimesVolume := v.priceTimesVolume[tt.args.symbol]
			gotSumVolume := v.sumVolume[tt.args.symbol]
			gotVWAP := v.vWAP[tt.args.symbol]
			gotCount := len(v.trades[tt.args.symbol])

			if gotPriceTimesVolume != tt.fields.wantPriceTimesVolume {
				t.Errorf("VwapComputer.Compute() produces priceTimesVolume = %v, want %v", gotPriceTimesVolume, tt.fields.wantPriceTimesVolume)
			}
			if gotSumVolume != tt.fields.wantSumVolume {
				t.Errorf("VwapComputer.Compute() makes prduces sumVolume = %v, want %v", gotSumVolume, tt.fields.wantSumVolume)
			}
			if gotVWAP != tt.fields.wantVWAP {
				t.Errorf("VwapComputer.Compute() makes prduces vWAP = %v, want %v", gotVWAP, tt.fields.wantVWAP)
			}
			if gotCount != tt.fields.wantWindowSize {
				t.Errorf("VwapComputer.Compute() makes count of trades count of = %v, want %v", gotCount, tt.fields.wantWindowSize)
			}
		})
	}
}

type CoinbaseMock struct {
	stopListener context.CancelFunc
}

func (c *CoinbaseMock) Subscribe() error { return nil }

func (c *CoinbaseMock) TurnOff() error { return nil }

func (c *CoinbaseMock) Read() (*Trade, error) {
	if count := len(AllTrades); count <= 0 {
		c.stopListener()
		return &Trade{}, nil
	}
	// pop one trade from slice
	op := AllTrades[0]
	AllTrades = AllTrades[1:]
	if op.ProviderName == "error" {
		return nil, errors.New("unexpected websocket error")
	}
	return &op, nil
}
