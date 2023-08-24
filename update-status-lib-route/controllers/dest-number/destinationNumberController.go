package destnumber

import (
	"net/http"
	"strconv"
	maincontroller "update-status/controllers/mainController"
	"update-status/entities"
	destnumbermodel "update-status/models/destNumberModel"
)

func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		response := entities.Response{
			Success: false,
			Message: "Method not allowed",
		}

		maincontroller.Response(w, http.StatusMethodNotAllowed, response)
		return 
	}

	dnumber := "+" + r.URL.Query().Get("msisdn")
	statusNumber, _ := strconv.Atoi(r.URL.Query().Get("status"))
	
	var routing entities.DestNumber

	routing.DestinationNumber = dnumber
	routing.Status = int64(statusNumber)

	if ok := destnumbermodel.UpdateStatus(dnumber, routing); !ok {
		response := entities.Response{
			Success: false,
			Message: "Failed to Update Status Number",
		}

		maincontroller.Response(w, http.StatusNotModified, response)
		return 
	}

	response := entities.Response{
		Success: true,
		Message: "Successfully Update Status Number",
	}

	maincontroller.Response(w, http.StatusOK, response)

}