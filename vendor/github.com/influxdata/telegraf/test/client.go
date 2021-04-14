package main

import (
	"fmt"
	"io/ioutil"
	"crypto/x509"
	"crypto/tls"
	"log"
	"io"
)


func main() {
	pem, _ := ioutil.ReadFile("cert2.pem")
	key, _ := ioutil.ReadFile("cert2.key")
	pks, _ := x509.ParsePKCS1PrivateKey(key)

	cert := tls.Certificate{
		Certificate: [][]byte{ pem },
		PrivateKey: pks,
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", "172.31.23.74:8095", &config)
	if err != nil {
		log.Fatalf("dial err", err)
	}
	defer conn.Close()
	log.Println("connected")

	state := conn.ConnectionState()
	for _, v := range state.PeerCertificates {
		fmt.Println(x509.MarshalPKIXPublicKey(v.PublicKey))
		fmt.Println(v.Subject)
	}
/*
SrcIp:v.Fields()["srcIp"].(string),
					DstIp:v.Fields()["dstIp"].(string),
					Seq:v.Fields()["seq"].(string),
					JumpIp:v.Fields()["jumpIp"].(string),
					Ttl:v.Fields()["ttl"].(string),
 */
	//text := `ping,srcIp=192.168.2.43,dstIp=192.168.49.9 jumps="{\"route\":[{\"seq\":1,\"ttl\":100,\"dstIp\":\"192.168.2.1\" },{\"seq\":2,\"ttl\":0,\"dstIp\":\"10.93.0.1\" },{\"seq\":3,\"ttl\":132,\"dstIp\":\"172.31.0.138\" },{\"seq\":4,\"ttl\":0,\"dstIp\":\"172.31.0.137\" },{\"seq\":5,\"ttl\":1,\"dstIp\":\"172.31.123.37\" },{\"seq\":6,\"ttl\":0,\"dstIp\":\"122.224.220.37\" },{\"seq\":7,\"ttl\":* ,\"dstIp\":* },{\"seq\":8,\"ttl\":* ,\"dstIp\":* },{\"seq\":9,\"ttl\":* ,\"dstIp\":* },{\"seq\":10,\"ttl\":* ,\"dstIp\":* },{\"seq\":11,\"ttl\":* ,\"dstIp\":* },{\"seq\":12,\"ttl\":* ,\"dstIp\":* },{\"seq\":13,\"ttl\":* ,\"dstIp\":* },{\"seq\":14,\"ttl\":* ,\"dstIp\":* },{\"seq\":15,\"ttl\":* ,\"dstIp\":* },{\"seq\":16,\"ttl\":* ,\"dstIp\":* },{\"seq\":17,\"ttl\":* ,\"dstIp\":* },{\"seq\":18,\"ttl\":* ,\"dstIp\":* },{\"seq\":19,\"ttl\":* ,\"dstIp\":* },{\"seq\":20,\"ttl\":* ,\"dstIp\":* },{\"seq\":21,\"ttl\":* ,\"dstIp\":* },{\"seq\":22,\"ttl\":* ,\"dstIp\":* },{\"seq\":23,\"ttl\":* ,\"dstIp\":* },{\"seq\":24,\"ttl\":* ,\"dstIp\":* },{\"seq\":25,\"ttl\":* ,\"dstIp\":* },{\"seq\":26,\"ttl\":* ,\"dstIp\":* },{\"seq\":27,\"ttl\":* ,\"dstIp\":* },{\"seq\":28,\"ttl\":* ,\"dstIp\":* },{\"seq\":29,\"ttl\":* ,\"dstIp\":* },{\"seq\":30,\"ttl\":* ,\"dstIp\":* }]}",ttl=0,lostRate=102,score=40`+"\n"
	//text := `pingNode2City,srcIp=192.168.2.3,dstIp=192.168.49.98 ttl=10`+"\n"
	text :=`pingNode2City,srcIp=192.168.2.199 srcCity="杭州市",dstCity="北京市",isp="电信",avlPathCntRate="3/4",ttl=500,lostRate=98,score=80,threshold=30,alertGroup="group1,group2",onAlert=1`+"\n"
	//text :=`traceImmediate,srcIp=1.1.1.1,dstIp=1.2.2.3 seq=1,jumpIp="1.3.4.4",ttl=55`+"\n"

	n, err := io.WriteString(conn, text)
	log.Println(n)
}
