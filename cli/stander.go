package cli

import (
	"context"
	"fmt"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/urfave/cli/v2"
)


var stateListMessagesCmd = &cli.Command{
	Name:  "list-messages",
	Usage: "list messages on chain matching given criteria",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "to",
			Usage: "return messages to a given address",
		},
		&cli.StringFlag{
			Name:  "from",
			Usage: "return messages from a given address",
		},
		&cli.Uint64Flag{
			Name:  "toheight",
			Usage: "don't look before given block height",
		},
		&cli.BoolFlag{
			Name:  "cids",
			Usage: "print message CIDs instead of messages",
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		//ts, err := LoadTipSet(ctx, cctx, api)
		//if err != nil {
		//	return err
		//}

		h := abi.ChainEpoch(cctx.Uint64("from"))
		h2 := abi.ChainEpoch(cctx.Uint64("to"))
		//if ts == nil {
		//head, err := api.ChainHead(ctx)
		//stander

		//var stout *lapi.ComputeStateOutput


		fmt.Println(h,h2)
		for i:=h;i<=h2;i++ {
			ts,err:=api.ChainGetTipSetByHeight(ctx,i,types.TipSetKey{})
			fmt.Printf("Loading epoch... %d \n",i)

			ctx1:=context.WithValue(ctx,"replay",true)
			fmt.Println("List 1")
			o, _ := api.StateListMessages(ctx1, nil,ts.Key(),i)
			fmt.Println("List 2")
			//fmt.Println("Load tipset key: "+ts.Key().String())
			if err != nil {return err}
			//stout = o
			fmt.Println(o)
		}

		return nil
	},
}
