package inform

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

type Client struct {
	mac     MAC
	key     Key
	address string
	http    *http.Client
}

func NewClient(m MAC, address string) *Client {
	if !m.IsValid() {
		panic("Invalid mac address")
	}
	return &Client{
		mac:     m,
		key:     DEFAULT_KEY,
		address: address,
		http: &http.Client{
			Transport: &http.Transport{
				DisableCompression: true,
				DisableKeepAlives:  true,
			},
		},
	}
}

func (c *Client) SendInform() (InformResponse, error) {
	p := NewPacket(c.mac, &fakeInform{}, c.key)
	body, err := p.Marshal()
	if err != nil {
		return nil, err
	}

	fmt.Printf("flags : %d \n", p.flags)

	r := bytes.NewReader(body)
	req, err := http.NewRequest("POST", c.address, r)
	if err != nil {
		return nil, err
	}

	req.Host = "192.168.1.15"
	//setHeader(req, "Host", )
	setHeader(req, "user-agent", "AirControl Agent v1.0")
	setHeader(req, "content-type", "application/x-binary")
	setHeader(req, "content-length", strconv.Itoa(len(body)))

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.StatusCode, resp.Status)

	return nil, nil
}

func setHeader(r *http.Request, key, value string) {
	r.Header.Del(key)

	//r.Header[key] = []string{value}
	r.Header.Set(key, value)
}

func (c *Client) StartDiscovery() {
	// TODO
	// see : https://github.com/fxkr/unifi-protocol-reverse-engineering#discover
}

type fakeInform struct{}

func (f fakeInform) Marshal() []byte {
	return []byte(f.String())
}

func (fakeInform) String() string {
	return `{
    "board_rev": 31,
    "bootrom_version": "unifi-v1.5.2.206-g44e4c8bc",
    "cfgversion": "?",
    "country_code": 0,
    "default": true,
    "discovery_response": true,
    "fingerprint": "b0:81:23:fe:07:a6:4a:28:64:5a:e0:a6:b1:d8:6f:40",
    "fw_caps": 75,
    "guest_token": "F341AB07BBEC990D0F46010B9654685E",
    "has_eth1": false,
    "has_speaker": false,
    "hostname": "UBNT",
    "if_table": [
        {
            "full_duplex": true,
            "ip": "0.0.0.0",
            "mac": "04:18:d6:e0:0f:af",
            "name": "eth0",
            "netmask": "0.0.0.0",
            "num_port": 2,
            "rx_bytes": 44763,
            "rx_dropped": 0,
            "rx_errors": 0,
            "rx_multicast": 93,
            "rx_packets": 316,
            "speed": 100,
            "tx_bytes": 20625,
            "tx_dropped": 0,
            "tx_errors": 0,
            "tx_packets": 169,
            "up": true
        }
    ],
    "inform_url": "http://192.168.1.15:8080/inform",
    "ip": "192.168.1.13",
    "isolated": false,
    "last_error": "Unable to resolve (http://unifi:8080/inform)",
    "locating": false,
    "mac": "04:18:d6:e0:0f:af",
    "model": "U7P",
    "model_display": "UAP-Pro",
    "netmask": "255.255.255.0",
    "qrid": "000000",
    "radio_table": [
        {
            "athstats": {
                "ast_ath_reset": 0,
                "ast_be_xmit": 0,
                "ast_cst": 0,
                "ast_deadqueue_reset": 0,
                "ast_fullqueue_stop": 0,
                "ast_txto": 0,
                "cu_self_rx": 0,
                "cu_self_tx": 0,
                "cu_total": 0,
                "n_rx_aggr": 0,
                "n_rx_pkts": 85,
                "n_tx_bawadv": 0,
                "n_tx_bawretries": 0,
                "n_tx_pkts": 36,
                "n_tx_queue": 0,
                "n_tx_retries": 0,
                "n_tx_xretries": 0,
                "n_txaggr_compgood": 0,
                "n_txaggr_compretries": 0,
                "n_txaggr_compxretry": 0,
                "n_txaggr_prepends": 0,
                "name": "wifi0"
            },
            "builtin_ant_gain": 0,
            "builtin_antenna": true,
            "has_dfs": true,
            "max_txpower": 23,
            "min_txpower": 4,
            "name": "wifi0",
            "nss": 2,
            "radio": "na",
            "scan_table": []
        },
        {
            "athstats": {
                "ast_ath_reset": 0,
                "ast_be_xmit": 2985,
                "ast_cst": 2,
                "ast_deadqueue_reset": 0,
                "ast_fullqueue_stop": 0,
                "ast_txto": 0,
                "cu_self_rx": 15,
                "cu_self_tx": 0,
                "cu_total": 18,
                "n_rx_aggr": 0,
                "n_rx_pkts": 19853,
                "n_tx_bawadv": 0,
                "n_tx_bawretries": 0,
                "n_tx_pkts": 0,
                "n_tx_queue": 0,
                "n_tx_retries": 0,
                "n_tx_xretries": 0,
                "n_txaggr_compgood": 0,
                "n_txaggr_compretries": 0,
                "n_txaggr_compxretry": 0,
                "n_txaggr_prepends": 0,
                "name": "wifi1"
            },
            "builtin_ant_gain": 0,
            "builtin_antenna": true,
            "max_txpower": 30,
            "min_txpower": 12,
            "name": "wifi1",
            "nss": 3,
            "radio": "ng",
            "scan_table": []
        }
    ],
    "required_version": "2.4.4",
    "selfrun_beacon": true,
    "serial": "0418D6E00FAF",
    "spectrum_scanning": false,
    "state": 1,
    "stream_token": "",
    "sys_stats": {
        "loadavg_1": "0.01",
        "loadavg_15": "0.04",
        "loadavg_5": "0.07",
        "mem_buffer": 0,
        "mem_total": 130375680,
        "mem_used": 47947776
    },
    "time": 1494091704,
    "uplink": "eth0",
    "uptime": 346,
    "vap_table": [
        {
            "bssid": "ff:ff:ff:ff:ff:ff",
            "ccq": 5203672,
            "channel": 36,
            "essid": "",
            "id": "user",
            "name": "ath0",
            "num_sta": 0,
            "radio": "na",
            "rx_bytes": 0,
            "rx_crypts": 0,
            "rx_dropped": 0,
            "rx_errors": 0,
            "rx_frags": 0,
            "rx_nwids": 0,
            "rx_packets": 0,
            "sta_table": [],
            "state": "INIT",
            "tx_bytes": 0,
            "tx_dropped": 2,
            "tx_errors": 0,
            "tx_packets": 0,
            "tx_power": 23,
            "tx_retries": 0,
            "up": false,
            "usage": "uplink"
        },
        {
            "bssid": "04:18:d6:e2:0f:af",
            "ccq": 5203672,
            "channel": 1,
            "essid": "0418D6E00FAF",
            "id": "user",
            "name": "ath1",
            "num_sta": 0,
            "radio": "ng",
            "rx_bytes": 0,
            "rx_crypts": 0,
            "rx_dropped": 0,
            "rx_errors": 0,
            "rx_frags": 0,
            "rx_nwids": 763,
            "rx_packets": 0,
            "sta_table": [],
            "state": "RUN",
            "tx_bytes": 0,
            "tx_dropped": 151,
            "tx_errors": 0,
            "tx_packets": 0,
            "tx_power": 30,
            "tx_retries": 0,
            "up": true,
            "usage": "user"
        }
    ],
    "version": "3.7.39.6089-custom",
    "wifi_caps": 117
}`
}
