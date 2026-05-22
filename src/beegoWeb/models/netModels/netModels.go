package netModels

type NetProxy struct {
	ID        int64  `bson:"id" json:"id"`
	SrcIp     string `bson:"src_ip" json:"src_ip"`
	SrcPort   string `bson:"src_port" json:"src_port"`
	IsoCode   string `bson:"iso_code" json:"iso_code"`
	ProxyIp   string `bson:"proxy_ip" json:"proxy_ip"`
	ProxyPort string `bson:"proxy_port" json:"proxy_port"`
	Priority  int8   `bson:"priority" json:"priority"`
	Status    string `bson:"status" json:"status"`
	Desc      string `bson:"desc" json:"desc"`
}

type NetProxyList struct {
	ProxyList []*NetProxy
}
