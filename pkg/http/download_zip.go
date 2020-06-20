package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Download(url string) ([]byte, error){
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func DownloadB3HistoryZip(year uint) ([]byte, error){
	zip, err := Download(fmt.Sprintf("http://bvmf.bmfbovespa.com.br/InstDados/SerHist/COTAHIST_A%v.ZIP", year))
	if err != nil {
		return nil, err
	}

	return zip, err
}