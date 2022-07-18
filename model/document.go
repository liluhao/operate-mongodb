package model

//go代码中定义一个Studet类型如下
type ClothTemplate struct {
	Class    string `mapstructure:"_class" json:"_class" yaml:"_class" toml:"_class"`
	Cid      string `mapstructure:"cid" json:"cid" yaml:"cid" toml:"cid"`
	Name     string `mapstructure:"name" json:"name" yaml:"name" toml:"name"`
	Type     int32  `mapstructure:"mongodb" json:"mongodb" yaml:"mongodb" toml:"mongodb"`
	SuitId   int32  `mapstructure:"suitId" json:"suitId" yaml:"suitId" toml:"suitId"`
	Position string `mapstructure:"position" json:"position" yaml:"position" toml:"position"`
	Scale    string `mapstructure:"scale" json:"scale" yaml:"scale" toml:"scale"`
	Price    int32  `mapstructure:"price" json:"price" yaml:"price" toml:"price"`
	GetType  int32  `mapstructure:"getType" json:"getType" yaml:"getType" toml:"getType"`
}
