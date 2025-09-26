package main

import ()

func main() {
	log.Println("starting payments microservice")

	defer log.Println("payments microservice closed")

	ctx := cmd.Context()

	paymentsInterface := createPaymentsMicroservice()

	if err := paymentsInterface.Run(ctx); err != nil {
		panic(err)
	}

}


func createOrdersMicroservice() amqp.paymentsInterface {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	paymentsSercive := payments_app.NewPaymentsService(
		payments_infra_orders.NewHTTPClient(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")),
	)

	paymentsInterface, err := amqp.NewPaymentsInterface(
		fmt.Sprintf("amqp://%s", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUQ"),
		paymentsSercive,
	)
	if err != nil {
		panic(err)
	}	

	return paymentsInterface

}






} 