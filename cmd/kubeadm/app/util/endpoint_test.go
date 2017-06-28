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
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"testing"
)

func TestGetMasterEndpoint(t *testing.T) {
	var tests = []struct {
		name     string
		cfg      *kubeadmapi.MasterConfiguration
		endpoint string
		expected bool
	}{
		{
			name: "Valid IPv4 endpoint",
			cfg: &kubeadmapi.MasterConfiguration{
				API: kubeadmapi.API{
					AdvertiseAddress: "1.2.3.4",
					BindPort:         1234,
				},
			},
			endpoint: "https://1.2.3.4:1234",
			expected: true,
		},
		{
			name: "Valid IPv6 endpoint",
			cfg: &kubeadmapi.MasterConfiguration{
				API: kubeadmapi.API{
					AdvertiseAddress: "2001:db8::1",
					BindPort:         4321,
				},
			},
			endpoint: "https://[2001:db8::1]:4321",
			expected: true,
		},
		{
			name: "Invalid IPv4 endpoint",
			cfg: &kubeadmapi.MasterConfiguration{
				API: kubeadmapi.API{
					AdvertiseAddress: "1.2.3.4",
					BindPort:         1234,
				},
			},
			endpoint: "https://[1.2.3.4]:1234",
			expected: false,
		},
		{
			name: "Invalid IPv6 endpoint",
			cfg: &kubeadmapi.MasterConfiguration{
				API: kubeadmapi.API{
					AdvertiseAddress: "2001:db8::1",
					BindPort:         4321,
				},
			},
			endpoint: "https://2001:db8::1:4321",
			expected: false,
		},
		{
			name: "Invalid IPv4 AdvertiseAddress",
			cfg: &kubeadmapi.MasterConfiguration{
				API: kubeadmapi.API{
					AdvertiseAddress: "1.2.34",
					BindPort:         1234,
				},
			},
			endpoint: "https://1.2.3.4:1234",
			expected: false,
		},
		{
			name: "Invalid IPv6 AdvertiseAddress",
			cfg: &kubeadmapi.MasterConfiguration{
				API: kubeadmapi.API{
					AdvertiseAddress: "2001::db8::1",
					BindPort:         4321,
				},
			},
			endpoint: "https://[2001:db8::1]:4321",
			expected: false,
		},
	}
	for _, rt := range tests {
		actual, err := GetMasterEndpoint(rt.cfg)
		if err != nil && rt.expected {
			t.Error(err)
		}
		if actual != rt.endpoint && rt.expected {
			t.Errorf(
				"failed GenerateMasterEndpoint:\n\texpected: %s\n\t actual: %s",
				rt.endpoint,
				(actual),
			)
		}
	}
}
