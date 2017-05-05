package hop

import (
    "errors"
    "code.google.com/p/gcfg"
)
// Server Config
type HopServerConfig struct {
    HopStart int
    HopEnd int
    ListenAddr string
    Addr string
    MTU int
    Key string
    FixMSS bool
    MorphMethod string
}

// Client Config
type HopClientConfig struct {
    Server string
    HopStart int
    HopEnd int
    Key string
    MTU int
    FixMSS bool
    Local bool
    MorphMethod string
    Redirect_gateway bool
    Net_gateway []string
}

type HopConfig struct {
    Default struct {
        Mode string
    }
    Server HopServerConfig
    Client HopClientConfig
}

func ParseHopConfig(filename string) (interface{}, error) {
    cfg := new(HopConfig)
    err := gcfg.ReadFileInto(cfg, filename)
    if err != nil {
        return nil, err
    }
    switch cfg.Default.Mode {
    case "server":
        return cfg.Server, nil
    case "client":
        return cfg.Client, nil
    default:
        return nil, errors.New("Wrong mode")
    }
}
