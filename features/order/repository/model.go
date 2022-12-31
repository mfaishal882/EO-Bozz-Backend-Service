package repository

import (
	"capstone-alta1/features/order"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	EventName         string
	StartDate         time.Time
	StartTime         time.Duration
	EndDate           time.Time
	EndTime           time.Duration
	EventLocation     string
	EventAddress      string
	NotesForPartner   string
	ServiceName       string
	ServicePrice      uint
	GrossAmmount      uint
	PaymentMethod     string
	OrderStatus       string
	PayoutRecieptFile string
	PayoutDate        time.Time `gorm:"default:null"`
	ServiceID         uint
	ClientID          uint
	DetailOrder       []DetailOrder
}

type DetailOrder struct {
	gorm.Model
	AdditionalName      string
	AdditionalPrice     uint
	Qty                 uint
	DetailOrderTotal    uint
	ServiceAdditionalID uint
	OrderID             uint
}

type ServiceAdditional struct {
	gorm.Model
	AdditionalID uint
	ServiceID    uint
	DetailOrders []DetailOrder
}

type Additional struct {
	gorm.Model
	AdditionalName    string
	AdditionalPrice   uint
	PartnerID         uint
	ServiceAdditional []ServiceAdditional
}

type Client struct {
	gorm.Model
	Gender          string
	Address         string
	City            string
	Phone           string
	ClientImageFile string
	UserID          uint
	User            User
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
}

type Service struct {
	gorm.Model
	ServiceName        string
	ServiceDescription string
	ServiceCategory    string
	ServicePrice       uint
	AverageRating      float64
	ServiceImageFile   string
	City               string
	PartnerID          uint
	ServiceAdditional  []ServiceAdditional
}

// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore order.Core) Order {
	modelData := Order{
		EventName:       dataCore.EventName,
		StartDate:       dataCore.StartDate,
		EndDate:         dataCore.EndDate,
		EventLocation:   dataCore.EventLocation,
		EventAddress:    dataCore.EventAddress,
		NotesForPartner: dataCore.NotesForPartner,
		PaymentMethod:   dataCore.PaymentMethod,
		ServiceID:       dataCore.ServiceID,
		ClientID:        dataCore.ClientID,
	}
	return modelData
}

func fromDetailOrder(dataCore order.DetailOrder) DetailOrder {
	modelData := DetailOrder{
		ServiceAdditionalID: dataCore.ServiceAdditionalID,
		Qty:                 dataCore.Qty,
	}
	return modelData
}

func fromDetailOrderList(dataCore []order.DetailOrder) []DetailOrder {
	var dataModel []DetailOrder
	for _, v := range dataCore {
		dataModel = append(dataModel, fromDetailOrder(v))
	}
	return dataModel
}

func fromCoreStatus(dataCore order.Core) Order {
	modelData := Order{
		OrderStatus: dataCore.OrderStatus,
	}
	return modelData
}

func fromCorePayout(dataCore order.Core) Order {
	modelData := Order{
		PayoutRecieptFile: dataCore.PayoutRecieptFile,
		PayoutDate:        time.Now(),
	}
	return modelData
}

// mengubah struct model gorm ke struct core
func (dataModel *Order) toCore() order.Core {
	return order.Core{
		ID:            dataModel.ID,
		EventName:     dataModel.EventName,
		StartDate:     dataModel.StartDate,
		EndDate:       dataModel.EndDate,
		EventLocation: dataModel.EventLocation,
		ServiceName:   dataModel.ServiceName,
		GrossAmmount:  dataModel.GrossAmmount,
		OrderStatus:   dataModel.OrderStatus,
		ServiceID:     dataModel.ServiceID,
		ClientID:      dataModel.ClientID,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Order) []order.Core {
	var dataCore []order.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

func (dataModel *Order) toCoreOrder() order.Core {
	return order.Core{
		ID:                dataModel.ID,
		EventName:         dataModel.EventName,
		StartDate:         dataModel.StartDate,
		StartTime:         dataModel.StartTime,
		EndDate:           dataModel.EndDate,
		EndTime:           dataModel.EndTime,
		EventLocation:     dataModel.EventLocation,
		EventAddress:      dataModel.EventAddress,
		NotesForPartner:   dataModel.NotesForPartner,
		ServiceName:       dataModel.ServiceName,
		ServicePrice:      dataModel.ServicePrice,
		GrossAmmount:      dataModel.GrossAmmount,
		PaymentMethod:     dataModel.PaymentMethod,
		OrderStatus:       dataModel.OrderStatus,
		PayoutRecieptFile: dataModel.PayoutRecieptFile,
		PayoutDate:        dataModel.PayoutDate,
		ServiceID:         dataModel.ServiceID,
		ClientID:          dataModel.ClientID,
	}
}

func (dataModel *DetailOrder) toCoreDetailOrder() order.DetailOrder {
	return order.DetailOrder{
		ID:                  dataModel.ID,
		ServiceAdditionalID: dataModel.ServiceAdditionalID,
		AdditionalName:      dataModel.AdditionalName,
		AdditionalPrice:     dataModel.AdditionalPrice,
		Qty:                 dataModel.Qty,
		DetailOrderTotal:    dataModel.DetailOrderTotal,
		OrderID:             dataModel.OrderID,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreDetailOrderList(dataModel []DetailOrder) []order.DetailOrder {
	var dataCore []order.DetailOrder
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreDetailOrder())
	}
	return dataCore
}
