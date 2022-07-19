package model

type Chart struct {
	//Id        string `json:"_id"`
	Name      string `json:"name" yaml:"name" toml:"name" bson:"name"`
	Type      string `json:"type" yaml:"type" toml:"type" bson:"type"`
	DatasetId string `json:"dataset_id" yaml:"dataset_id" toml:"dataset_id" bson:"dataset_id"`
	Inputs    []struct {
		Name   string `json:"name" yaml:"name"  toml:"name" bson:"name"`
		Label  string `json:"label" yaml:"label" toml:"label" bson:"label"`
		Type   string `json:"type" yaml:"type" toml:"type" bson:"type"`
		Source struct {
			Type  string `json:"type" yaml:"type" toml:"type" bson:"type"`
			Value []struct {
				Label string `json:"label" yaml:"label" toml:"label" bson:"label"`
				Value string `json:"value" yaml:"value" toml:"value" bson:"value"`
			} `json:"value" yaml:"value" toml:"value" bson:"value"`
		} `json:"source,omitempty" yaml:"source" toml:"source" bson:"source" `
		DefaultOption struct {
			Label string `json:"label" yaml:"label" toml:"label" bson:"label"`
			Value string `json:"value" yaml:"label" toml:"value" bson:"value"`
		} `json:"default_option,omitempty" yaml:"default_option" toml:"default_option" bson:"default_option"`
		DefaultValue string `json:"default_value" bson:"default_value" toml:"default_value" bson:"default_value"`
		Required     bool   `json:"required" yaml:"required" toml:"required" bson:"required"`
	} `json:"inputs" yaml:"inputs" toml:"inputs" bson:"inputs"`
	Columns []string `json:"columns" yaml:"columns" toml:"columns" bson:"columns"`
	OrderBy []struct {
		Column string `json:"column" yaml:"column" toml:"columns" bson:"column"`
		Type   string `json:"type" yaml:"type" toml:"type" bson:"type"`
	} `json:"orderby" yaml:"orderby" toml:"orderby" bson:"orderby"`
	Exportable bool `json:"exportable" yaml:"exportable" toml:"exportable" bson:"exportable"`
}

type DashBoard struct {
	//Id        string   `json:"_id"`
	AppId     int32    `json:"app_id" yaml:"app_id" toml:"app_id" bson:"app_id"`
	Name      string   `json:"name" yaml:"name" toml:"name" bson:"name"`
	ChartsIds []string `json:"charts_ids" yaml:"charts_ids" toml:"charts_ids" bson:"charts_ids"`
	Sortby    int32    `json:"sortby" yaml:"sortby" toml:"sortby" bson:"sortby"`
}

type DataBase struct {
	//Id       string `json:"_id"`
	DbName   string `json:"db_name" yaml:"db_name" tmol:"db_name" bson:"db_name"`
	Username string `json:"username" yaml:"username" toml:"username" bson:"username"`
	Password string `json:"password" yaml:"password" toml:"password" bson:"password"`
	Host     string `json:"host" yaml:"host" toml:"host" bson:"host"`
	Port     int32  `json:"port" yaml:"port" toml:"port" bson:"port"`
	Driver   string `json:"driver" yaml:"driver" toml:"driver" bson:"driver"`
	Label    string `json:"label" yaml:"label" toml:"label"  bson:"label"`
}

type DataSet struct {
	Id      string `json:"_id" yaml:"id"  toml:"id" bson:"id"`
	Query   string `json:"query" yaml:"query" toml:"query" bson:"query"`
	Columns []struct {
		Name       string `json:"name" yaml:"name" toml:"name" bson:"name"`
		Label      string `json:"label" yaml:"label" toml:"label" bson:"label"`
		Transtable struct {
			AppId int    `json:"app_id" yaml:"app_id" toml:"app_id" bson:"app_id"`
			Type  string `json:"type" yaml:"type" toml:"type" bson:"type"`
		} `json:"transtable,omitempty" yaml:"transtable" toml:"transtable" bson:"transtable"`
	} `json:"columns" yaml:"columns" toml:"columns" bson:"columns"`
	DatabaseId string `json:"database_id" yaml:"database_id" toml:"database_id" bson:"database_id"`
}

type Transtable struct {
	//Id    string `json:"_id"`
	AppId int32  `json:"app_id" yaml:"app_id" toml:"app_id" bson:"app_id"`
	Type  string `json:"type" yaml:"type" toml:"type" bson:"type"`
	Key   string `json:"key" yaml:"key" toml:"key" bson:"key"`
	Value string `json:"value" yaml:"value" toml:"value" bson:"value"`
}

type ClothTemplate struct {
	/*Id struct {
		Oid string `json:"$oid"`
	} `json:"_id"`*/
	Cid      string `json:"cid" yaml:"cid" toml:"cid" bson:"cid"`
	Name     string `json:"name" yaml:"name" toml:"name" bson:"name"`
	Type     int32  `json:"type" yaml:"type" toml:"type" bson:"type"`
	SuitId   int32  `json:"suitId" yaml:"suitId" toml:"suitId" bson:"suitId"`
	Position string `json:"position" yaml:"position" toml:"position" bson:"position"`
	Scale    string `json:"scale" yaml:"scale" toml:"scale" bson:"scale"`
	Price    int32  `json:"price" yaml:"price" toml:"price" bson:"price"`
	GetType  int32  `json:"getType" yaml:"getType" toml:"getType" bson:"getType"`
	Class    string `json:"_class" yaml:"_class" toml:"_class" bson:"_class"`
}
