/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/client/clientset/versioned/typed/longhorn/v1beta2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLonghornV1beta2 struct {
	*testing.Fake
}

func (c *FakeLonghornV1beta2) BackingImages(namespace string) v1beta2.BackingImageInterface {
	return &FakeBackingImages{c, namespace}
}

func (c *FakeLonghornV1beta2) BackingImageDataSources(namespace string) v1beta2.BackingImageDataSourceInterface {
	return &FakeBackingImageDataSources{c, namespace}
}

func (c *FakeLonghornV1beta2) BackingImageManagers(namespace string) v1beta2.BackingImageManagerInterface {
	return &FakeBackingImageManagers{c, namespace}
}

func (c *FakeLonghornV1beta2) Backups(namespace string) v1beta2.BackupInterface {
	return &FakeBackups{c, namespace}
}

func (c *FakeLonghornV1beta2) BackupTargets(namespace string) v1beta2.BackupTargetInterface {
	return &FakeBackupTargets{c, namespace}
}

func (c *FakeLonghornV1beta2) BackupVolumes(namespace string) v1beta2.BackupVolumeInterface {
	return &FakeBackupVolumes{c, namespace}
}

func (c *FakeLonghornV1beta2) Engines(namespace string) v1beta2.EngineInterface {
	return &FakeEngines{c, namespace}
}

func (c *FakeLonghornV1beta2) EngineImages(namespace string) v1beta2.EngineImageInterface {
	return &FakeEngineImages{c, namespace}
}

func (c *FakeLonghornV1beta2) InstanceManagers(namespace string) v1beta2.InstanceManagerInterface {
	return &FakeInstanceManagers{c, namespace}
}

func (c *FakeLonghornV1beta2) Nodes(namespace string) v1beta2.NodeInterface {
	return &FakeNodes{c, namespace}
}

func (c *FakeLonghornV1beta2) Orphans(namespace string) v1beta2.OrphanInterface {
	return &FakeOrphans{c, namespace}
}

func (c *FakeLonghornV1beta2) RecurringJobs(namespace string) v1beta2.RecurringJobInterface {
	return &FakeRecurringJobs{c, namespace}
}

func (c *FakeLonghornV1beta2) Replicas(namespace string) v1beta2.ReplicaInterface {
	return &FakeReplicas{c, namespace}
}

func (c *FakeLonghornV1beta2) Settings(namespace string) v1beta2.SettingInterface {
	return &FakeSettings{c, namespace}
}

func (c *FakeLonghornV1beta2) ShareManagers(namespace string) v1beta2.ShareManagerInterface {
	return &FakeShareManagers{c, namespace}
}

func (c *FakeLonghornV1beta2) Snapshots(namespace string) v1beta2.SnapshotInterface {
	return &FakeSnapshots{c, namespace}
}

func (c *FakeLonghornV1beta2) SupportBundles(namespace string) v1beta2.SupportBundleInterface {
	return &FakeSupportBundles{c, namespace}
}

func (c *FakeLonghornV1beta2) SystemBackups(namespace string) v1beta2.SystemBackupInterface {
	return &FakeSystemBackups{c, namespace}
}

func (c *FakeLonghornV1beta2) SystemRestores(namespace string) v1beta2.SystemRestoreInterface {
	return &FakeSystemRestores{c, namespace}
}

func (c *FakeLonghornV1beta2) Volumes(namespace string) v1beta2.VolumeInterface {
	return &FakeVolumes{c, namespace}
}

func (c *FakeLonghornV1beta2) VolumeAttachments(namespace string) v1beta2.VolumeAttachmentInterface {
	return &FakeVolumeAttachments{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeLonghornV1beta2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
