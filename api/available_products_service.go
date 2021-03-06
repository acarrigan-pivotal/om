package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const availableProductsEndpoint = "/api/v0/available_products"

type UploadAvailableProductInput struct {
	ContentLength   int64
	Product         io.Reader
	ContentType     string
	PollingInterval int
}

type ProductInfo struct {
	Name    string `json:"name"`
	Version string `json:"product_version"`
}

type UploadAvailableProductOutput struct{}

type AvailableProductsOutput struct {
	ProductsList []ProductInfo
}

type DeleteAvailableProductsInput struct {
	ProductName             string
	ProductVersion          string
	ShouldDeleteAllProducts bool
}

func (a Api) UploadAvailableProduct(input UploadAvailableProductInput) (UploadAvailableProductOutput, error) {
	req, err := http.NewRequest("POST", availableProductsEndpoint, input.Product)
	if err != nil {
		return UploadAvailableProductOutput{}, err
	}

	req.Header.Set("Content-Type", input.ContentType)
	req.ContentLength = input.ContentLength

	req = req.WithContext(context.WithValue(req.Context(), "polling-interval", time.Duration(input.PollingInterval)*time.Second))

	resp, err := a.progressClient.Do(req)
	if err != nil {
		return UploadAvailableProductOutput{}, fmt.Errorf("could not make api request to available_products endpoint: %s", err)
	}

	defer resp.Body.Close()

	if err = validateStatusOK(resp); err != nil {
		return UploadAvailableProductOutput{}, err
	}

	return UploadAvailableProductOutput{}, nil
}

func (a Api) ListAvailableProducts() (AvailableProductsOutput, error) {
	avReq, err := http.NewRequest("GET", availableProductsEndpoint, nil)
	if err != nil {
		return AvailableProductsOutput{}, err
	}

	resp, err := a.client.Do(avReq)
	if err != nil {
		return AvailableProductsOutput{}, fmt.Errorf("could not make api request to available_products endpoint: %s", err)
	}
	defer resp.Body.Close()

	if err = validateStatusOK(resp); err != nil {
		return AvailableProductsOutput{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AvailableProductsOutput{}, err
	}

	var availableProducts []ProductInfo
	err = json.Unmarshal(respBody, &availableProducts)
	if err != nil {
		return AvailableProductsOutput{}, fmt.Errorf("could not unmarshal available_products response: %s", err)
	}

	return AvailableProductsOutput{ProductsList: availableProducts}, nil
}

func (a Api) DeleteAvailableProducts(input DeleteAvailableProductsInput) error {
	req, err := http.NewRequest("DELETE", availableProductsEndpoint, nil)

	if !input.ShouldDeleteAllProducts {
		query := url.Values{}
		query.Add("product_name", input.ProductName)
		query.Add("version", input.ProductVersion)

		req.URL.RawQuery = query.Encode()
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("could not make api request to available_products endpoint: %s", err)
	}

	defer resp.Body.Close()

	if err = validateStatusOK(resp); err != nil {
		return err
	}

	return nil
}
