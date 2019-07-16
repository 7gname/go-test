package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"strconv"
	"os"
)

const MAX_NUM_DB int = 16
const TAB_NUM_PER_DB int = 32
const DB_NAME_PRE string = "ofo_game"
const DB_TAB_NAME_PRE string = "g_user_props"
//const DB_DSN string = "ofo_game_w:ofo_game_w@tcp(10.6.42.138:7501)/%s"
const DB_DSN string = "ofo_game_w:UUlZAMP0KngZ3EzjbdTv@tcp(10.6.37.190:4458)/%s"

//各种翅膀碎片id
const (
	XiaoTianShiWingShatterType uint32 = 13
	DaTianShiWingShatterType   uint32 = 15
	ShengLingWingShatterType   uint32 = 18
	ShenShengWingShatterType   uint32 = 23
	BaiTianShiWingShatterType  uint32 = 24
	NormalWingShatterType      uint32 = 27 //暂时定义，尚未添加
)

//sql
const (
	SQL_GET_DISTINCT_USER_ID        string = "SELECT DISTINCT `user_id` FROM %s WHERE `status` = 1;"
	SQL_GET_WING_SHATTER_BY_USER_ID string = "SELECT `user_props_id`, `props_id`,`cnt` FROM %s WHERE `user_id` = %d AND `typ` = 1 AND `props_id` != %d AND `status` = 1;"
	SQL_UPDATE_WING_SHATTER         string = "UPDATE %s SET `status` = 99 WHERE `user_id` = %d AND `user_props_id` IN (%s);"
	SQL_INSERT_NORMAL_WING_SHATTER  string = "INSERT INTO %s (`user_id`, `typ`, `props_id`, `cnt`, `status`) VALUES (%d, 1, %d, %d, 1);"
)

func main() {
	//for i := 0; i < MAX_NUM_DB; i++ {
	//	fixData(i)
	//}

	ch := make([]chan int, MAX_NUM_DB)
	for i := 0; i < MAX_NUM_DB; i++ {
		ch[i] = make(chan int)
		go func(ch chan int, i int) {
			fixData(i)
			ch <- i
		}(ch[i], i)
	}

	for _, c := range ch {
		n := <-c
		fmt.Println(n)
	}

	//fixData(0)
}

func fixData(dbIndex int) () {
	fn := "log/db_" + strconv.Itoa(dbIndex);
	logFile, _ := os.Create(fn);
	log := log.New(logFile, "", 0)

	var dbName string = DB_NAME_PRE + "_" + strconv.Itoa(dbIndex)
	db, err := sql.Open("mysql", fmt.Sprintf(DB_DSN, dbName))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connect db: ", dbName, " success")
	}
	defer db.Close()

	for n := dbIndex * TAB_NUM_PER_DB; n < (dbIndex+1)*TAB_NUM_PER_DB; n++ {
		var tbName string = DB_TAB_NAME_PRE + "_" + strconv.Itoa(n)
		row, err := db.Query(fmt.Sprintf(SQL_GET_DISTINCT_USER_ID, tbName))
		if err != nil {
			fmt.Println(err)
		}

		var userIdArr []uint64
		for row.Next() {
			var userId uint64
			err := row.Scan(&userId)
			if err != nil {
				fmt.Println(err)
			}
			userIdArr = append(userIdArr, userId)
		}
		row.Close()
		//fmt.Println(tbName, ":", len(userIdArr))
		//continue
		//var t = 0

		for _, userId := range userIdArr {
			//fmt.Println(fmt.Sprintf("select `user_props_id`, `cnt` from %s where `user_id` = %d `props_id` = 27 and `status` = 1", tbName, userId))
			//continue
			r, err := db.Query(fmt.Sprintf("select `user_props_id`, `cnt` from %s where `user_id` = %d and typ = 1 and `status` = 1 order by create_time", tbName, userId))
			if err != nil {
				fmt.Println(err)
			}

			var cntMap = make(map[uint32]uint32)
			for r.Next() {
				var userPropsId uint32
				var cnt uint32
				err := r.Scan(&userPropsId, &cnt)
				if err != nil {
					fmt.Println(err)
				}
				cntMap[userPropsId] = cnt
			}

			var i = 1
			var l = len(cntMap)
			if l > 1 {
				//t += l
				for upid, _ := range cntMap {
					if i < l {
						log.Println(fmt.Sprintf("delete from %s where `user_id` = %d and `user_props_id` = %d", tbName, userId, upid))
					}
					i ++
				}
			}
		}
		//fmt.Println(t)
	}

	fmt.Println(dbName, " ok!")
}
