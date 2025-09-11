package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct {
	Symbol string `json:"symbol" jsonschema:"the symbol of the stock to get"`
}

type Output struct {
	Symbol                     string `json:"symbol" jsonschema:"the symbol of the stock"`
	AssetType                  string `json:"assetType" jsonschema:"the type of the asset"`
	Name                       string `json:"name" jsonschema:"the name of the stock"`
	Description                string `json:"description" jsonschema:"the description of the stock"`
	CIK                        string `json:"CIK" jsonschema:"the CIK of the stock"`
	Exchange                   string `json:"exchange" jsonschema:"the exchange of the stock"`
	Currency                   string `json:"currency" jsonschema:"the currency of the stock"`
	Country                    string `json:"country" jsonschema:"the country of the stock"`
	Sector                     string `json:"sector" jsonschema:"the sector of the stock"`
	Industry                   string `json:"industry" jsonschema:"the industry of the stock"`
	Address                    string `json:"address" jsonschema:"the address of the stock"`
	OfficialName               string `json:"officialName" jsonschema:"the official name of the stock"`
	FiscalYearEnd              string `json:"fiscalYearEnd" jsonschema:"the fiscal year end of the stock"`
	LatestQuarter              string `json:"latestQuarter" jsonschema:"the latest quarter of the stock"`
	MarketCapitalization       string `json:"marketCapitalization" jsonschema:"the market capitalization of the stock"`
	EBITDA                     string `json:"ebitda" jsonschema:"the EBITDA of the stock"`
	PERatio                    string `json:"peRatio" jsonschema:"the PERatio of the stock"`
	PEGRatio                   string `json:"pegRatio" jsonschema:"the PEGRatio of the stock"`
	BookValue                  string `json:"bookValue" jsonschema:"the BookValue of the stock"`
	DividendPerShare           string `json:"dividendPerShare" jsonschema:"the DividendPerShare of the stock"`
	DividendYield              string `json:"dividendYield" jsonschema:"the DividendYield of the stock"`
	EPS                        string `json:"eps" jsonschema:"the EPS of the stock"`
	RevenuePerShareTTM         string `json:"revenuePerShareTTM" jsonschema:"the RevenuePerShareTTM of the stock"`
	ProfitMargin               string `json:"profitMargin" jsonschema:"the ProfitMargin of the stock"`
	OperatingMarginTTM         string `json:"operatingMarginTTM" jsonschema:"the OperatingMarginTTM of the stock"`
	ReturnOnAssetsTTM          string `json:"returnOnAssetsTTM" jsonschema:"the ReturnOnAssetsTTM of the stock"`
	ReturnOnEquityTTM          string `json:"returnOnEquityTTM" jsonschema:"the ReturnOnEquityTTM of the stock"`
	RevenueTTM                 string `json:"revenueTTM" jsonschema:"the RevenueTTM of the stock"`
	GrossProfitTTM             string `json:"grossProfitTTM" jsonschema:"the GrossProfitTTM of the stock"`
	DilutedEPSTTM              string `json:"dilutedEPSTTM" jsonschema:"the DilutedEPSTTM of the stock"`
	QuarterlyEarningsGrowthYOY string `json:"quarterlyEarningsGrowthYOY" jsonschema:"the QuarterlyEarningsGrowthYOY of the stock"`
	QuarterlyRevenueGrowthYOY  string `json:"quarterlyRevenueGrowthYOY" jsonschema:"the QuarterlyRevenueGrowthYOY of the stock"`
	AnalystTargetPrice         string `json:"analystTargetPrice" jsonschema:"the AnalystTargetPrice of the stock"`
	AnalystRatingStrongBuy     string `json:"analystRatingStrongBuy" jsonschema:"the AnalystRatingStrongBuy of the stock"`
	AnalystRatingBuy           string `json:"analystRatingBuy" jsonschema:"the AnalystRatingBuy of the stock"`
	AnalystRatingHold          string `json:"analystRatingHold" jsonschema:"the AnalystRatingHold of the stock"`
	AnalystRatingSell          string `json:"analystRatingSell" jsonschema:"the AnalystRatingSell of the stock"`
	AnalystRatingStrongSell    string `json:"analystRatingStrongSell" jsonschema:"the AnalystRatingStrongSell of the stock"`
	TrailingPE                 string `json:"trailingPE" jsonschema:"the TrailingPE of the stock"`
	ForwardPE                  string `json:"forwardPE" jsonschema:"the ForwardPE of the stock"`
	PriceToSalesRatioTTM       string `json:"priceToSalesRatioTTM" jsonschema:"the PriceToSalesRatioTTM of the stock"`
	PriceToBookRatio           string `json:"priceToBookRatio" jsonschema:"the PriceToBookRatio of the stock"`
	EVToRevenue                string `json:"evToRevenue" jsonschema:"the EVToRevenue of the stock"`
	EVToEBITDA                 string `json:"evToEBITDA" jsonschema:"the EVToEBITDA of the stock"`
	Beta                       string `json:"beta" jsonschema:"the Beta of the stock"`
	FiftyTwoWeekHigh           string `json:"fiftyTwoWeekHigh" jsonschema:"the FiftyTwoWeekHigh of the stock"`
	FiftyTwoWeekLow            string `json:"fiftyTwoWeekLow" jsonschema:"the FiftyTwoWeekLow of the stock"`
	FiftyDayMovingAverage      string `json:"fiftyDayMovingAverage" jsonschema:"the FiftyDayMovingAverage of the stock"`
	TwoHundredDayMovingAverage string `json:"twoHundredDayMovingAverage" jsonschema:"the TwoHundredDayMovingAverage of the stock"`
	SharesOutstanding          string `json:"sharesOutstanding" jsonschema:"the SharesOutstanding of the stock"`
	SharesFloat                string `json:"sharesFloat" jsonschema:"the SharesFloat of the stock"`
	PercentInsiders            string `json:"percentInsiders" jsonschema:"the PercentInsiders of the stock"`
	PercentInstitutions        string `json:"percentInstitutions" jsonschema:"the PercentInstitutions of the stock"`
	DividendDate               string `json:"dividendDate" jsonschema:"the DividendDate of the stock"`
	ExDividendDate             string `json:"exDividendDate" jsonschema:"the ExDividendDate of the stock"`
}

