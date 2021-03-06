package core

func Main() {
	// Load config file and connect to kubernetes cluster
	K8sConnect, YamlConfig := Init()
	// Continuously Watch node changes
	WatchNodes(K8sConnect, YamlConfig)
}