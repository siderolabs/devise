package modifiers

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8s struct {
	Secret    *K8sSecret
	ConfigMap *K8sConfigMap
}

type K8sSecret struct {
	ClientSet *kubernetes.Clientset
}

type K8sConfigMap struct {
	ClientSet *kubernetes.Clientset
}

func NewK8s() (*K8s, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	secret := &K8sSecret{ClientSet: clientSet}
	configMap := &K8sConfigMap{ClientSet: clientSet}

	return &K8s{Secret: secret, ConfigMap: configMap}, nil
}

func (k *K8sSecret) Get(s string, c string, ns string) (*string, error) {
	configMap, err := k.ClientSet.Secrets(ns).Get(c)
	if err != nil {
		return nil, err
	}

	value := string(configMap.Data[s])

	return &value, nil
}

func (k *K8sConfigMap) Get(s string, c string, ns string) (*string, error) {
	configMap, err := k.ClientSet.ConfigMaps(ns).Get(c)
	if err != nil {
		return nil, err
	}

	value := string(configMap.Data[s])

	return &value, nil
}
