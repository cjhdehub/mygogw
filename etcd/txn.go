package etcd

import (
	"context"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/google/uuid"
	"go.etcd.io/etcd/clientv3"
	"gogw-server/logger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var txnFactory *txnFactory_t
var txnFactoryOnce sync.Once

type txnFactory_t struct {
	etcd *clientv3.Client
}

type Txn_t struct {
	id         string //仅用来打日志，排错
	txn        clientv3.Txn
	ops        []clientv3.Op
	cmps       []clientv3.Cmp
	created    int64
	keysWithOp []string
}

func GetTxnFactory() *txnFactory_t {
	if txnFactory == nil {
		txnFactoryOnce.Do(func() {

			cfg := clientv3.Config{
				Endpoints: []string{"http://47.110.37.178:2379"},
			}
			etcd, err := clientv3.New(cfg)
			if err != nil {
				logger.Error("new etcd err %v", err)
			}

			txnFactory = &txnFactory_t{
				etcd: etcd,
			}
		})
	}
	return txnFactory
}

func (f *txnFactory_t) NewTxn() *Txn_t {
	id := uuid.New().String()
	logger.Debug("id:%s,tx start", id)
	return &Txn_t{
		id:      id,
		created: time.Now().Unix(),
		txn:     f.etcd.Txn(context.Background()),
	}
}

type tmpMetadata struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              struct {
		NodeName string
	}
}

func (t *Txn_t) VniCreateOp(key string, n int) {
	t.keysWithOp = append(t.keysWithOp, "create "+strconv.Itoa(n)+" "+key)
}

func (t *Txn_t) VniIncOp(key string, n int) {
	t.keysWithOp = append(t.keysWithOp, "add "+strconv.Itoa(n)+" "+key)
}

func (t *Txn_t) VniDecOp(key string, n int) {
	t.keysWithOp = append(t.keysWithOp, "dec "+strconv.Itoa(n)+" "+key)
}

func (t *Txn_t) VniDelOp(key string) {
	t.keysWithOp = append(t.keysWithOp, "del "+key)
}

func (t *Txn_t) IpsecDelOp(key, nodeName string) {
	t.keysWithOp = append(t.keysWithOp, "del "+key+" "+nodeName)
}

func (t *Txn_t) AppendPutOp(key, val string) {
	logger.Debug("id:%s,AppendPutOp key:%s", t.id, key)
	if strings.Contains(key, "ipsec") {
		var tmp tmpMetadata
		err := yaml.Unmarshal([]byte(val), &tmp)
		if err != nil {
			logger.Error("Unmarshal error:", err)
		} else {
			//log.Debugf("%d %d", t.created, tmp.CreationTimestamp.UnixNano())
			if t.created <= tmp.CreationTimestamp.UnixNano() {
				t.keysWithOp = append(t.keysWithOp, "add "+key+" "+tmp.Spec.NodeName)
			}
		}
	}
	op := clientv3.OpPut(key, val)
	t.ops = append(t.ops, op)
}

func (t *Txn_t) AppendPutOpWithLease(key, val string, option clientv3.OpOption) {
	logger.Debug("id:%s,AppendPutOp key:", t.id, key)
	op := clientv3.OpPut(key, val, option)
	t.ops = append(t.ops, op)
}

func (t *Txn_t) GetId() string {
	return t.id
}

func (t *Txn_t) AppendDeleteOp(key string) {
	logger.Debug("id:%s,AppendDeleteOp key:", t.id, key)
	op := clientv3.OpDelete(key)
	t.ops = append(t.ops, op)
}

func (t *Txn_t) AppendDeleteWithPrefixOp(key string) {
	logger.Debug("id:%s,AppendDeleteOp key:", t.id, key)
	op := clientv3.OpDelete(key, clientv3.WithPrefix())
	t.ops = append(t.ops, op)
}

func (t *Txn_t) GetKeys() (keys []string) {
	for _, op := range t.ops {
		keys = append(keys, string(op.KeyBytes()))
	}
	return
}

func (t *Txn_t) AppendCmps(cmp []clientv3.Cmp) {
	t.cmps = append(t.cmps, cmp...)
}
func (t *Txn_t) AppendCmp(cmp clientv3.Cmp) {
	t.cmps = append(t.cmps, cmp)
}
func (t *Txn_t) ClearCmp() {
	t.cmps = t.cmps[0:0]
}

func (t *Txn_t) ClearOps() {
	t.ops = t.ops[0:0]
}

func (t *Txn_t) AppendOps(ops []clientv3.Op) {
	t.ops = append(t.ops, ops...)
}

func (t *Txn_t) GetOps() []clientv3.Op {
	return t.ops
}

func (t *Txn_t) Commit() error {
	txn := t.txn
	if len(t.cmps) > 0 {
		txn = t.txn.If(t.cmps...)
	}

	for _, key := range t.GetKeys() {
		if strings.Contains(key, "vni") {
			keyCompared := false

			for _, cmp := range t.cmps {
				if string(cmp.Key) == key {
					keyCompared = true
					break
				}
			}

			if !keyCompared {
				err := fmt.Errorf("key not compared:%s", key)
				return err
			}
		}
	}
	txnResp, err := txn.Then(t.ops...).Commit()
	if err != nil {
		logger.Error("id:%s, txn commit err:", t.id, err)
	} else if txnResp.Succeeded {
		pc, _, _, _ := runtime.Caller(2)
		logger.Debug("fun:%s, id:%s,Commit ok, keys:%v, vniRefKeys:%v", runtime.FuncForPC(pc).Name(), t.id, t.GetKeys(), t.keysWithOp)
	} else {
		err = fmt.Errorf("txn commit failed")
	}
	return err
}

