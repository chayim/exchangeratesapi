package exchangeratesapi_test

import (
	"exchangeratesapi"
	"testing"
	"time"
)

func TestCovertForCurrency(t *testing.T) {
	currencyData, err := exchangeratesapi.ConvertForCurrency("CAD")
	if err != nil {
		t.Error(err)
	}
	if currencyData <= 0.0 {
		t.Error("Received invalid currency data.")
	}
}

func TestCovertForCurrencyWithBase(t *testing.T) {
	currencyData, err := exchangeratesapi.ConvertForCurrencyWithBase("CAD", "CAD")
	if err != nil {
		t.Error(err)
	}
	if currencyData != 1.0 {
		t.Error("Received invalid currency data.")
	}
}

func TestGetLatest(t *testing.T) {
	currencyData, err := exchangeratesapi.GetLatest()
	if err != nil {
		t.Error(err)
	}
	for key, _ := range currencyData {
		if len(key) != 3 {
			t.Errorf("%s is an invalid key.", key)
		}
		if currencyData[key] < 0 {
			t.Errorf("Invalid currency amount %f", currencyData[key])
		}
	}

	currencyData, err = exchangeratesapi.GetLatest("CAD", "USD", "ILS")
	if err != nil {
		t.Error(err)
	}

	if currencyData["CAD"] < 0.1 {
		t.Errorf("Invalid CAD currency data %f.", currencyData["CAD"])
	}

	if currencyData["USD"] < 0.1 {
		t.Errorf("Invalid USD currency data %f.", currencyData["USD"])
	}

	if currencyData["ILS"] < 0.1 {
		t.Errorf("Invalid ILS currency data %f.", currencyData["ILS"])
	}

}

func TestGetLatestForCurrency(t *testing.T) {
	currencyData, err := exchangeratesapi.GetLatestForCurrency("CAD")
	if err != nil {
		t.Error(err)
	}
	for key, _ := range currencyData {
		if len(key) != 3 {
			t.Errorf("%s is an invalid key.", key)
		}
		if currencyData[key] < 0 {
			t.Errorf("Invalid currency amount %f", currencyData[key])
		}
	}

	currencyData, err = exchangeratesapi.GetLatestForCurrency("CAD", "USD", "ILS")
	if err != nil {
		t.Error(err)
	}

	if currencyData["CAD"] != 0.0 {
		t.Errorf("Invalid CAD currency data %f.", currencyData["CAD"])
	}

	if currencyData["USD"] < 0.5 {
		t.Errorf("Invalid USD currency data %f.", currencyData["USD"])
	}

	if currencyData["ILS"] < 0.1 {
		t.Errorf("Invalid ILS currency data %f.", currencyData["ILS"])
	}

}

func TestGetForDate(t *testing.T) {

	date, _ := time.Parse("2006-01-02", "2003-12-31")
	currencyData, err := exchangeratesapi.GetForDate(date)
	if err != nil {
		t.Error(err)
	}
	for key, _ := range currencyData {
		if len(key) != 3 {
			t.Errorf("%s is an invalid key.", key)
		}
		if currencyData[key] < 0 {
			t.Errorf("Invalid currency amount %f", currencyData[key])
		}
	}

	currencyData, err = exchangeratesapi.GetForDate(date, "CAD", "USD")
	if err != nil {
		t.Error(err)
	}

	if currencyData["CAD"] != 1.6234 {
		t.Errorf("CAD should have been 1.6234, not %f", currencyData["CAD"])
	}

	if currencyData["USD"] < 1.263 {
		t.Errorf("USD should have been 1.263, not %f", currencyData["USD"])
	}

}

func TestGetForDateForCurrency(t *testing.T) {

	date, _ := time.Parse("2006-01-02", "2003-12-31")
	currencyData, err := exchangeratesapi.GetForDateForCurrency(date, "CAD")
	if err != nil {
		t.Error(err)
	}
	for key, _ := range currencyData {
		if len(key) != 3 {
			t.Errorf("%s is an invalid key.", key)
		}
		if currencyData[key] < 0 {
			t.Errorf("Invalid currency amount %f", currencyData[key])
		}
	}

	currencyData, err = exchangeratesapi.GetForDateForCurrency(date, "CAD", "GBP", "USD")
	if err != nil {
		t.Error(err)
	}

	if currencyData["GBP"] != 0.7048 {
		t.Errorf("GBP should have been 0.7048, not %f", currencyData["GBP"])
	}

	if currencyData["USD"] < 1.263 {
		t.Errorf("USD should have been 1.263, not %f", currencyData["USD"])
	}

}

func TestGetBetweenDates(t *testing.T) {

	date, _ := time.Parse("2006-01-02", "2003-12-31")
	endDate, _ := time.Parse("2006-01-02", "2004-01-11")
	currencyData, err := exchangeratesapi.GetBetweenDates(date, endDate)
	if err != nil {
		t.Error(err)
	}
	if len(currencyData) <= 1 {
		t.Error("ExchangesRatesAPI failed to return more than one dated item")
	}

	currencyData, err = exchangeratesapi.GetBetweenDates(date, endDate, "CAD", "GBP", "USD")
	if err != nil {
		t.Error(err)
	}

	for r, _ := range currencyData {
		if currencyData[r]["GBP"] <= 0.0 {
			t.Errorf("Invalid GBP currency data.")
		}

		if currencyData[r]["USD"] <= 0.0 {
			t.Errorf("Invalid USD currency data.")
		}
	}

}

func TestGetBetweenDatesForCurrency(t *testing.T) {

	date, _ := time.Parse("2006-01-02", "2003-12-31")
	endDate, _ := time.Parse("2006-01-02", "2004-01-11")
	currencyData, err := exchangeratesapi.GetBetweenDatesForCurrency(date, endDate, "CAD")
	if err != nil {
		t.Error(err)
	}
	if len(currencyData) <= 1 {
		t.Error("ExchangesRatesAPI failed to return more than one dated item")
	}

	currencyData, err = exchangeratesapi.GetBetweenDatesForCurrency(date, endDate, "CAD", "CAD", "GBP", "USD")
	if err != nil {
		t.Error(err)
	}

	for r, _ := range currencyData {
		if currencyData[r]["CAD"] != 1.0 {
			t.Errorf("CAD should be 1.0 instead of %f", currencyData[r]["CAD"])
		}

		if currencyData[r]["GBP"] <= 0.0 {
			t.Errorf("Invalid GBP currency data.")
		}

		if currencyData[r]["USD"] <= 0.0 {
			t.Errorf("Invalid USD currency data.")
		}
	}

}
