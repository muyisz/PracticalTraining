package database

import (
	"log"
	"sync"
	"time"

	"github.com/muyisz/PracticalTraining/constant"
	"github.com/muyisz/PracticalTraining/database/model"
	"github.com/muyisz/PracticalTraining/util"
)

var (
	lock             sync.Mutex
	orderIncrementID int
)

func GetUserOrder(oid int) ([]model.Order, error) {
	orderList := []model.Order{}

	if err := db.Table(constant.TableOrder).Where("id=?", oid).Find(&orderList).Error; !util.DBQueryErr(err) {
		log.Printf("[GetProductByID] First err,oid:%+v,err:%+v", oid, err)
		return nil, err
	}

	return orderList, nil
}

func CreateOrder(mid int, address string) (int, error) {
	order := model.Order{
		ID:      OrderIncrementID(),
		Time:    time.Now(),
		State:   0,
		Mid:     mid,
		Address: address,
	}

	if err := db.Table(constant.TableOrder).Create(&order).Error; err != nil {
		log.Printf("[CreateOrder] Create err,order:%+v,err:%+v", order, err)
		return 0, err
	}

	return order.ID, nil
}

func GetMaxIncrementID() (int, error) {
	var maxID int
	if err := db.Table(constant.TableOrder).Select("max(id)").Row().Scan(&maxID); err != nil {
		log.Printf("[GetMaxIncrementID] Scan err,err:%+v", err)
		return 0, err
	}

	return maxID, nil
}

func OrderIncrementID() int {
	lock.Lock()
	defer lock.Unlock()
	orderIncrementID++
	return orderIncrementID
}

func SetOrderIncrementID() {
	lock.Lock()
	defer lock.Unlock()
	maxID, err := GetMaxIncrementID()
	if err != nil {
		panic(err)
	}
	orderIncrementID = maxID
}
