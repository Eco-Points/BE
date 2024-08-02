package repository

import (
	"eco_points/internal/features/excelize"
	"eco_points/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
)

type excelMakeQuery struct {
	db    *gorm.DB
	excel utils.ExcelMakeUtilityInterface
	cloud utils.CloudinaryUtilityInterface
}

func NewExcelQuery(d *gorm.DB, e utils.ExcelMakeUtilityInterface, c utils.CloudinaryUtilityInterface) excelize.QueryMakeExcelInterface {
	return &excelMakeQuery{
		db:    d,
		excel: e,
		cloud: c,
	}

}

var DbSchema string

func (q *excelMakeQuery) SetDbSchema(schema string) {
	DbSchema = schema
}

func (q *excelMakeQuery) MakeExcelFromDeposit(table string, date string, userid uint, isVerif bool, limit uint) (string, string, error) {

	var err error
	var jsonData []byte
	if strings.Contains(table, "depo") || strings.Contains(table, "Depo") {
		var deposit []WasteDeposit

		if userid == 0 {
			query := fmt.Sprintf(`select wd.id, t.trash_type, wd.point, wd.status, wd.quantity, u.fullname, wd.created_at from "%s".waste_deposits wd 
								join "%s".trashes t on t.id = wd.trash_id 
								join "%s".users u on u.id = wd.user_id 
								where wd."deleted_at" IS NULL `, DbSchema, DbSchema, DbSchema)
			if !isVerif {
				query = query + `and wd."status" <> 'verified' `
			} else {
				query = query + `and wd."status" <> 'rejected' and wd."status" <> 'pending' `
			}
			if limit != 0 {
				query = query + fmt.Sprintf("limit %d", limit)
			}
			err := q.db.Debug().Raw(query).Scan(&deposit).Error
			if err != nil {
				log.Println("error get data", err)
				return "", "", err
			}
			jsonData, err = json.Marshal(toWasteDepositInterface(deposit))
			if err != nil {
				log.Println("error marshal ", err)
				return "", "", nil
			}

		} else {
			query := fmt.Sprintf(`select wd.id, t.trash_type, wd.point, wd.status, wd.quantity, u.fullname, wd.created_at from "%s".waste_deposits wd 
								join "%s".trashes t on t.id = wd.trash_id 
								join "%s".users u on u.id = wd.user_id 
								where wd.user_id = %d and wd."deleted_at" IS NULL `, DbSchema, DbSchema, DbSchema, userid)

			if !isVerif {
				query = query + ` and wd."status" <> 'verified' `
			} else {
				query = query + ` and wd."status" <> 'rejected' and wd."status" <> 'pending' `
			}
			if limit != 0 {
				query = query + fmt.Sprintf("limit %d", limit)
			}
			err := q.db.Debug().Raw(query).Scan(&deposit).Error
			if err != nil {
				log.Println("error get data", err)
				return "", "", err
			}
			jsonData, err = json.Marshal(toWasteDepositInterface(deposit))
			if err != nil {
				log.Println("error marshal ", err)
				return "", "", nil
			}
		}

	} else if strings.Contains(table, "rew") || strings.Contains(table, "Rew") {
		var exchange []Exchange
		if userid == 0 {
			query := fmt.Sprintf(`select e.point_used, r."name" as rewards, u.fullname as name, e.created_at from "%s".exchanges e 
								join "%s".rewards r on r.id = e.reward_id	
								join "%s".users u on u.id = e.user_id
								where e."deleted_at" IS NULL `, DbSchema, DbSchema, DbSchema)
			if limit != 0 {
				query = query + fmt.Sprintf("limit %d", limit)
			}
			err := q.db.Debug().Raw(query).Scan(&exchange).Error
			if err != nil {
				log.Println("error get data", err)
				return "", "", err
			}
			jsonData, err = json.Marshal(toExchangeInterface(exchange))
			if err != nil {
				log.Println("error marshal ", err)
				return "", "", nil
			}
		} else {
			query := fmt.Sprintf(`select e.point_used, r."name" as rewards, u.fullname as name, e.created_at from "%s".exchanges e 
								join "%s".rewards r on r.id = e.reward_id	
								join "%s".users u on u.id = e.user_id
								where user_id = %d and e."deleted_at" IS NULL `, DbSchema, DbSchema, DbSchema, userid)
			if limit != 0 {
				query = query + fmt.Sprintf("limit %d;", limit)
			}
			err := q.db.Debug().Raw(query).Scan(&exchange).Error
			if err != nil {
				log.Println("error get data", err)
				return "", "", err
			}
			jsonData, err = json.Marshal(toExchangeInterface(exchange))
			if err != nil {
				log.Println("error marshal ", err)
				return "", "", nil
			}
		}
	} else {
		return "", "url", errors.New("query param invalid")

	}

	file, err := q.excel.MakeExcel(table, date, string(jsonData))
	if err != nil {
		log.Println("error make excel ", err)
		return "", "", err
	}
	src, err := os.Open(fmt.Sprintf("./%s", file))
	if err != nil {
		log.Println("error when open image", err)
		return "", "", err
	}
	defer src.Close()

	url, err := q.cloud.UploadToCloudinary(src, file)
	if err != nil {
		log.Println("error when upload image", err)
		return "", "", err
	}

	return file, url, nil
}
