package framework

// Config 配置檔
type Config struct {
	Title     string               `toml:"title"`
	Servers   map[string]ServerVer `toml:"servers"`
	Owenor    `toml:"owenor"`
	RPCServer `toml:"rpc_server"`
	// Clients   struct {
	// 	Data [][]string `toml:"data"`
	// } `toml:"clients"`
}

type Owenor struct {
	Name    string `toml:"name"`
	Account string `toml:"account"`
}

type RPCServer struct {
	Server        string `toml:"server"`
	Port          string `toml:"port"`
	ConnectionMax int    `toml:"connection_max"`
	Enabled       bool   `toml:"enabled"`
}

type ServerVer struct {
	IP string `toml:"ip"`
	Dc string `toml:"dc"`
}
