package payment

import (
	"fmt"
	"github.com/robiokidenis/microservice-mvc-2/conf"
	"github.com/robiokidenis/microservice-mvc-3/services"
	"runtime"
)

func UpdateExpired() {
	runtime.GOMAXPROCS(1)

	cfg := conf.NewConfig()

	db, err := conf.MysqlConnection(cfg)
	if err != nil {
		fmt.Printf("someting when wrong : %s", err.Error())
		return
	}

	paymentContract := services.NewPaymentServicesContract(db)

	data, err := paymentContract.FindExpiredPayment()
	if err != nil {
		fmt.Printf("someting when wrong : %s", err.Error())
		return
	}

	tx := db.Begin()
	defer tx.Rollback()

	for _, payment := range data {
		payment.Status = "EXPIRED"

		if err := paymentContract.UpdatePaymentStatus(payment, tx); err != nil {
			fmt.Printf("someting when wrong : %s", err.Error())
			return
		}
	}

	tx.Commit()

	return
}
