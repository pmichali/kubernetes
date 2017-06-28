/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"fmt"
	"net"
	"strconv"

	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
)

// GetMasterEndpoint returns a properly formatted Master Endpoint
// or error if the IP address can not be parsed.
func GetMasterEndpoint(cfg *kubeadmapi.MasterConfiguration) (string, error) {
	masterIP := net.ParseIP(cfg.API.AdvertiseAddress)
	if masterIP == nil {
		return "", fmt.Errorf("error parsing address %s", cfg.API.AdvertiseAddress)
	}
	masterEndpoint := net.JoinHostPort(masterIP.String(), strconv.Itoa(int(cfg.API.BindPort)))
	return fmt.Sprintf("https://%s", masterEndpoint), nil
}
