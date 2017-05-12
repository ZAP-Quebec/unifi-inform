package inform

import (
	"bytes"
	"errors"
	messages "github.com/ZAP-Quebec/unifi-inform/data"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Client struct {
	mac     messages.MacAddr
	key     messages.Key
	address string
}

var httpClient = &http.Client{
	Transport: &http.Transport{
		DisableCompression: true,
		DisableKeepAlives:  true,
	},
}

func NewClient(m messages.MacAddr, address string) *Client {
	if !m.IsValid() {
		panic("Invalid mac address")
	}
	return &Client{
		mac:     m,
		key:     messages.DEFAULT_KEY,
		address: address,
	}
}

func (c *Client) SendInform() (messages.InformResponse, error) {
	p := NewPacket(c.mac, &fakeInform{
		m: c.mac,
	}, c.key)
	body, err := p.Marshal()
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	req, err := http.NewRequest("POST", c.address, r)
	if err != nil {
		return nil, err
	}

	addr, err := url.Parse(c.address)
	if err != nil {
		return nil, err
	}

	req.Host = addr.Hostname()
	setHeader(req, "user-agent", "AirControl Agent v1.0")
	setHeader(req, "content-type", "application/x-binary")
	setHeader(req, "content-length", strconv.Itoa(len(body)))

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 || resp.Header.Get("Content-Type") != "application/x-binary" {
		return messages.ResponseFromHttpCode(resp.StatusCode), nil
	}

	data, err := ioutil.ReadAll(resp.Body)

	rPacket := &Packet{}

	err = rPacket.Unmarshal(data, func(ap messages.MacAddr) (messages.Key, error) {
		return messages.DEFAULT_KEY, nil
	})
	if err != nil {
		return nil, err
	}

	if informResp, ok := rPacket.Msg.(messages.InformResponse); !ok {
		return nil, errors.New("Invalid")
	} else {
		return informResp, nil
	}
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

type fakeInform struct {
	m messages.MacAddr
}

func (f fakeInform) Marshal() []byte {
	return []byte(f.String())
}

func (f fakeInform) String() string {
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
            "mac": "` + f.m.String() + `",
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
    "mac": "` + f.m.String() + `",
    "model": "U7P",
    "model_display": "UAP-Pro",
    "netmask": "255.255.255.0",
    "qrid": "000000",
    "radio_table": [],
    "required_version": "2.4.4",
    "selfrun_beacon": true,
    "serial": "` + strings.ToUpper(f.m.HexString()) + `",
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
    "vap_table": [],
    "version": "3.7.39.6089-custom3",
    "wifi_caps": 117
}`
}
