package retriever

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetOrderStatusById(ids []string, msg chan string, done chan bool) {

	for _, id := range ids {
		error_data := "-,-,-,-,-,"

		if id == "" {
			msg <- fmt.Sprintf("%s%s", error_data, "bad id")
			continue
		}

		res, err := http.Get(APIURL + id)

		if err != nil {
			msg <- fmt.Sprintf("%s%s", error_data, err)
			continue
		} else if res.StatusCode != 200 {
			msg <- fmt.Sprintf("%s%s", error_data, fmt.Sprintf("status code: %d", res.StatusCode))
			continue
		}
		defer res.Body.Close()

		order := &OrderStatusResponse{}
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(order)
		if err != nil {
			msg <- fmt.Sprintf("%s%s", error_data, err)
			continue
		}

		value := fmt.Sprintf("%s,%s,%s,%s,%s,-", order.ID, order.Status, order.Amount, order.CreatedAt, order.UpdatedAt)

		msg <- value
	}

	done <- true
}

func GetOrders(orders []string, nRoutines int) (res []string) {
	nRequestsPerRoutine := len(orders) / nRoutines
	offset := len(orders) % nRoutines
	responses := make([]string, 0, len(orders)+1)
	responses = append(responses, "id,status,amount,created,updated,error")

	message := make(chan string)
	done := make(chan bool, nRoutines)

	for i := 0; i < nRoutines; i++ {
		width := nRequestsPerRoutine

		if i == (nRoutines-1) && offset != 0 {
			width = offset + nRequestsPerRoutine
		}
		currentIds := orders[i:(i + width)]

		go GetOrderStatusById(currentIds, message, done)
	}

	i := 0

	for i < nRoutines {
		select {
		case msg := <-message:
			responses = append(responses, msg)
		case <-done:
			i++
		}
	}
	return responses
}
