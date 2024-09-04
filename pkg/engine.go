package pkg

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

func NewComputerVWAP(windowSize int) *ComputerVWAP {
	return &ComputerVWAP{
		trades:           map[string][]*Trade{}, // key: parity symbol (e.g. BTC-USD, ETH-USD etc)
		sumVolume:        map[string]float64{},
		priceTimesVolume: map[string]float64{},
		vWAP:             map[string]float64{},
		windowSize:       windowSize, // assume window size is > 0
	}
}

type ComputerVWAP struct {
	trades           map[string][]*Trade
	sumVolume        map[string]float64
	priceTimesVolume map[string]float64
	vWAP             map[string]float64
	windowSize       int
}

func (v *ComputerVWAP) Listen(_ctx context.Context, _cancelFunc context.CancelFunc, wsf WebSocketFeed) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	// Subscribe to socket
	// TODO Subscribe and listen to consume the socket feed
	if err := wsf.Subscribe(); err != nil {
		log.Fatal(err)
	}
	// defer wsf.TurnOff() // log out error
	for {
		select {
		case <-_ctx.Done():
			fmt.Println("cancelling...")
			wsf.TurnOff()
			return
		default:
			//Consume the socket
			trade, err := wsf.Read()
			if err != nil {
				// _cancelFunc()
				log.Fatal(err)
			}

			// single compute example
			v.Compute(trade)

			log.Printf(
				"Symbol: %s Trade Sum:%3d VWAP: %s %.2f\n",
				trade.TradeSymbol, len(v.trades[trade.TradeSymbol]), trade.Currency, v.vWAP[trade.TradeSymbol],
			)
			time.Sleep(time.Second * 1)
		}
	}
}

// Compute does the main calculation formula of VWAP price.
// https://github.com/shopspring/decimal
func (v *ComputerVWAP) Compute(trade *Trade) {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
	key := trade.TradeSymbol
	// if _, ok := v.trades[key]; !ok {
	// 	v.trades[key] = make([]*Trade, 0) // TODO: to check. slice of pointer is instantiated, not as nil
	// }
	vol, err := strconv.ParseFloat(trade.Volume, 64)
	if err != nil {
		err = errors.New(err.Error())
		return
	}
	price, err := strconv.ParseFloat(trade.Price, 64)
	if err != nil {
		err = errors.New(err.Error())
		return
	}
	// TODO: to remove double calculation, replace using pointer
	// check window size
	if len(v.trades[key]) >= v.windowSize {
		firstTrade := v.trades[key][0]
		price, _ := strconv.ParseFloat(firstTrade.Price, 64) // assume parseable
		vol, _ := strconv.ParseFloat(firstTrade.Volume, 64)  // assume parseable
		v.trades[key] = v.trades[key][1:]
		v.sumVolume[key] -= vol
		v.priceTimesVolume[key] -= price * vol
		// v.vWAP[key] = v.priceTimesVolume[key] / v.sumVolume[key]
	}

	v.trades[key] = append(v.trades[key], trade)
	v.sumVolume[key] += vol
	v.priceTimesVolume[key] += price * vol
	v.vWAP[key] = v.priceTimesVolume[key] / v.sumVolume[key]
}
