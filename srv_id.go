package client

import "github.com/eden-framework/client/client_srv_id"

func GetUniqueID(client client_srv_id.ClientSrvIDInterface) (uint64, error) {
	resp, err := client.GenerateID()
	if err != nil {
		return 0, err
	}
	return resp.Body.ID, nil
}
