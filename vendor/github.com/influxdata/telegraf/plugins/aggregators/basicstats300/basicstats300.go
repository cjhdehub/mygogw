package basicstats

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/aggregators"
	"time"
	"sync"
)

type BasicStats300 struct {
	cache map[uint64]aggregate
	locker *sync.Mutex
}

func NewBasicStats300() telegraf.Aggregator {
	mm := &BasicStats300{}
	mm.locker = new(sync.Mutex)
	mm.Reset()
	return mm
}

type aggregate struct {
	fields map[string]basicstats
	name   string
	tags   map[string]string
}

type basicstats struct {
	count float64
	min   float64
	max   float64
	mean  float64
	lostCnt  float64
	lostRate float64
	score float64
	M2    float64 //intermedia value for variance/stdev
}

var sampleConfig = `
  ## General Aggregator Arguments:
  ## The period on which to flush & clear the aggregator.
  period = "30s"
  ## If true, the original metric will be dropped by the
  ## aggregator and will not get sent to the output plugins.
  drop_original = false
`

func (m *BasicStats300) SampleConfig() string {
	return sampleConfig
}

func (m *BasicStats300) Description() string {
	return "Keep the aggregate basicstats of each metric passing through."
}

func (m *BasicStats300) Add(in telegraf.Metric) {
	if in.HasField("jumps") || in.HasField("failureDuration"){
		return
	}
	if in.Name() != "ping" {
		return
	}
	id := in.HashID()
	m.locker.Lock()
	if _, ok := m.cache[id]; !ok {
		// hit an uncached metric, create caches for first time:
		a := aggregate{
			name:   in.Name(),
			tags:   in.Tags(),
			fields: make(map[string]basicstats),
		}
		for k, v := range in.Fields() {
			if fv, ok := convert(v); ok {
				var lostCnt float64
				if (fv == 0) {
					lostCnt = 1
				}else{
					lostCnt = 0
				}

				a.fields[k] = basicstats{
					count: 1,
					min:   fv,
					max:   fv,
					mean:  fv,
					lostCnt: lostCnt,
					lostRate: 0.0,
					M2:    0.0,
				}
			}
		}
		m.cache[id] = a
	} else {
		for k, v := range in.Fields() {
			if fv, ok := convert(v); ok {
				if _, ok := m.cache[id].fields[k]; !ok {
					// hit an uncached field of a cached metric
					m.cache[id].fields[k] = basicstats{
						count: 1,
						min:   fv,
						max:   fv,
						mean:  fv,
						lostCnt: 0,
						lostRate: 0.0,
						M2:    0.0,
					}
					continue
				}

				tmp := m.cache[id].fields[k]

				//https://en.m.wikipedia.org/wiki/Algorithms_for_calculating_variance
				//variable initialization
				x := fv
				mean := tmp.mean
				M2 := tmp.M2
				//counter compute
				n := tmp.count + 1
				tmp.count = n
				//mean compute
				delta := x - mean
				mean = mean + delta/n
				tmp.mean = mean
				if k == "ttl" {
					//variance/stdev compute
					M2 = M2 + delta*(x-mean)
					tmp.M2 = M2
				}
				//max/min compute
				if fv < tmp.min {
					tmp.min = fv
				} else if fv > tmp.max {
					tmp.max = fv
				}

				m.cache[id].fields[k] = tmp
			}
		}
	}
	m.locker.Unlock()
}

func (m *BasicStats300) Push(acc telegraf.Accumulator,period time.Duration) {
	for _, aggregate := range m.cache {
		fields := map[string]interface{}{}
		for k, v := range aggregate.fields {
				if k == "ttl" || k=="lost" || k == "score"{
					fields[k+"_"+period.String()+"_mean"] = v.mean
				}

				//v.count always >=1
				if v.count > 1 && k == "ttl" {
					variance := v.M2 / (v.count - 1)
					fields[k+"_"+period.String()+"_s2"] = variance
				}

		}
		//fields["score_"+period.String()] = score
		aggregate.tags["period"] = period.String()
		acc.AddFields(aggregate.name, fields, aggregate.tags)
	}
}

func (m *BasicStats300) Reset() {
	m.locker.Lock()
	m.cache = make(map[uint64]aggregate)
	m.locker.Unlock()
}

func convert(in interface{}) (float64, bool) {
	switch v := in.(type) {
	case float64:
		return v, true
	case int64:
		return float64(v), true
	default:
		return 0, false
	}
}

func init() {
	aggregators.Add("basicstats300", func() telegraf.Aggregator {
		return NewBasicStats300()
	})
}
