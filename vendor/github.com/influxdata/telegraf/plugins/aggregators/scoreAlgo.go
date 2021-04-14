package aggregators


func ScoreByTTL(TTL float64) float64 {
	/*延时在[0,40]ms之间，得分为100分；
	延时在(40,60]ms之间，得分为80分；
	延时在(60,80]ms之间，得分为60分；
	延时在(80,~)ms之间，得分为40分；*/
	if TTL >=0 && TTL <=40 {
		return 100
	}else if TTL > 40 && TTL <=60{
		return 80
	}else if TTL > 60 && TTL <=80{
		return 60
	}else {
		return 40
	}
}

func ScoreByLostRate(lostRate float64) float64 {
	/*丢包率为0，得100分；
丢包率为(0,1.7]，得80分；
丢包率为(1.7,5]，得60分；
丢包率为(5,100]，得40分；*/
	if lostRate == 0.0 {
		return 100
	}else if  lostRate > 0.0  && lostRate <=1.7{
		return 80
	}else if  lostRate > 1.7 && lostRate <=5{
		return 60
	}else {
		return 40
	}
}