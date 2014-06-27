package testheka

import (
	"database/sql"
	"fmt"
	//"encoding/json"
	"errors"
	"strings"
)

type SQLOutputConfig struct {
	Host     string
	Username string
	Password string
	Database string
	Table    string
}

type SQLOutput struct {
	config  *SQLOutputConfig
	DB      *sql.DB
	encoder Encoder
}

func (m *SQLOutput) ConfigStruct() interface{} {
	return &SQLOutputConfig{
		Host:     "localhost:3306",
		Username: "root",
		Password: "mypassword",
		Database: "hekadb",
		Table:    "outt",
	}
}

func (m *SQLOutput) Init(config interface{}) (err error) {
	m.config = config.(*SQLOutputConfig)
	path := m.config.Username + ":" + m.config.Password + "@" + "tcp(" +
		m.config.Host + ")/" + m.config.Database

	db, err = sql.Open("mysql", path)
	if err != nil {
		return fmt.Errorf("Error opening MySql database path '%s': %s", path, err)
	}
	m.DB = db
	fmt.Println("Opened SQL conn")
	return
}

func init() {
	RegisterPlugin("SQLOutput", func() interface{} {
		return new(SQLOutput)
	})
}

func (m *SQLOutput) Run(or OutputRunner, h PluginHelper) (err error) {
	encoder := or.Encoder()
	if encoder == nil {
		return errors.New("Encoder required.")
	}
	var (
		pack     *PipelinePack
		message  string
		outBytes []byte
	)
	inChan := or.InChan()
	for pack = range inChan {
		if outBytes, err = encoder.Encode(pack); err != nil {
			or.LogError(fmt.Errorf("Error encoding message: %s", err))
			pack.Recycle()
			continue
		}
		message = (string(outBytes))
		//  var payload payloadStruct
		// err := json.Unmarshal([]byte(message), &payload)
		query := "insert into out(val1,count) values(1,2)"
		_, err := m.DB.Exec(query)
		if err != nil {
			fmt.Errorf("Error querying '%s': %s", query, err)
			or.LogError(err)
		}
		pack.Recycle()
	}
	return

}
