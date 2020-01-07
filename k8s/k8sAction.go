package k8s

import (
	"fmt"
	"strconv"

	"tiodeploy/k8s/data"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	tiomodel "github.com/tio-serverless/model/database/model"
)

// CreateNewDeploy 返回可用的Kubernetes Deploy对象
func CreateNewDeploy(config *data.B, k MyK8s, s tiomodel.Server) (string, error) {
	d := deploy{
		Image: s.Image,
		Name:  s.Name,
	}
	var meta data.MyDeploy

	_, err := toml.Decode(s.Raw, &meta)
	if err != nil {
		return "", err
	}

	env := make(map[string]string)

	env["MY_POD_PORT"] = "80"
	env["MY_SERVICE_NAME"] = s.Name
	env["CONSUL_ADDRESS"] = config.K.Consul
	env["MY_SERVICE_TYPE"] = strconv.Itoa(s.Stype)
	env["MY_SERVICE_URL"] = meta.Meta.Url
	env["DEBUG"] = "debug"

	d.Env = env

	isExist, err := k.IsHasDeploy(d.Name)
	if err != nil {
		logrus.Errorf("Check Exist Deployment Error. %s", err)
		isExist = false
	}

	if isExist {
		logrus.Debugf("Replace Exist Deployment[%s]", d.Name)
		return d.Name, k.ReplaceDeploy(d)
	}

	return k.NewDeploy(d)
}

// ScalaInstances 调整实例数, num为实例具体个数
func ScalaInstances(k MyK8s, id string, num int) error {
	return k.Scala(id, num)
}

// DeleteDeploy 删除Deployment, id deploy名称
func DeleteDeploy(k MyK8s, id string) error {
	return k.Delete(id)
}

// GetDeploymentLog 获取Deployment日志
func GetDeploymentLog(k MyK8s, name string, logs chan string) error {
	return k.GetLog(deploy{
		Name: name,
	}, logs)
}

// UpdateDeployment 更新Deployment环境变量
func UpdateDeployment(k MyK8s, name string, env map[string]string) error {
	return k.Update(deploy{
		Name: name,
		Env:  env,
	})
}

// GetPodEndpoint  获取指定Pod IP地址
func GetPodEndpoint(k MyK8s, name string) (string, error) {
	pod, err := k.GetPodInfo(name)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:80", pod.Status.PodIP), nil
}

// GetDeployInstanceNum 获取Deployment实例数量
func GetDeployInstanceNum(k MyK8s, name string) (int, error) {
	d, err := k.GetDeploymentInfo(name)
	if err != nil {
		return 0, err
	}

	return int(d.Status.Replicas), nil
}
