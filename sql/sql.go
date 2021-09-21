package sql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	maps "roboInsert/googleMaps"
	"roboInsert/models"

	_ "github.com/denisenkom/go-mssqldb" //bblablalba
)

// SQLStr ...
type SQLStr struct {
	url *url.URL
	db  *sql.DB
}

var connection *SQLStr

//SetSQLConn ...
func SetSQLConn(c *SQLStr) {
	connection = c
}

//SearchClient ...
func (s *SQLStr) SearchClient() error {
	rows, err := s.db.QueryContext(context.Background(), `SELECT LTRIM(RTRIM(NOME_CLIFOR)) AS NOME_CLIFOR,LTRIM(RTRIM(A.ENDERECO)) AS ENDERECO,LTRIM(RTRIM(A.NUMERO)) AS NUMERO,LTRIM(RTRIM(A.BAIRRO)) AS BAIRRO,LTRIM(RTRIM(A.CIDADE)) AS CIDADE,LTRIM(RTRIM(A.UF)) AS UF,LTRIM(RTRIM(A.CEP)) AS CEP,LTRIM(RTRIM(A.PAIS)) AS PAIS,LTRIM(RTRIM(A.CLIFOR)) AS CLIFOR, LTRIM(RTRIM(B.LAT)) AS LAT, LTRIM(RTRIM(B.LONG)) AS LONG, b.DATA_PARA_TRANSFERENCIA
	FROM (SELECT * FROM LINX_TBFG..CADASTRO_CLI_FOR WHERE INDICA_CLIENTE='1' AND PJ_PF = '1') A 
	LEFT JOIN
	(SELECT CLIENTE_ATACADO, CAST("01203" AS FLOAT) AS LAT, CAST("01204" AS FLOAT) AS LONG, DATA_PARA_TRANSFERENCIA
	FROM
	(
	  SELECT CLIENTE_ATACADO, VALOR_PROPRIEDADE, PROPRIEDADE, DATA_PARA_TRANSFERENCIA
	  FROM LINX_TBFG..PROP_CLIENTES_ATACADO
	  WHERE PROPRIEDADE IN ('01203', '01204')
	) d
	PIVOT
	(
	  max(VALOR_PROPRIEDADE)
	  FOR PROPRIEDADE IN ("01203", "01204")
	) piv) B ON A.NOME_CLIFOR=B.CLIENTE_ATACADO
	WHERE A.DATA_PARA_TRANSFERENCIA>B.DATA_PARA_TRANSFERENCIA OR B.DATA_PARA_TRANSFERENCIA IS NULL`, nil)
	if err != nil {
		// fmt.Println(err)
		return nil
	}
	for rows.Next() {
		client := models.Client{}
		if err := rows.Scan(&client.Nome, &client.Endereco, &client.Numero, &client.Bairro, &client.Cidade, &client.Uf, &client.Cep, &client.Pais, &client.Clifor, &client.Lat, &client.Long, &client.Data); err != nil {
			fmt.Println(err)
			continue
		}
		lat, long := maps.RequestMapsNewclient(client)
		if client.Data != nil {
			connection.UpdateRow(fmt.Sprintf("%f", lat), "01203")
			connection.UpdateRow(fmt.Sprintf("%f", long), "01204")
			continue
		}
		connection.InsertRow(fmt.Sprintf("%f", lat), fmt.Sprintf("%f", long), client.Nome)
	}
	return nil
}
func (s *SQLStr) InsertRow(lat, long string, nome string) {
	_, err := s.db.QueryContext(context.Background(), fmt.Sprintf(`insert into LINX_TBFG..PROP_CLIENTES_ATACADO (PROPRIEDADE,CLIENTE_ATACADO,ITEM_PROPRIEDADE, VALOR_PROPRIEDADE)
	VALUES 
	('01203', '%s', 1, '%s'),
	('01204', '%s', 1, '%s);`, nome, lat, nome, long))
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (s *SQLStr) UpdateRow(value string, condition string) {
	_, err := s.db.QueryContext(context.Background(), fmt.Sprintf(update, table, "VALOR_PROPRIEDADE ="+value, "PROPRIEDADE="+condition))
	if err != nil {
		fmt.Println(err)
		return
	}
}

//MakeSql ...
func MakeSQL(host, port, username, password string) (*SQLStr, error) {

	s := &SQLStr{}
	s.url = &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(username, password),
		Host:     fmt.Sprintf("%s:%s", host, port),
		RawQuery: url.Values{}.Encode(),
	}
	return s, s.Connect()
}

//Connect ...
func (s *SQLStr) Connect() error {
	var err error
	if s.db, err = sql.Open("sqlserver", s.url.String()); err != nil {
		return err
	}
	return s.db.PingContext(context.Background())
}

const (
	update = "UPDATE %s SET %s WHERE %s"
)
const table = "LINX_TBFG..PROP_CLIENTES_ATACADO"

// func (s *SQLStr) disconnect() error {
