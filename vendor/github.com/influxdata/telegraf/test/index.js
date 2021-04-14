var http = require('http');
var random = require("random-js")();
var srcIps=['1.1.1.1','1.1.1.2','1.1.1.3','1.1.1.4','1.1.1.5','1.1.1.6','1.1.1.7','1.1.1.8','1.1.1.9','1.1.1.10'];
var dstIps=['2.1.1.1','3.1.1.2','4.1.1.3','5.1.1.4','6.1.1.5','7.1.1.6','8.1.1.7','9.1.1.8','10.1.1.9','11.1.1.10'];
function RandomNumBoth(Min,Max){
    var Range = Max - Min;
    var Rand = Math.random();
    var num = Min + Math.round(Rand * Range); //四舍五入
    return num;
}

var options = {
    hostname: '127.0.0.1',
    port: 8186,
    path: '/write',
    method: 'POST'
};


function sendMetric(data) {
    console.log(data)
    var req = http.request(options, function (res) {
        res.on('data', function (chunk) {
            console.log('BODY: ' + chunk);
        });
    });
    req.write(data);
    req.end();
}
var failureMap = {}
var time = 0
function collectMetricByPath(index) {
    if(time % 1000 == 0)
        console.log(time+'\n')
    var timestamp = new Date().getTime()*1000000;
    var ttl = RandomNumBoth(1,500);
    ttl = ttl>400?0:ttl; //20%
    var srcIp = srcIps[index],
        dstIP = dstIps[index];

    var tag = 'ping,srcIp='+srcIp+',dstIp='+dstIP

    sendMetric(tag+' ttl='+ttl+' '+timestamp);

    if (++time % 3 == 0){
        var lostRate = RandomNumBoth(0,100),
            score = RandomNumBoth(0,100);
        var randomStr='lostRate='+lostRate+',score='+score+',ttl='+ttl
        console.log('score='+score +'\n');
        console.log('lostRate='+lostRate +'\n');
        if (score < 80) {
            if(failureMap[srcIp+'_'+dstIP] == undefined){
                failureMap[srcIp+'_'+dstIP] = timestamp
                randomStr += ',jumps="{\\"route\\":[{\\"seq\\":1,\\"ttl\\":20,\\"dstIp\\":\\"2.2.2.5\\"}]}"'
                sendMetric(tag + ' ' + randomStr + ' ' + timestamp);
            }
        }else if(failureMap[srcIp+'_'+dstIP] != undefined){
                var failureTimestamp = failureMap[srcIp+'_'+dstIP],
                    failureDuration =  (timestamp - failureTimestamp)/1000000000;
                failureDuration = 'failureDuration='+failureDuration
                failureMap[srcIp+'_'+dstIP] = undefined
                sendMetric(tag + ' ' + failureDuration + ' ' + failureTimestamp);
        }
    }

}
function collectMetric(){
    var i=random.integer(0,9);
        (function (index){
            var j = index
            collectMetricByPath(j)

        })(i)
}
setInterval(collectMetric,1000)


