/*
 * Copyright (C) 2018 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy ofthe License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specificlanguage governing permissions and
 * limitations under the License.
 *
 */

package rbac

import (
	"context"
	"runtime"

	"github.com/casbin/casbin/persist"
	etcd "go.etcd.io/etcd/client/v2"
)

// EtcdWatcher listens for etcd events
type EtcdWatcher struct {
	kapi     etcd.KeysAPI
	running  bool
	callback func(string)
	cancel   context.CancelFunc
}

// finalizer is the destructor for EtcdWatcher.
func finalizer(w *EtcdWatcher) {
	w.running = false
	w.cancel()
}

// NewEtcdWatcher returns new etcd change watcher
func NewEtcdWatcher(kapi etcd.KeysAPI, parent context.Context) persist.Watcher {
	ctx, cancel := context.WithCancel(parent)

	w := &EtcdWatcher{
		kapi:    kapi,
		running: true,
		cancel:  cancel,
	}

	// Call the destructor when the object is released.
	runtime.SetFinalizer(w, finalizer)

	go w.startWatch(ctx)

	return w
}

// SetUpdateCallback sets the callback function that the watcher will call
// when the policy in DB has been changed by other instances.
// A classic callback is Enforcer.LoadPolicy().
func (w *EtcdWatcher) SetUpdateCallback(callback func(string)) error {
	w.callback = callback
	return nil
}

// Update calls the update callback of other instances to synchronize their policy.
// It is usually called after changing the policy in DB, like Enforcer.SavePolicy(),
// Enforcer.AddPolicy(), Enforcer.RemovePolicy(), etc.
func (w *EtcdWatcher) Update() error {
	return nil
}

func (w *EtcdWatcher) Close() {
	w.cancel()
}

// startWatch is a goroutine that watches the policy change.
func (w *EtcdWatcher) startWatch(ctx context.Context) error {
	watcher := w.kapi.Watcher(etcdPolicyKey, &etcd.WatcherOptions{Recursive: false})
	for {
		if !w.running {
			return nil
		}

		res, err := watcher.Next(ctx)
		if err != nil {
			return err
		}

		if res.Action == "set" || res.Action == "update" {
			if w.callback != nil {
				w.callback(res.Node.Value)
			}
		}
	}
}
