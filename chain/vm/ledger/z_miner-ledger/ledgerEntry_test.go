package z_miner_ledger

//func TestLedgerEntry(t *testing.T){
//
//	str_cid:="bafy2bzacecqybhyybzna67odowqsrpcol3rg74yxmobspfycxayda7f3iqt5a"
//	cid,_:=cid.Parse(str_cid)
//	bi,_:=big.FromString("1083849489732541506155204973")
////	c.removeEntry(cid.String())
//
//	addr,_:=Address.NewFromString("f099608")
//	//fmt.Println(addr)
//	e:= LedgerEntry{
//		Id:          str_cid,
//		EntryCid:    cid,
//		Address:     addr,
//		Offset:      Address.Address{},
//		Method:   1,
//		Amount:       abi.NewTokenAmount(145),
//		TotalAmount: DimBalance{Available: bi, InitialPledge:abi.NewTokenAmount(123)},
//		Amount:      DimBalance{Available: bi},
//		Balance:     DimBalance{},
//		GasFee:      big.MustFromString("1083849489732541506155204973"),
//		GasUsed:     104,
//		CallDepth:   105,
//		Sector:      abi.SectorID{
//			Miner:  99806,
//			Number: 10600000,
//		},
//		Deals: nil,
//	}
//
//	InsertEntry(e)
//	r:= FindEntry(cid.String())
//
//	if r.Id!=e.Id { t.Fatalf("Cid incorrect %s",r.Id)}
//	if r.EntryCid !=e.EntryCid { t.Fatalf("EntryCid incorrect %s",r.EntryCid)}
//
//	if r.Address!=e.Address { t.Fatalf("Address incorrect %s",r.Address)}
//	if r.Offset!=e.Offset { t.Fatalf("Offset Address incorrect %s",r.Offset)}
//	if r.Offset!=Address.Undef { t.Fatalf("Offset Address incorrect %s, should be %s",r.Offset,Address.Undef.String())}
//
//	if r.Method!=e.Method { t.Fatalf("Method incorrect %s",r.Method)}
//
//	if r.Amount.Int.String()!=e.Amount.Int.String() { t.Fatalf("Amount incorrect %s -> %s ",
//		e.Amount.Int.String(),
//		r.Amount.Int.String())}
//
//	if !reflect.DeepEqual(r.Amount,e.Amount) { t.Fatalf("Amount incorrect %s -> %s ",
//		e.Amount.Int.String(),
//		r.Amount.Int.String())}
//
//	if !reflect.DeepEqual(r.TotalAmount,e.TotalAmount) { t.Fatalf("TotalAmount incorrect %v",r.TotalAmount)}
//	if !reflect.DeepEqual(r.Amount,e.Amount) { t.Fatalf("Amount incorrect %v",r.Amount)}
//
//
//	if !reflect.DeepEqual(r.Balance,e.Balance) { t.Fatalf("Available incorrect %v -> %v",e.Balance,r.Balance)}
//
//	if r.GasUsed!=e.GasUsed { t.Fatalf("GasUsed incorrect %v",r.GasUsed)}
//	if r.GasUsed!=104 { t.Fatalf("GasUsed incorrect %v",r.GasUsed)}
//
//	if r.CallDepth!=e.CallDepth { t.Fatalf("CallDepth incorrect %v %v",e.CallDepth,r.CallDepth)}
//	if r.CallDepth!=105 { t.Fatalf("CallDepth incorrect %v ",r.CallDepth)}
//
//	if r.Sector!=e.Sector { t.Fatalf("Sector incorrect %v %v",e.Sector,r.Sector)}
//
//	//if r.Deals !=e.Deals { t.Fatalf("Deals incorrect %v",r.Deals)}
//	//if r.Deals !=107 { t.Fatalf("Deals incorrect %v",r.Deals)}
//	//m1:=make(map[string]abi.TokenAmount)
//	//m2:=make(map[string]abi.TokenAmount)
//	//if !reflect.DeepEqual(m1,m2) { t.Fatalf("maps not equal %v -> %v",m1,m2)}
//
//}
