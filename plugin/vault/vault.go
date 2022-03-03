package vault

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hashicorp/vault/api"

	pb "github.com/regen-network/keystone2/keystone"
	krplugin "github.com/regen-network/keystone2/plugin"
)

type vaultConfig struct {
	Token string `json:"token"`
	Addr  string `json:"addr"`
}

type keyring struct {
	vaultClient *api.Client
}

// Init initializes this keyring using the passed in file path
// which should implement the KeyringPlugin interface @@TODO
func Init(configPath string) (krplugin.Plugin, error) {
	var config vaultConfig
	jsonConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonConfig, &config)
	if err != nil {
		return nil, err
	}
	client, err := api.NewClient(&api.Config{
		Address: config.Addr,
	})
	if err != nil {
		return nil, err
	}
	client.SetToken(config.Token)
	return &keyring{
		vaultClient: client,
	}, nil
}

// NewKey creates a new private key(pair) on the existing keyring that
func (*keyring) NewKey(in *pb.KeySpec) (*pb.KeyRef, error) {
	return nil, nil
}

//PutKey stores a key in the keyring
func (*keyring) PubKey(in *pb.KeySpec) (*pb.PublicKey, error) {
	return nil, nil
}

func (*keyring) Sign(in *pb.Msg) (*pb.Signed, error) {
	return nil, nil
}