type MarketData struct {
	APIURL string `json:"apiURL" jsonschema:"the APIURL of the stock"`
	APIKey string `json:"apiKey" jsonschema:"the APIKey of the stock"`
}

func NewMarketData(apiURL, apiKey string) *MarketData {
	return &MarketData{
		APIURL: apiURL,
		APIKey: apiKey,
	}
}

func (md *MarketData) GetStock(ctx context.Context, req *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, Output, error) {
	res, err := http.Get(fmt.Sprintf("%s/query?function=OVERVIEW&symbol=%s&apikey=%s", md.APIURL, strings.ToUpper(input.Symbol), md.APIKey))
	if err != nil {
		return nil, Output{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, Output{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var data Output

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, Output{}, err
	}

	return nil, data, nil
}

type Config struct {
	APIURL string `json:"api_url" jsonschema:"the API URL for the Alpha Vantage API"`
	APIKey string `json:"api_key" jsonschema:"the API key for the Alpha Vantage API"`
}

func NewConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		APIURL: os.Getenv("API_URL"),
		APIKey: os.Getenv("API_KEY"),
	}
}

func main() {
	cfg := NewConfig()
	impl := &mcp.Implementation{
		Name:    "Market-Data",
		Version: "v1.0.0",
	}
	ctx := context.Background()
	server := mcp.NewServer(impl, nil)
	md := NewMarketData(cfg.APIURL, cfg.APIKey)

	mcp.AddTool(server, &mcp.Tool{Name: "get-stock", Description: "Get specific stock into market using Stock symbol"}, md.GetStock)

	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
