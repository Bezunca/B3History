package b3

import (
	"github.com/Bezunca/B3History/pkg/http"
	"github.com/Bezunca/B3History/pkg/internal/parser"
	"github.com/Bezunca/B3History/pkg/models"
	"github.com/Bezunca/ZipInMemory/pkg/zip"
)

func GetHistory(year uint) ([]models.AssetInfo, error){
	responseData, err := http.DownloadB3HistoryZip(year)
	if err != nil {
		return nil, err
	}
	encodedContent, err := zip.ExtractInMemory(responseData)
	if err != nil {
		return nil, err
	}
	data, err := parser.ParseHistoricDataFromBytes(encodedContent, int(year))
	if err != nil {
		return nil, err
	}

	return data, nil
}