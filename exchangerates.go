package exchangeratesapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type data struct {
	Base  string
	Date  string
	Rates map[string]float32
}

type datedData struct {
	Base    string
	StartAt string
	EndAt   string
	Rates   map[string]map[string]float32
}

// Reformat the time object to meet the date format required by the exchangeratesapi.
func reformatDate(date time.Time) string {
	return date.Format("2006-01-02")
}

// This helper function fetches the latest exchange rate for the specified security.
func ConvertForCurrency(currency string) (float32, error) {
	res, err := GetLatest(currency)
	if err != nil {
		return -1.0, err
	}

	return res[currency], nil
}

// This helper function fetches the latest exchange rate for the specified security.
func ConvertForCurrencyWithBase(base string, currency string) (float32, error) {
	res, err := GetLatestForCurrency(base, currency)
	if err != nil {
		return -1.0, err
	}

	return res[currency], nil
}

// Fetch the latest forex data, optionally limiting to specific currencies.
func GetLatest(currencies ...string) (map[string]float32, error) {
	url := fmt.Sprintf("%s/latest", baseURL)
	all := strings.Join(currencies, ",")
	if all != "" {
		url = fmt.Sprintf("%s?symbols=%s", url, all)
	}
	return getCurrencyData(url)
}

// Fetch the latest forex data specifying another currency as a baseline, optionally limiting to specific currencies.
// ExchangeRatesAPI defaults to EUR as the baseline currency.
func GetLatestForCurrency(currency string, currencies ...string) (map[string]float32, error) {
	all := strings.Join(currencies, ",")
	url := fmt.Sprintf("%s/latest?base=%s", baseURL, currency)
	if all != "" {
		url = fmt.Sprintf("%s&symbols=%s", url, all)
	}
	return getCurrencyData(url)
}

// Fetch forex data for a specific date, optionally limiting to specific currencies.
func GetForDate(t time.Time, currencies ...string) (map[string]float32, error) {
	all := strings.Join(currencies, ",")
	date := reformatDate(t)
	url := fmt.Sprintf("%s/%s", baseURL, date)
	if all != "" {
		url = fmt.Sprintf("%s?&symbols=%s", url, all)
	}
	return getCurrencyData(url)
}

// Fetch forex data for a specific date, baselined to another currency, optionally limiting to specific currencies.
// ExchangeRatesAPI defaults to EUR as the baseline currency.
func GetForDateForCurrency(t time.Time, currency string, currencies ...string) (map[string]float32, error) {
	all := strings.Join(currencies, ",")
	date := reformatDate(t)
	url := fmt.Sprintf("%s/%s?currency=%s", baseURL, date, currency)
	if all != "" {
		url = fmt.Sprintf("%s&symbols=%s", url, all)
	}
	return getCurrencyData(url)
}

// Fetch forex data for between dates, optionally limiting to specific currencies.
func GetBetweenDates(startDate time.Time, endDate time.Time, currencies ...string) (map[string]map[string]float32, error) {
	all := strings.Join(currencies, ",")
	date1 := reformatDate(startDate)
	date2 := reformatDate(endDate)
	url := fmt.Sprintf("%s/history?&start_at=%s&end_at=%s", baseURL, date1, date2)
	if all != "" {
		url = fmt.Sprintf("%s&symbols=%s", url, all)
	}
	return getDatedCurrencyData(url)
}

// Fetch forex data for between dates, baselined to another currency, optionally limiting to specific currencies.
// ExchangeRatesAPI defaults to EUR as the baseline currency.
func GetBetweenDatesForCurrency(startDate time.Time, endDate time.Time, currency string, currencies ...string) (map[string]map[string]float32, error) {
	all := strings.Join(currencies, ",")
	date1 := reformatDate(startDate)
	date2 := reformatDate(endDate)
	url := fmt.Sprintf("%s/history?base=%s&start_at=%s&end_at=%s", baseURL, currency, date1, date2)
	if all != "" {
		url = fmt.Sprintf("%s&symbols=%s", url, all)
	}
	return getDatedCurrencyData(url)
}

func getCurrencyData(url string) (map[string]float32, error) {
	resp, _ := http.Get(url)
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("Failed to fetch currency.\n%s", body)
		return nil, err
	}

	d := data{}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return d.Rates, e
	}
	json.Unmarshal(body, &d)
	return d.Rates, nil
}

func getDatedCurrencyData(url string) (map[string]map[string]float32, error) {
	resp, _ := http.Get(url)
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("Failed to fetch currency.\n%s", body)
		return nil, err
	}

	d := datedData{}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return d.Rates, e
	}
	json.Unmarshal(body, &d)
	return d.Rates, nil
}
