## v4.0.0

* Fix error management
* Fix connection leak to ETCD (`clientv3/concurrency.Session` was never closed)
* More go-ish API
  * `lock.Error` -> `lock.ErrAlreadyLocked`
  * `lock.WithTrylockTimeout` -> `lock.WithTryLockTimeout`
* Better configurability in `lock.NewEtcdLocker()`
  * `lock.WithMaxTryLockTimeout`: Set a max duration a caller can wait for a lock in `WaitAcquire` (default: 2 minutes)
  * `lock.WithCooldownTryLockDuration`: Set a duration between calls to ETCD to acquire a lock (default: 1 second)
* Prevent `clientv3/concurrency.Mutex.Lock(ctx)` to be called with an infinite `context.Background()` which was basically blocking forever.

## v3.4.2

* Simpler dependencies management

## v3.4.1

* Fix race condition in specs

## v3.4.0

* ETCD Client to v3.4.3

## v2.0

* Change ETCD client

## v0.2

```go
func Wait(client *etcd.Client, key string) error
func WaitAcquire(client *etcd.Client, key string, uint64 ttl) (*Lock, erro)
```


## v0.1

```go
func Acquire(client *etcd.Client, key string, uint64 ttl) (*Lock, error)
func (lock *Lock) Release() error
```
