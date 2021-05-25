package ledg_util

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/vm/ledger/models"
	logging "github.com/ipfs/go-log/v2"
	"strconv"
)

var llog			= logging.Logger("gledger")
var AddressCache map[string]int32

func loadAddressCache(){
	AddressCache=make(map[string]int32)
	_db:=GetPgDatabase()

	var allAccounts []models.Account
	_db.Find(&allAccounts)

	for _,a:=range allAccounts{
		AddressCache[a.Address]=a.ID
	}
}

func getNewAddressIdFromDb(addr address.Address,epoch abi.ChainEpoch)(int32,error){
	if id,err:=address.IDFromAddress(addr);err==nil {return int32(id),nil}

	var minId int32
	//err:=db.Clauses(clause.Locking{
	//	Strength: "UPDATE",
	//	Table: clause.Table{ Name:"accounts"}}).
	//	Select("min(id)").Row().Scan(&minId)
	err:=db.Table("accounts").Select("min(id)").Row().Scan(&minId)
	return minId-1,err
}
func findInDb(addr address.Address,epoch abi.ChainEpoch) (int32, bool) {
	var accInDb models.Account
	result := db.Where("Address=?", addr.String()).First(&accInDb)
	if result.Error!=nil {
		llog.Info("findInDB: "+result.Error.Error())
		return 0, false
	}

		//if epoch<accInDb.CreationEpoch {
		//	db.Model(models.Account{}).Where("Address=?",addr.String()).Update("creation_epoch",epoch)
		//}


		//llog.Info("FoundInDb id: "+accInDb.Address)
		return accInDb.ID, true


}

func GetOrCreateAccountFromId( addrId int32,name string,epoch abi.ChainEpoch) (int32,error) {

	addr,_:=address.NewIDAddress(uint64(addrId))
	return GetOrCreateAccountFromAddress(addr,name,epoch)
}

func GetOrCreateAccountFromAddress( addr address.Address,name string,epoch abi.ChainEpoch) (int32,error){


	accId,foundInCache:=AddressCache[addr.String()]
	if foundInCache {
		return 	accId,nil //ret from acc cache
	}else {
		llog.Info("NotFound in cache id: "+addr.String()+" "+strconv.Itoa( int(accId)))
	}

	llog.Infof("Searching %s",addr.String())
	accId,foundInDb:=findInDb(addr,epoch)
	if foundInDb {
		AddressCache[addr.String()]=accId
		return accId,nil
	}//ret from db

	//creating new Address in db and cache
		//db.Begin()
		newId,err := getNewAddressIdFromDb(addr,epoch)
		llog.Infof("New id for %s : %d",addr.String(),newId)
		if err!=nil {
			llog.Infof("getting newID:"+err.Error())
			return 0,err
		}
		a := models.Account{
			ID:            newId,
			Address:       addr.String(),
			Name:          name,
			CompanyId:     "",
			Signature:     "",
			CreationEpoch: epoch,
			Protocol:      addr.Protocol(),
		}
		db.Create(&a)
		//db.Commit()

		return 0,err

}



func GetOrCreateAccount_del(minerId int32,name string){

	db:=GetPgDatabase()

	//if result.RowsAffected==0{
	minerAddress,_:=address.NewIDAddress(uint64(minerId))
	if name==""{name="generated: "+minerAddress.String()}
	m:=models.Account{
		ID: minerId,
		Address: minerAddress.String(),
		Name: name,
	}
	db.FirstOrCreate(&m)
	//}
}