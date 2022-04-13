package orgAcct

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Account represents an account in the form3 org section.
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

type Client struct {
	BaseURL        string
	TimeoutSeconds int64
}

func (c *Client) createHttpClient() *http.Client {
	return &http.Client{Timeout: time.Second * time.Duration(c.TimeoutSeconds)}
}

// Create method sends a POST method to the REST API server
// for the creation of a new account
// It returns the AccountData object, and error returned by
// the API
func (c *Client) Create(accountData AccountData) (*AccountData, error) {
	//configure the http client
	httpClient := c.createHttpClient()

	// construct the request body
	reqBody, err := json.Marshal(accountData)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.BaseURL, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Perform the POST request to the requestUrl
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()

	// To check the response's Status Code
	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("Error - Status Code %d", resp.StatusCode))
	}

	// Reading the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	respAccountData := AccountData{}
	err = json.Unmarshal(body, &respAccountData)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &respAccountData, nil
}

// Fetch method sends a GET method to the REST API server
// to retrieve a specific account
// It returns the AccountData object, and error returned by
// the API
func (c *Client) Fetch(accountID string) (*AccountData, error) {
	//configure the http client
	httpClient := c.createHttpClient()

	// baseUrl/{account_id}
	requestUrl := fmt.Sprintf("%s/%s", c.BaseURL, accountID)
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Perform the GET request to the requestUrl
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()

	// To check the response's Status Code
	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("Account Not Found")
	} else if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Error - Status Code %d", resp.StatusCode))
	}

	// Reading the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	accountData := AccountData{}
	err = json.Unmarshal(body, &accountData)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &accountData, nil
}

// Delete method sends a DELETE method to the REST API server
// to delete a specific account
// It returns the AccountData object, and error returned by
// the API
func (c *Client) Delete(accountID string, version int) error {
	//configure the http client
	httpClient := c.createHttpClient()

	// baseUrl/{account_id}
	requestUrl := fmt.Sprintf("%s/%s", c.BaseURL, accountID)

	// Prepare the request
	req, err := http.NewRequest(http.MethodDelete, requestUrl, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Perform the DELETE request to the requestUrl
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("Account Not Found")
	} else if resp.StatusCode == http.StatusConflict {
		return errors.New("Incorrect Specified Version Number")
	} else if resp.StatusCode != http.StatusNoContent {
		return errors.New(fmt.Sprintf("Error - Status Code %d", resp.StatusCode))
	}

	return nil
}
