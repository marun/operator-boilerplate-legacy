package operator

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"

	"monis.app/go/openshift/controller"
)

type KeySyncer interface {
	Key() (v1.Object, error)
	controller.Syncer
}

var _ controller.KeySyncer = &wrapper{}

type wrapper struct {
	KeySyncer
}

func (s *wrapper) Key(namespace, name string) (v1.Object, error) {
	return s.KeySyncer.Key()
}
