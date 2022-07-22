package model

type Chart struct {
	Id        string `mapstructure:"_id" bson:"_id,omitempty"`
	Name      string `mapstructure:"name" bson:"name,omitempty"`
	Type      string `mapstructure:"type" bson:"type,omitempty"`
	Tags      string `mapstructure:"tags" bson:"tags,omitempty"`
	DatasetId string `mapstructure:"dataset_id" bson:"dataset_id,omitempty"`
	Inputs    []struct {
		Behavior      string      `mapstructure:"behavior" bson:"behavior,omitempty"`
		Name          string      `mapstructure:"name" bson:"name,omitempty"`
		Label         string      `mapstructure:"label" bson:"label,omitempty"`
		Type          string      `mapstructure:"type" bson:"type,omitempty"`
		Source        interface{} `mapstructure:"source" bson:"source,omitempty" `
		DefaultOption interface{} `mapstructure:"default_option" bson:"default_option,omitempty"`
		DefaultValue  string      `mapstructure:"default_value" bson:"default_value,omitempty"`
		Required      bool        `mapstructure:"required" bson:"required"`
		Params        []struct {
			Name  string `mapstructure:"name" bson:"name,omitempty"`
			Value string `mapstructure:"value" bson:"value,omitempty"`
		} `mapstructure:"params" bson:"params,omitempty"`
	} `mapstructure:"inputs" bson:"inputs,omitempty"`
	Columns []string `mapstructure:"columns" bson:"columns,omitempty"`
	OrderBy []struct {
		Column string `mapstructure:"column" bson:"column,omitempty"`
		Type   string `mapstructure:"type" bson:"type,omitempty"`
	} `mapstructure:"orderby" bson:"orderby,omitempty"`
	Exportable bool `mapstructure:"exportable" bson:"exportable,omitempty"`
}

type DashBoard struct {
	Id        string   `mapstructure:"_id" bson:"_id,omitempty,omitempty"`
	AppId     int32    `mapstructure:"app_id" bson:"app_id,omitempty"`
	Name      string   `mapstructure:"name" bson:"name,omitempty"`
	ChartsIds []string `mapstructure:"charts_ids" bson:"charts_ids,omitempty"`
	Sortby    int32    `mapstructure:"sortby" bson:"sortby,omitempty"`
}

type DataBase struct {
	Id       string `mapstructure:"_id" bson:"_id,omitempty"`
	DbName   string `mapstructure:"db_name" bson:"db_name,omitempty"`
	Username string `mapstructure:"username" bson:"username"`
	Password string `mapstructure:"password" bson:"password"`
	Host     string `mapstructure:"host" bson:"host,omitempty"`
	Port     int32  `mapstructure:"port" bson:"port,omitempty"`
	Driver   string `mapstructure:"driver" bson:"driver,omitempty"`
	Label    string `mapstructure:"label" bson:"label,omitempty"`
}

type DataSet struct {
	Id      string `mapstructure:"_id" bson:"_id,omitempty"`
	Query   string `mapstructure:"query" bson:"query,omitempty"`
	Columns []struct {
		Name       string      `mapstructure:"name" bson:"name,omitempty"`
		Label      string      `mapstructure:"label" bson:"label,omitempty"`
		Transtable interface{} `mapstructure:"transtable" bson:"transtable,omitempty"`
	} `mapstructure:"columns"  bson:"columns,omitempty"`
	DatabaseId string `mapstructure:"database_id"  bson:"database_id,omitempty"`
}

type Transtable struct {
	Id    string `mapstructure:"_id" bson:"_id,omitempty"`
	AppId int32  `mapstructure:"app_id" bson:"app_id,omitempty"`
	Type  string `mapstructure:"type" bson:"type,omitempty"`
	Key   string `mapstructure:"key" bson:"key,omitempty"`
	Value string `mapstructure:"value" bson:"value,omitempty"`
}
