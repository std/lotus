package ledg_util

import (
	"github.com/filecoin-project/lotus/chain/vm/ledger/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=fcw password=fcw dbname=ledger port=5432 sslmode=disable TimeZone=Europe/Riga"

var db *gorm.DB

//func Exists(v interface{},id int) bool {
//	//var found models.Account
//	result:=db.First(v,id)
//	if result.RowsAffected>0{ return true }
//	return false
//}
func create_epochLedgerEntriesPartition(){
	sql:=`CREATE TABLE ledger_entries_50000 PARTITION OF epoch_ledger_entries for values from (0) to (50000)`
	_db:=GetPgDatabase()
	_db.Exec(sql)
}
func create_epochLedgerEntriesTable(){
	sql:=`create table epoch_ledger_entries
(
    epoch integer,
    id            bigint not null,

    address_id    integer
        constraint fk_ledger_entries_address
            references accounts,
    offset_id     integer
        constraint fk_ledger_entries_offset
            references accounts,
    amount        text,
    method        bigint,
    call_depth    bigint,
    sector_id     integer,
    miner_id      integer
        constraint fk_ledger_entries_miner
            references accounts,
    tx_id         bigint
        constraint fk_ledger_entries_tx
            references tx_messages,
    entry_type    text,
    method_name   text,
    implicit      boolean,
    constraint fk_ledger_entries_sector
        foreign key (miner_id, sector_id) references sectors,
         constraint part_ledger_entries_pkey
            primary key (epoch,id)
) partition by range (epoch)
`
	_db:=GetPgDatabase()
	_db.Migrator().DropTable("epoch_ledger_entries")

	_db.Exec(sql)
}
func fillDefaultSector(){

	_db:=GetPgDatabase()

	sec:=models.Sector{
		MinerId:          0,
		ID:               0,

	}
	//if Exists(&sec,0){		return}
	_db.Create(&sec)
	//fmt.Println("fillDefaultSector")
	//_db.FirstOrCreate(&sec,&sec)
}
func fillMinerTypes(){
	_db:=GetPgDatabase()

	types:= map[int32]string{
		0:"System",
		1:"Init", //f01
		2:"Reward",                        //f02
		3:"Cron",                          //f03
		4:"Power",                         //f04
		5:"Market",                        //f05
		6:"Registry",                      //f06
		99:"Burn",//f099
		1000:"Miner",
	}
	for id,descr:=range types {
		//if !Exists(&models.AccountType{},id) {
			_db.FirstOrCreate(&models.AccountType{
				Id:    id,
				Descr: descr,
			})
			

				//actorAddres,_:=address.NewIDAddress(uint64(id))
				//_db.FirstOrCreate(&models.Account{
				//	ID:            int32(id),
				//	Address: actorAddres.String(),
				//	Name:          descr,
				//
				//	Protocol:     models.ActorTypeConst(id),
				//})
			}
		//}

	for id,descr:=range types {
		if id != 1000 {
			GetOrCreateAccountFromId(id, descr,0)
		}
	}

}

func migrate(){
	db.AutoMigrate(

		models.Block{},
		models.Tipset{},
		models.TxMessage{},
		//models.Account{},
		//models.AccountType{},
		models.Sector{},
		models.PowerEntry{},
		//models.LedgerEntry{},
		models.TablePart{},

		)
	loadAddressCache()
	GetOrCreateAccountFromId(0,"System#",0)

	fillMinerTypes()
	fillDefaultTablePart()

	//create_epochLedgerEntriesTable()
	truncate_ledgerEntries()

	//fillDefaultSector()
}

func truncate_ledgerEntries() {
	sql:=`truncate table ledger_entries`
	GetPgDatabase().Exec(sql)
}

func fillDefaultTablePart() {
	_db:=GetPgDatabase()

	sec:=models.TablePart{
		Id:0,
		TablePartName: "default",
	}
	//if Exists(&sec,0){		return}
	_db.Create(&sec)
}
func GetPgDatabase() *gorm.DB{
	if db==nil {
		newdb, err:= gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err!=nil {panic(err)}
		db=newdb
		dropTables()
		migrate()
	}
	return db
}

func dropTables() {
	_db:=GetPgDatabase()
	_db.Migrator().DropTable(models.TablePart{})
	//_db.Migrator().DropTable(models.LedgerEntry{})
	_db.Migrator().DropTable(models.PowerEntry{})
	_db.Migrator().DropTable(models.TxMessage{})
	_db.Migrator().DropTable(models.Sector{})
	_db.Migrator().DropTable(models.Block{})
	_db.Migrator().DropTable(models.Tipset{})
	//_db.Migrator().DropTable(models.Account{})
	//_db.Migrator().DropTable(models.AccountType{})


}
