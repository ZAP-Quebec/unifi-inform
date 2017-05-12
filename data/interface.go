package data

type Interface struct {
	full_duplex  bool    `json:"full_duplex"`
	ip           string  `json:"ip"`
	mac          MacAddr `json:"mac"`
	name         string  `json:"name"`
	netmask      string  `json:"netmask"`
	num_port     int     `json:"num_port"`
	rx_bytes     int     `json:"rx_bytes"`
	rx_dropped   int     `json:"rx_dropped"`
	rx_errors    int     `json:"rx_errors"`
	rx_multicast int     `json:"rx_multicast"`
	rx_packets   int     `json:"rx_packets"`
	speed        int     `json:"speed"`
	tx_bytes     int     `json:"tx_bytes"`
	tx_dropped   int     `json:"tx_dropped"`
	tx_errors    int     `json:"tx_errors"`
	tx_packets   int     `json:"tx_packets"`
	up           bool    `json:"up"`
}
