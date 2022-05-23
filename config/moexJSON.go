package config

type MOEXCURRENT struct {
	Securities []struct {
		Secid               string      `json:"SECID"`
		Boardid             string      `json:"BOARDID"`
		Shortname           string      `json:"SHORTNAME"`
		Prevprice           float64     `json:"PREVPRICE"`
		Lotsize             int         `json:"LOTSIZE"`
		Facevalue           float64     `json:"FACEVALUE"`
		Status              string      `json:"STATUS"`
		Boardname           string      `json:"BOARDNAME"`
		Decimals            int         `json:"DECIMALS"`
		Secname             string      `json:"SECNAME"`
		Remarks             interface{} `json:"REMARKS"`
		Marketcode          string      `json:"MARKETCODE"`
		Instrid             string      `json:"INSTRID"`
		Sectorid            interface{} `json:"SECTORID"`
		Minstep             float64     `json:"MINSTEP"`
		Prevwaprice         float64     `json:"PREVWAPRICE"`
		Faceunit            string      `json:"FACEUNIT"`
		Prevdate            string      `json:"PREVDATE"`
		Issuesize           int64       `json:"ISSUESIZE"`
		Isin                string      `json:"ISIN"`
		Latname             string      `json:"LATNAME"`
		Regnumber           string      `json:"REGNUMBER"`
		Prevlegalcloseprice float64     `json:"PREVLEGALCLOSEPRICE"`
		Prevadmittedquote   float64     `json:"PREVADMITTEDQUOTE"`
		Currencyid          string      `json:"CURRENCYID"`
		Sectype             string      `json:"SECTYPE"`
		Listlevel           int         `json:"LISTLEVEL"`
		Settledate          string      `json:"SETTLEDATE"`
	} `json:"securities"`
	Marketdata []struct {
		Secid                         string      `json:"SECID"`
		Boardid                       string      `json:"BOARDID"`
		Bid                           interface{} `json:"BID"`
		Biddepth                      interface{} `json:"BIDDEPTH"`
		Offer                         interface{} `json:"OFFER"`
		Offerdepth                    interface{} `json:"OFFERDEPTH"`
		Spread                        float64     `json:"SPREAD"`
		Biddeptht                     int         `json:"BIDDEPTHT"`
		Offerdeptht                   int         `json:"OFFERDEPTHT"`
		Open                          float64     `json:"OPEN"`
		Low                           float64     `json:"LOW"`
		High                          float64     `json:"HIGH"`
		Last                          float64     `json:"LAST"`
		Lastchange                    float64     `json:"LASTCHANGE"`
		Lastchangeprcnt               float64     `json:"LASTCHANGEPRCNT"`
		Qty                           int         `json:"QTY"`
		Value                         float64     `json:"VALUE"`
		ValueUsd                      float64     `json:"VALUE_USD"`
		Waprice                       float64     `json:"WAPRICE"`
		Lastcngtolastwaprice          float64     `json:"LASTCNGTOLASTWAPRICE"`
		Waptoprevwapriceprcnt         float64     `json:"WAPTOPREVWAPRICEPRCNT"`
		Waptoprevwaprice              float64     `json:"WAPTOPREVWAPRICE"`
		Closeprice                    interface{} `json:"CLOSEPRICE"`
		Marketpricetoday              float64     `json:"MARKETPRICETODAY"`
		Marketprice                   float64     `json:"MARKETPRICE"`
		Lasttoprevprice               float64     `json:"LASTTOPREVPRICE"`
		Numtrades                     int64       `json:"NUMTRADES"`
		Voltoday                      int         `json:"VOLTODAY"`
		Valtoday                      int64       `json:"VALTODAY"`
		ValtodayUsd                   int         `json:"VALTODAY_USD"`
		Etfsettleprice                interface{} `json:"ETFSETTLEPRICE"`
		Tradingstatus                 string      `json:"TRADINGSTATUS"`
		Updatetime                    string      `json:"UPDATETIME"`
		Admittedquote                 float64     `json:"ADMITTEDQUOTE"`
		Lastbid                       interface{} `json:"LASTBID"`
		Lastoffer                     interface{} `json:"LASTOFFER"`
		Lcloseprice                   float64     `json:"LCLOSEPRICE"`
		Lcurrentprice                 float64     `json:"LCURRENTPRICE"`
		Marketprice2                  float64     `json:"MARKETPRICE2"`
		Numbids                       interface{} `json:"NUMBIDS"`
		Numoffers                     interface{} `json:"NUMOFFERS"`
		Change                        float64     `json:"CHANGE"`
		Time                          string      `json:"TIME"`
		Highbid                       interface{} `json:"HIGHBID"`
		Lowoffer                      interface{} `json:"LOWOFFER"`
		Priceminusprevwaprice         float64     `json:"PRICEMINUSPREVWAPRICE"`
		Openperiodprice               float64     `json:"OPENPERIODPRICE"`
		Seqnum                        int         `json:"SEQNUM"`
		Systime                       string      `json:"SYSTIME"`
		Closingauctionprice           float64     `json:"CLOSINGAUCTIONPRICE"`
		Closingauctionvolume          int         `json:"CLOSINGAUCTIONVOLUME"`
		Issuecapitalization           float64     `json:"ISSUECAPITALIZATION"`
		IssuecapitalizationUpdatetime string      `json:"ISSUECAPITALIZATION_UPDATETIME"`
		Etfsettlecurrency             interface{} `json:"ETFSETTLECURRENCY"`
		ValtodayRur                   float64     `json:"VALTODAY_RUR"`
		Tradingsession                interface{} `json:"TRADINGSESSION"`
	} `json:"marketdata"`
	Dataversion []struct {
		DataVersion int `json:"data_version"`
		Seqnum      int `json:"seqnum"`
	} `json:"dataversion"`
	MarketdataYields []interface{} `json:"marketdata_yields"`
}

type LASTVOLUME struct {
	Charsetinfo struct {
		Name string `json:"name"`
	} `json:"charsetinfo,omitempty"`
	History []struct {
		Boardid                 string      `json:"BOARDID"`
		Tradedate               string      `json:"TRADEDATE"`
		Shortname               string      `json:"SHORTNAME"`
		Secid                   string      `json:"SECID"`
		Numtrades               int64       `json:"NUMTRADES"`
		Value                   float64     `json:"VALUE"`
		Open                    interface{} `json:"OPEN"`
		Low                     interface{} `json:"LOW"`
		High                    interface{} `json:"HIGH"`
		Legalcloseprice         float64     `json:"LEGALCLOSEPRICE"`
		Waprice                 interface{} `json:"WAPRICE"`
		Close                   float64     `json:"CLOSE"`
		Volume                  int64       `json:"VOLUME"`
		Marketprice2            interface{} `json:"MARKETPRICE2"`
		Marketprice3            interface{} `json:"MARKETPRICE3"`
		Admittedquote           float64     `json:"ADMITTEDQUOTE"`
		Mp2Valtrd               float64     `json:"MP2VALTRD"`
		Marketprice3Tradesvalue float64     `json:"MARKETPRICE3TRADESVALUE"`
		Admittedvalue           float64     `json:"ADMITTEDVALUE"`
		Waval                   int         `json:"WAVAL"`
		Tradingsession          int         `json:"TRADINGSESSION"`
	} `json:"history,omitempty"`
	HistoryCursor []struct {
		Index    int `json:"INDEX"`
		Total    int `json:"TOTAL"`
		Pagesize int `json:"PAGESIZE"`
	} `json:"history.cursor,omitempty"`
}
