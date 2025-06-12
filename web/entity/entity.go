package entity

import (
	"crypto/tls"
	"net"
	"strings"

	"x-ui/util/common"
)

type Msg struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Obj     any    `json:"obj"`
}

type AllSetting struct {
	WebListen                   string `json:"webListen" form:"webListen"`
	WebDomain                   string `json:"webDomain" form:"webDomain"`
	WebPort                     int    `json:"webPort" form:"webPort"`
	WebCertFile                 string `json:"webCertFile" form:"webCertFile"`
	WebKeyFile                  string `json:"webKeyFile" form:"webKeyFile"`
	WebBasePath                 string `json:"webBasePath" form:"webBasePath"`
	SessionMaxAge               int    `json:"sessionMaxAge" form:"sessionMaxAge"`
	PageSize                    int    `json:"pageSize" form:"pageSize"`
	ExpireDiff                  int    `json:"expireDiff" form:"expireDiff"`
	TrafficDiff                 int    `json:"trafficDiff" form:"trafficDiff"`
	RemarkModel                 string `json:"remarkModel" form:"remarkModel"`
	TwoFactorEnable				bool   `json:"twoFactorEnable" form:"twoFactorEnable"`
	TwoFactorToken				string `json:"twoFactorToken" form:"twoFactorToken"`
	ExternalTrafficInformEnable bool   `json:"externalTrafficInformEnable" form:"externalTrafficInformEnable"`
	ExternalTrafficInformURI    string `json:"externalTrafficInformURI" form:"externalTrafficInformURI"`
}

func (s *AllSetting) CheckValid() error {
	if s.WebListen != "" {
		ip := net.ParseIP(s.WebListen)
		if ip == nil {
			return common.NewError("web listen is not valid ip:", s.WebListen)
		}
	}

	if s.WebPort <= 0 || s.WebPort > 65535 {
		return common.NewError("web port is not a valid port:", s.WebPort)
	}

	if s.WebCertFile != "" || s.WebKeyFile != "" {
		_, err := tls.LoadX509KeyPair(s.WebCertFile, s.WebKeyFile)
		if err != nil {
			return common.NewErrorf("cert file <%v> or key file <%v> invalid: %v", s.WebCertFile, s.WebKeyFile, err)
		}
	}

	if !strings.HasPrefix(s.WebBasePath, "/") {
		s.WebBasePath = "/" + s.WebBasePath
	}
	if !strings.HasSuffix(s.WebBasePath, "/") {
		s.WebBasePath += "/"
	}

	return nil
}
