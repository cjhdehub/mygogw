package printer

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/processors"
	"log"
	"github.com/golang/glog"
	"github.com/influxdata/telegraf/plugins/aggregators"
	"math"
)

type Printer struct {
	mapPath2FailureTime map[uint64]int64
}

var sampleConfig = `
`

func (p *Printer) SampleConfig() string {
	return sampleConfig
}

func (p *Printer) Description() string {
	return "Print all metrics that pass through this filter."
}

/*func (p *Printer) loadAlertConfig() string {

}*/

func (p *Printer) Apply(in ...telegraf.Metric) []telegraf.Metric {
	if p.mapPath2FailureTime == nil {
		log.Println("mapPath2Failure is nil")
		p.mapPath2FailureTime = make(map[uint64]int64)
	}

	convert := func (in interface{}) (float64, bool) {
		switch v := in.(type) {
			case float64:
			return v, true
			case int64:
			return float64(v), true
			default:
			return 0, false
		}
	}
	for _, metric := range in {
		if metric.Name() == "ping" {
			if !metric.HasTag("srcIp") {
				glog.V(10).Info("drop metric, no srcIp")
				return nil
			}
			if !metric.HasTag("dstIp") {
				glog.V(10).Info("drop metric, no dstIp")
				return nil
			}

			if metric.HasTag("period") {
				//聚合处理过的数据
				metric.SetName("aggregators"+metric.Tags()["period"])

				var tags2beRemoved []string
				for key,_ := range metric.Tags(){
					if/* key != "period" &&*/ key !="srcIp" && key !="dstIp" {
						tags2beRemoved = append(tags2beRemoved,key)
					}
				}
				glog.V(10).Info("tags2beRemoved = ",tags2beRemoved)
				for _,value :=range tags2beRemoved{
					metric.RemoveTag(value)
				}

			}else{
				//原始数据


				if metric.HasField("jumps") {
					hashId := metric.HashID()
					if failureTime := p.mapPath2FailureTime[hashId];failureTime!=0 {
						glog.V(10).Info("drop metric, had failed in this path already")
						return nil
					}

					metric.SetName("failure")
					//故障开始
					if !metric.HasField("lostRate") {
						glog.V(10).Info("drop metric, no lostRate")
						return nil
					}else if _, ok := convert(metric.Fields()["lostRate"]); !ok {
						glog.V(10).Info("drop metric, lostRate can't convert to float")
						return nil
					}
					if !metric.HasField("score") {
						glog.V(10).Info("drop metric, no score")
						return nil
					}else if _, ok := convert(metric.Fields()["score"]); !ok {
						glog.V(10).Info("drop metric, score can't convert to float")
						return nil
					}
					if !metric.HasField("ttl") {
						glog.V(10).Info("drop metric, no ttl")
						return nil
					}
					for key,_ := range metric.Fields() {
						if key != "lostRate" && key != "score" && key != "jumps" && key != "ttl" {
							glog.V(10).Info("drop metric, unnecessary field :",key)
							return nil
						}
					}
					metric.AddField("lostSolved","0")
					p.mapPath2FailureTime[hashId] = metric.Time().UnixNano()

				}else if metric.HasField("failureDuration"){
					hashId := metric.HashID()
					if failureTime := p.mapPath2FailureTime[hashId];failureTime==0 {
						glog.V(10).Info("drop metric, had not failed in this path")
						return nil
					}else if failureTime != metric.Time().UnixNano(){
						glog.V(10).Info("drop metric, metric timestamp not equal failureTime")
						return nil
					}

					metric.SetName("failure")

					//故障结束
					for key,_ := range metric.Fields() {
						if key != "failureDuration" {
							glog.V(10).Info("drop metric, unnecessary field :",key)
							return nil
						}
					}
					metric.AddField("lostSolved","1")
					p.mapPath2FailureTime[hashId] = 0

				}else{
					if !metric.HasField("ttl") {
						glog.V(10).Info("drop metric, no ttl")
						return nil
					}
					if !metric.HasField("lost") {
						glog.V(10).Info("drop metric, no lost")
						return nil
					}

					if !metric.HasField("mdev") {
						glog.V(10).Info("drop metric, no mdev")
						return nil
					}
					ttl := metric.Fields()["ttl"]
					ttlVal := 0.0
					switch v := ttl.(type) {
					case float64:
						ttlVal = v
					case int64:
						ttlVal = float64(v)
					}

					lost := metric.Fields()["lost"]
					lostVal := 0.0
					switch v := lost.(type) {
					case float64:
						lostVal = v
					case int64:
						lostVal = float64(v)
					}

					lostRateScore := aggregators.ScoreByLostRate(lostVal)
					ttlScore := aggregators.ScoreByTTL(ttlVal)
					score := math.Min(lostRateScore,ttlScore)

					metric.AddField("score",score)

					if _, ok := convert(ttl); ok {

						metric.SetName("ping")
						for key,_ := range metric.Fields() {
							if key != "ttl" && key != "lost" && key != "score" && key != "mdev" {
								glog.V(10).Info("drop metric, unnecessary field :",key)
								return nil
							}
						}

					}else{
						//不会运行到这里,丢弃,除非恶意数据
						glog.V(10).Info("drop metric,ttl can't convert to float")
						return nil
					}

				}
			}
		} else if metric.Name() == "pingNode2City"{
			if !metric.HasTag("srcIp") {
				glog.V(0).Info("drop metric, no srcIp")
				return nil
			}
			/*if !metric.HasTag("dstIp") {
				glog.V(10).Info("drop metric, no dstIp")
				return nil
			}*/
			if !metric.HasTag("dstCity") {
				glog.V(0).Info("drop metric, no dstCity",metric)
				return nil
			}

			if !metric.HasField("srcCity") {
				glog.V(0).Info("drop metric, no srcCity")
				return nil
			}

			if !metric.HasField("avlPathCntRate") {
				glog.V(0).Info("drop metric, no avlPathCntRate")
				return nil
			}
			if !metric.HasField("ttl") {
				glog.V(0).Info("drop metric, no ttl")
				return nil
			}
			if !metric.HasField("lostRate") {
				glog.V(0).Info("drop metric, no lostRate")
				return nil
			}
			if !metric.HasField("score") {
				glog.V(0).Info("drop metric, no score")
				return nil
			}
			if !metric.HasField("threshold") {
				glog.V(0).Info("drop metric, no threshold")
				return nil
			}
			if !metric.HasField("onAlert") {
				glog.V(0).Info("drop metric, no onAlert")
				return nil
			}
			if !metric.HasField("alertGroup") {
				glog.V(0).Info("drop metric, no onAlert")
				return nil
			}
			if !metric.HasField("isp") {
				glog.V(0).Info("drop metric, no isp")
				return nil
			}
			if !metric.HasField("idc") {
				glog.V(0).Info("drop metric, no idc")
				return nil
			}
			for key,_ := range metric.Fields() {
				if  key != "onAlert" &&
				   	key != "avlPathCntRate" &&
				   	key != "ttl" &&
					key != "lostRate" &&
					key != "score" &&
					key != "alertGroup" &&
					key != "srcCity" &&
					key != "idc" &&
					key != "isp" &&
					key != "threshold"{
					glog.V(0).Info("drop metric, unnecessary field :",key)
					return nil
				}
			}

			if metric.Fields()["onAlert"] == "1" {
				metric.AddField("alert",metric.Tags()["alertGroup"])
			}
		}else  if metric.Name() == "tmpMtr" {
			if !metric.HasTag("srcIp") {
				glog.V(0).Info("drop metric, no srcIp")
				return nil
			}
			if !metric.HasTag("dstIp") {
				glog.V(0).Info("drop metric, no dstIp")
				return nil
			}

			if !metric.HasTag("id") {
				glog.V(0).Info("drop metric, no id")
				return nil
			}
			if !metric.HasField("record") {
				glog.V(0).Info("drop metric, no record")
				return nil
			}

			for key,_ := range metric.Fields() {
				if  key != "record"{
					glog.V(0).Info("drop metric, unnecessary field :",key)
					return nil
				}
			}
		}else if metric.Name() == "monitorMtr"{
			if !metric.HasTag("srcIp") {
				glog.V(0).Info("drop metric, no srcIp")
				return nil
			}
			if !metric.HasTag("dstIp") {
				glog.V(0).Info("drop metric, no dstIp")
				return nil
			}
			if !metric.HasTag("srcCity") {
				glog.V(0).Info("drop metric, no srcCity",metric)
				return nil
			}
			if !metric.HasTag("dstCity") {
				glog.V(0).Info("drop metric, no dstCity")
				return nil
			}
			if !metric.HasTag("id") {
				glog.V(0).Info("drop metric, no id")
				return nil
			}
			if !metric.HasField("record") {
				glog.V(0).Info("drop metric, no record")
				return nil
			}

			for key,_ := range metric.Fields() {
				if  key != "record" {
					glog.V(0).Info("drop metric, unnecessary field :",key)
					return nil
				}
			}
		}else if metric.Name() == "tmpPing" {
			if !metric.HasTag("srcIp") {
				glog.V(0).Info("drop metric, no srcIp")
				return nil
			}
			if !metric.HasTag("dstIp") {
				glog.V(0).Info("drop metric, no dstIp")
				return nil
			}
			if !metric.HasTag("id") {
				glog.V(0).Info("drop metric, no id")
				return nil
			}

			var summaryPing bool
			if metric.HasField("mdev") &&  metric.HasField("min") && metric.HasField("avg") && metric.HasField("max") && metric.HasField("transmitted") && metric.HasField("received") && metric.HasField("loss") {
				summaryPing = true
			}else if (!metric.HasField("transmitted") && !metric.HasField("received") && !metric.HasField("loss")){
				summaryPing = false
			}else{
				glog.V(0).Info("drop metric, neither summary ping nor normal ping")
				return nil
			}
			//glog.V(0).Info(metric)

			for key,_ := range metric.Fields() {
				if !summaryPing && key != "ttl"&& key != "dura"&& key != "icmp_seq" {
					glog.V(0).Info("drop metric, unnecessary field :",key)
					return nil
				}else if summaryPing &&  key != "max" && key != "mdev" && key != "min"&& key != "avg" &&  key != "received" && key != "ttl"&& key != "dura"&& key != "icmp_seq"&& key != "transmitted"&& key != "loss" {
					glog.V(0).Info("drop metric, unnecessary field :",key)
					return nil
				}
			}
		}else  if metric.Name() == "monitorPing"{
			if !metric.HasTag("srcIp") {
				glog.V(0).Info("drop metric, no srcIp")
				return nil
			}
			if !metric.HasTag("dstIp") {
				glog.V(0).Info("drop metric, no dstIp")
				return nil
			}
			if !metric.HasTag("id") {
				glog.V(0).Info("drop metric, no id")
				return nil
			}

			var summaryPing bool
			if metric.HasField("mdev") &&  metric.HasField("min") && metric.HasField("avg") && metric.HasField("max") && metric.HasField("transmitted") && metric.HasField("received") && metric.HasField("loss") {
				summaryPing = true
			}else if (!metric.HasField("transmitted") && !metric.HasField("received") && !metric.HasField("loss")){
				summaryPing = false
			}else{
				glog.V(0).Info("drop metric, neither summary ping nor normal ping")
				return nil
			}
			//glog.V(0).Info(metric)
			if !metric.HasTag("idc") {
				glog.V(0).Info("drop metric, no idc")
				return nil
			}
			if !metric.HasTag("isp") {
				glog.V(0).Info("drop metric, no isp")
				return nil
			}

			for key,_ := range metric.Fields() {
				if !summaryPing && key != "ttl"&& key != "dura"&& key != "icmp_seq"{
					glog.V(0).Info("drop metric, unnecessary field :",key)
					return nil
				}else if summaryPing && key != "score" && key != "max" && key != "mdev" && key != "min"&& key != "avg" &&  key != "received" && key != "ttl"&& key != "dura"&& key != "icmp_seq"&& key != "transmitted"&& key != "loss" {
					glog.V(0).Info("drop metric, unnecessary field :",key)
					return nil
				}
			}
		}
	}

	return in
}

func init() {
	processors.Add("printer", func() telegraf.Processor {
		return &Printer{}
	})
}
