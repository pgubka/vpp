// Code generated by 'create-test-data.sh' on Thu Dec 13 00:27:38 PST 2018. DO NOT EDIT.

package testdata

func getRawK8sNodeTestData() []string {
	return []string{

		`{
			"addresses": [
				{
					"address": "10.20.0.2",
					"type": "NodeInternalIP"
				},
				{
					"address": "k8s-master",
					"type": "NodeHostName"
				}
			],
			"name": "k8s-master",
			"node_info": {
				"Architecture": "amd64",
				"KubeProxyVersion": "v1.11.4",
				"OperatingSystem": "linux",
				"boot_ID": "8ea33bf5-d11c-4533-bce2-bc0e05ba572e",
				"container_runtime_version": "docker://18.3.0",
				"kernel_version": "4.4.0-21-generic",
				"kubelet_version": "v1.11.4",
				"machine_ID": "91550c3d3d1bca06c11d4f64575584db",
				"os_image": "Ubuntu 16.04 LTS",
				"system_UUID": "1BDE171A-2CB1-496A-B97F-41B0B6587F34"
			},
			"pod_CIDR": "10.0.0.0/24"
		}`,
		`{
			"addresses": [
				{
					"address": "10.20.0.10",
					"type": "NodeInternalIP"
				},
				{
					"address": "k8s-worker1",
					"type": "NodeHostName"
				}
			],
			"name": "k8s-worker1",
			"node_info": {
				"Architecture": "amd64",
				"KubeProxyVersion": "v1.11.4",
				"OperatingSystem": "linux",
				"boot_ID": "dd5daa7a-785c-4554-9621-c26730769c6d",
				"container_runtime_version": "docker://18.3.0",
				"kernel_version": "4.4.0-21-generic",
				"kubelet_version": "v1.11.4",
				"machine_ID": "91550c3d3d1bca06c11d4f64575584db",
				"os_image": "Ubuntu 16.04 LTS",
				"system_UUID": "FD7A2EBD-22B4-4DD7-A7AC-A903E4EA695F"
			},
			"pod_CIDR": "10.0.1.0/24"
		}`,
		`{
			"addresses": [
				{
					"address": "10.20.0.11",
					"type": "NodeInternalIP"
				},
				{
					"address": "k8s-worker2",
					"type": "NodeHostName"
				}
			],
			"name": "k8s-worker2",
			"node_info": {
				"Architecture": "amd64",
				"KubeProxyVersion": "v1.11.4",
				"OperatingSystem": "linux",
				"boot_ID": "1acfcb2c-3f4d-4bfb-9919-0fb8667e03c1",
				"container_runtime_version": "docker://18.3.0",
				"kernel_version": "4.4.0-21-generic",
				"kubelet_version": "v1.11.4",
				"machine_ID": "91550c3d3d1bca06c11d4f64575584db",
				"os_image": "Ubuntu 16.04 LTS",
				"system_UUID": "EB5A6F59-3AAD-4213-B706-CCC34127A096"
			},
			"pod_CIDR": "10.0.2.0/24"
		}`,
	}
}
