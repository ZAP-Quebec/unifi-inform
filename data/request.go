package data

type InformRequest struct {
	BoardRevision        int         `json:"board_rev"`
	BootRomVersion       string      `json:"bootrom_version"`
	ConfigVersion        string      `json:"cfgversion"`
	CountryCode          int         `json:"country_code"`
	Default              bool        `json:"default"`
	DiscoveryResponse    bool        `json:"discovery_response"`
	Fingerprint          string      `json:"fingerprint"`
	FirmwareCapabilities int         `json:"fw_caps"`
	GuestToken           string      `json:"guest_token"`
	HasEth1              bool        `json:"has_eth1"`
	HasSpeaker           bool        `json:"has_speaker"`
	Hostname             string      `json:"hostname"`
	IntfTable            []Interface `json:"if_table"`
	InformUrl            string      `json:"inform_url"`
	Ip                   string      `json:"ip"`
	LastError            string      `json:"last_error"`
	Locating             bool        `json:"locating"`
	Mac                  MacAddr     `json:"mac"`
	Model                string      `json:"model"`
	ModelDisplay         string      `json:"model_display"`
	Netmask              string      `json:"netmask"`
	QrId                 string      `json:"qrid"`
	RadioTable           []Radio     `json:"radio_table"`
	RequiredVersion      string      `json:"required_version"`
	SelfrunBeacon        bool        `json:"selfrun_beacon"`
	Serial               string      `json:"serial"`
	SpectrumScanning     bool        `json:"spectrum_scanning"`
	State                int         `json:"state"`
	StreamToken          string      `json:"stream_token"`
	SysStats             SysStats    `json:"sys_stats"`
	Time                 int         `json:"time"`
	Uplink               string      `json:"uplink"`
	Uptime               int         `json:"uptime"`
	VApTable             []VAp       `json:"vap_table"`
	Version              string      `json:"version"`
	WifiCapabilities     int         `json:"wifi_caps"`
}
