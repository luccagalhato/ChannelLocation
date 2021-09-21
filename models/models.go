package models

//AutoGenerated ...
type AutoGenerated struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Bounds struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"bounds"`
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Viewport     struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		PlaceID string   `json:"place_id"`
		Types   []string `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
}

//Response ...
type Response struct {
	Status string
	Error  string
	Data   interface{}
}

// Clients ..
type Client struct {
	Nome     string  `json:"NOME_CLIFOR,omitempty"`
	Endereco string  `json:"ENDERECO,omitempty"`
	Numero   string  `json:"NUMERO,omitempty"`
	Bairro   string  `json:"BAIRRO,omitempty"`
	Cidade   string  `json:"CIDADE,omitempty"`
	Uf       string  `json:"UF,omitempty"`
	Cep      string  `json:"CEP,omitempty"`
	Pais     string  `json:"PAIS,omitempty"`
	Clifor   string  `json:"CLIFOR,omitempty"`
	Lat      *string `json:"LAT,omitempty"`
	Long     *string `json:"LONG,omitempty"`
	Data     *string `json:"DATA_PARA_TRANSFERENCIA,omitempty"`
}

//Clients ...
type Clients struct {
	Nome      string  `json:"NOME_CLIFOR,omitempty"`
	Endereco  string  `json:"ENDERECO,omitempty"`
	Numero    string  `json:"NUMERO,omitempty"`
	Bairro    string  `json:"BAIRRO,omitempty"`
	Cidade    string  `json:"CIDADE,omitempty"`
	Uf        string  `json:"UF,omitempty"`
	Cep       string  `json:"CEP,omitempty"`
	Pais      string  `json:"PAIS,omitempty"`
	Clifor    string  `json:"CLIFOR,omitempty"`
	Latitude  float64 `json:"LATITUDE,omitempty"`
	Longitude float64 `json:"LONGITUDE,omitempty"`
}

//Column ...
type Column struct {
	ColumName  string `json:"name,omitempty"`
	ColumValue string `json:"value,omitempty"`
}

//UpdateRow ...
type UpdateRow struct {
	Columns   []Column `json:"columns,omitempty"`
	Condition string   `json:"condition,omitempty"`
}

//Columns ...
type Columns []Column