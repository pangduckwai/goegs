package common

func Rand2(rnd uint64) int {
	if rnd < 9223372036854775807 {
		return 0
	} else {
		return 1
	}
}

func Rand3(rnd uint64) int {
	if rnd < 6148914691236517205 {
		return 0
	} else if rnd < 12297829382473034410 {
		return 1
	} else {
		return 2
	}
}

func Rand10(rnd uint64) int {
	if rnd < 1844674407370955161 {
		return 0
	} else if rnd < 3689348814741910322 {
		return 1
	} else if rnd < 5534023222112865484 {
		return 2
	} else if rnd < 7378697629483820645 {
		return 3
	} else if rnd < 9223372036854775807 {
		return 4
	} else if rnd < 11068046444225730968 {
		return 5
	} else if rnd < 12912720851596686130 {
		return 6
	} else if rnd < 14757395258967641291 {
		return 7
	} else if rnd < 16602069666338596453 {
		return 8
	} else {
		return 9
	}
}

func RandN(n int, rnd uint64) int {
	switch n {
	case 2:
		if rnd < 9223372036854775807 { // diff:1; 2÷1=2
			return 0
		} else { // < 18446744073709551615; calc:18446744073709551614
			return 1
		} // num:2
	case 3:
		if rnd < 6148914691236517205 { // diff:0; 3÷0=0
			return 0
		} else if rnd < 12297829382473034410 { // diff:0
			return 1
		} else { // < 18446744073709551615; calc:18446744073709551615
			return 2
		} // num:3
	case 4:
		if rnd < 4611686018427387903 { // diff:3; 4÷3=1.3333333333333333
			return 0
		} else if rnd < 9223372036854775807 { // diff:2
			return 1
		} else if rnd < 13835058055282163711 { // diff:1
			return 2
		} else { // < 18446744073709551615; calc:18446744073709551612
			return 3
		} // num:4
	case 5:
		if rnd < 3689348814741910323 { // diff:0; 5÷0=0
			return 0
		} else if rnd < 7378697629483820646 { // diff:0
			return 1
		} else if rnd < 11068046444225730969 { // diff:0
			return 2
		} else if rnd < 14757395258967641292 { // diff:0
			return 3
		} else { // < 18446744073709551615; calc:18446744073709551615
			return 4
		} // num:5
	case 6:
		if rnd < 3074457345618258602 { // diff:3; 6÷3=2
			return 0
		} else if rnd < 6148914691236517204 { // diff:3
			return 1
		} else if rnd < 9223372036854775807 { // diff:2
			return 2
		} else if rnd < 12297829382473034409 { // diff:2
			return 3
		} else if rnd < 15372286728091293012 { // diff:1
			return 4
		} else { // < 18446744073709551615; calc:18446744073709551612
			return 5
		} // num:6
	case 7:
		if rnd < 2635249153387078802 { // diff:1; 7÷1=7
			return 0
		} else if rnd < 5270498306774157604 { // diff:1
			return 1
		} else if rnd < 7905747460161236406 { // diff:1
			return 2
		} else if rnd < 10540996613548315208 { // diff:1
			return 3
		} else if rnd < 13176245766935394010 { // diff:1
			return 4
		} else if rnd < 15811494920322472812 { // diff:1
			return 5
		} else { // < 18446744073709551615; calc:18446744073709551614
			return 6
		} // num:7
	case 8:
		if rnd < 2305843009213693951 { // diff:7; 8÷7=1.1428571428571428
			return 0
		} else if rnd < 4611686018427387903 { // diff:6
			return 1
		} else if rnd < 6917529027641081855 { // diff:5
			return 2
		} else if rnd < 9223372036854775807 { // diff:4
			return 3
		} else if rnd < 11529215046068469759 { // diff:3
			return 4
		} else if rnd < 13835058055282163711 { // diff:2
			return 5
		} else if rnd < 16140901064495857663 { // diff:1
			return 6
		} else { // < 18446744073709551615; calc:18446744073709551608
			return 7
		} // num:8
	case 9:
		if rnd < 2049638230412172401 { // diff:6; 9÷6=1.5
			return 0
		} else if rnd < 4099276460824344803 { // diff:5
			return 1
		} else if rnd < 6148914691236517204 { // diff:5
			return 2
		} else if rnd < 8198552921648689606 { // diff:4
			return 3
		} else if rnd < 10248191152060862008 { // diff:3
			return 4
		} else if rnd < 12297829382473034409 { // diff:3
			return 5
		} else if rnd < 14347467612885206811 { // diff:2
			return 6
		} else if rnd < 16397105843297379213 { // diff:1
			return 7
		} else { // < 18446744073709551615; calc:18446744073709551609
			return 8
		} // num:9
	case 10:
		if rnd < 1844674407370955161 { // diff:5; 10÷5=2
			return 0
		} else if rnd < 3689348814741910322 { // diff:5
			return 1
		} else if rnd < 5534023222112865484 { // diff:4
			return 2
		} else if rnd < 7378697629483820645 { // diff:4
			return 3
		} else if rnd < 9223372036854775807 { // diff:3
			return 4
		} else if rnd < 11068046444225730968 { // diff:3
			return 5
		} else if rnd < 12912720851596686130 { // diff:2
			return 6
		} else if rnd < 14757395258967641291 { // diff:2
			return 7
		} else if rnd < 16602069666338596453 { // diff:1
			return 8
		} else { // < 18446744073709551615; calc:18446744073709551610
			return 9
		} // num:10
	case 11:
		if rnd < 1676976733973595601 { // diff:4; 11÷4=2.75
			return 0
		} else if rnd < 3353953467947191202 { // diff:4
			return 1
		} else if rnd < 5030930201920786804 { // diff:3
			return 2
		} else if rnd < 6707906935894382405 { // diff:3
			return 3
		} else if rnd < 8384883669867978006 { // diff:3
			return 4
		} else if rnd < 10061860403841573608 { // diff:2
			return 5
		} else if rnd < 11738837137815169209 { // diff:2
			return 6
		} else if rnd < 13415813871788764810 { // diff:2
			return 7
		} else if rnd < 15092790605762360412 { // diff:1
			return 8
		} else if rnd < 16769767339735956013 { // diff:1
			return 9
		} else { // < 18446744073709551615; calc:18446744073709551611
			return 10
		} // num:11
	case 12:
		if rnd < 1537228672809129301 { // diff:3; 12÷3=4
			return 0
		} else if rnd < 3074457345618258602 { // diff:3
			return 1
		} else if rnd < 4611686018427387903 { // diff:3
			return 2
		} else if rnd < 6148914691236517204 { // diff:3
			return 3
		} else if rnd < 7686143364045646506 { // diff:2
			return 4
		} else if rnd < 9223372036854775807 { // diff:2
			return 5
		} else if rnd < 10760600709663905108 { // diff:2
			return 6
		} else if rnd < 12297829382473034409 { // diff:2
			return 7
		} else if rnd < 13835058055282163711 { // diff:1
			return 8
		} else if rnd < 15372286728091293012 { // diff:1
			return 9
		} else if rnd < 16909515400900422313 { // diff:1
			return 10
		} else { // < 18446744073709551615; calc:18446744073709551612
			return 11
		} // num:12
	case 13:
		if rnd < 1418980313362273201 { // diff:2; 13÷2=6.5
			return 0
		} else if rnd < 2837960626724546402 { // diff:2
			return 1
		} else if rnd < 4256940940086819603 { // diff:2
			return 2
		} else if rnd < 5675921253449092804 { // diff:2
			return 3
		} else if rnd < 7094901566811366005 { // diff:2
			return 4
		} else if rnd < 8513881880173639206 { // diff:2
			return 5
		} else if rnd < 9932862193535912408 { // diff:1
			return 6
		} else if rnd < 11351842506898185609 { // diff:1
			return 7
		} else if rnd < 12770822820260458810 { // diff:1
			return 8
		} else if rnd < 14189803133622732011 { // diff:1
			return 9
		} else if rnd < 15608783446985005212 { // diff:1
			return 10
		} else if rnd < 17027763760347278413 { // diff:1
			return 11
		} else { // < 18446744073709551615; calc:18446744073709551613
			return 12
		} // num:13
	case 14:
		if rnd < 1317624576693539401 { // diff:1; 14÷1=14
			return 0
		} else if rnd < 2635249153387078802 { // diff:1
			return 1
		} else if rnd < 3952873730080618203 { // diff:1
			return 2
		} else if rnd < 5270498306774157604 { // diff:1
			return 3
		} else if rnd < 6588122883467697005 { // diff:1
			return 4
		} else if rnd < 7905747460161236406 { // diff:1
			return 5
		} else if rnd < 9223372036854775807 { // diff:1
			return 6
		} else if rnd < 10540996613548315208 { // diff:1
			return 7
		} else if rnd < 11858621190241854609 { // diff:1
			return 8
		} else if rnd < 13176245766935394010 { // diff:1
			return 9
		} else if rnd < 14493870343628933411 { // diff:1
			return 10
		} else if rnd < 15811494920322472812 { // diff:1
			return 11
		} else if rnd < 17129119497016012213 { // diff:1
			return 12
		} else { // < 18446744073709551615; calc:18446744073709551614
			return 13
		} // num:14
	case 15:
		if rnd < 1229782938247303441 { // diff:0; 15÷0=0
			return 0
		} else if rnd < 2459565876494606882 { // diff:0
			return 1
		} else if rnd < 3689348814741910323 { // diff:0
			return 2
		} else if rnd < 4919131752989213764 { // diff:0
			return 3
		} else if rnd < 6148914691236517205 { // diff:0
			return 4
		} else if rnd < 7378697629483820646 { // diff:0
			return 5
		} else if rnd < 8608480567731124087 { // diff:0
			return 6
		} else if rnd < 9838263505978427528 { // diff:0
			return 7
		} else if rnd < 11068046444225730969 { // diff:0
			return 8
		} else if rnd < 12297829382473034410 { // diff:0
			return 9
		} else if rnd < 13527612320720337851 { // diff:0
			return 10
		} else if rnd < 14757395258967641292 { // diff:0
			return 11
		} else if rnd < 15987178197214944733 { // diff:0
			return 12
		} else if rnd < 17216961135462248174 { // diff:0
			return 13
		} else { // < 18446744073709551615; calc:18446744073709551615
			return 14
		} // num:15
	case 16:
		if rnd < 1152921504606846975 { // diff:15; 16÷15=1.0666666666666667
			return 0
		} else if rnd < 2305843009213693951 { // diff:14
			return 1
		} else if rnd < 3458764513820540927 { // diff:13
			return 2
		} else if rnd < 4611686018427387903 { // diff:12
			return 3
		} else if rnd < 5764607523034234879 { // diff:11
			return 4
		} else if rnd < 6917529027641081855 { // diff:10
			return 5
		} else if rnd < 8070450532247928831 { // diff:9
			return 6
		} else if rnd < 9223372036854775807 { // diff:8
			return 7
		} else if rnd < 10376293541461622783 { // diff:7
			return 8
		} else if rnd < 11529215046068469759 { // diff:6
			return 9
		} else if rnd < 12682136550675316735 { // diff:5
			return 10
		} else if rnd < 13835058055282163711 { // diff:4
			return 11
		} else if rnd < 14987979559889010687 { // diff:3
			return 12
		} else if rnd < 16140901064495857663 { // diff:2
			return 13
		} else if rnd < 17293822569102704639 { // diff:1
			return 14
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 15
		} // num:16
	case 17:
		if rnd < 1085102592571150095 { // diff:0; 17÷0=0
			return 0
		} else if rnd < 2170205185142300190 { // diff:0
			return 1
		} else if rnd < 3255307777713450285 { // diff:0
			return 2
		} else if rnd < 4340410370284600380 { // diff:0
			return 3
		} else if rnd < 5425512962855750475 { // diff:0
			return 4
		} else if rnd < 6510615555426900570 { // diff:0
			return 5
		} else if rnd < 7595718147998050665 { // diff:0
			return 6
		} else if rnd < 8680820740569200760 { // diff:0
			return 7
		} else if rnd < 9765923333140350855 { // diff:0
			return 8
		} else if rnd < 10851025925711500950 { // diff:0
			return 9
		} else if rnd < 11936128518282651045 { // diff:0
			return 10
		} else if rnd < 13021231110853801140 { // diff:0
			return 11
		} else if rnd < 14106333703424951235 { // diff:0
			return 12
		} else if rnd < 15191436295996101330 { // diff:0
			return 13
		} else if rnd < 16276538888567251425 { // diff:0
			return 14
		} else if rnd < 17361641481138401520 { // diff:0
			return 15
		} else { // < 18446744073709551615; calc:18446744073709551615
			return 16
		} // num:17
	case 18:
		if rnd < 1024819115206086200 { // diff:15; 18÷15=1.2
			return 0
		} else if rnd < 2049638230412172401 { // diff:14
			return 1
		} else if rnd < 3074457345618258602 { // diff:13
			return 2
		} else if rnd < 4099276460824344803 { // diff:12
			return 3
		} else if rnd < 5124095576030431004 { // diff:11
			return 4
		} else if rnd < 6148914691236517204 { // diff:11
			return 5
		} else if rnd < 7173733806442603405 { // diff:10
			return 6
		} else if rnd < 8198552921648689606 { // diff:9
			return 7
		} else if rnd < 9223372036854775807 { // diff:8
			return 8
		} else if rnd < 10248191152060862008 { // diff:7
			return 9
		} else if rnd < 11273010267266948209 { // diff:6
			return 10
		} else if rnd < 12297829382473034410 { // diff:5
			return 11
		} else if rnd < 13322648497679120610 { // diff:5
			return 12
		} else if rnd < 14347467612885206811 { // diff:4
			return 13
		} else if rnd < 15372286728091293012 { // diff:3
			return 14
		} else if rnd < 16397105843297379213 { // diff:2
			return 15
		} else if rnd < 17421924958503465414 { // diff:1
			return 16
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 17
		} // num:18
	case 19:
		if rnd < 970881267037344821 { // diff:16; 19÷16=1.1875
			return 0
		} else if rnd < 1941762534074689643 { // diff:15
			return 1
		} else if rnd < 2912643801112034465 { // diff:14
			return 2
		} else if rnd < 3883525068149379287 { // diff:13
			return 3
		} else if rnd < 4854406335186724109 { // diff:12
			return 4
		} else if rnd < 5825287602224068931 { // diff:11
			return 5
		} else if rnd < 6796168869261413752 { // diff:11
			return 6
		} else if rnd < 7767050136298758574 { // diff:10
			return 7
		} else if rnd < 8737931403336103396 { // diff:9
			return 8
		} else if rnd < 9708812670373448218 { // diff:8
			return 9
		} else if rnd < 10679693937410793040 { // diff:7
			return 10
		} else if rnd < 11650575204448137862 { // diff:6
			return 11
		} else if rnd < 12621456471485482683 { // diff:6
			return 12
		} else if rnd < 13592337738522827505 { // diff:5
			return 13
		} else if rnd < 14563219005560172327 { // diff:4
			return 14
		} else if rnd < 15534100272597517149 { // diff:3
			return 15
		} else if rnd < 16504981539634861971 { // diff:2
			return 16
		} else if rnd < 17475862806672206793 { // diff:1
			return 17
		} else { // < 18446744073709551615; calc:18446744073709551599
			return 18
		} // num:19
	case 20:
		if rnd < 922337203685477580 { // diff:15; 20÷15=1.3333333333333333
			return 0
		} else if rnd < 1844674407370955161 { // diff:14
			return 1
		} else if rnd < 2767011611056432742 { // diff:13
			return 2
		} else if rnd < 3689348814741910322 { // diff:13
			return 3
		} else if rnd < 4611686018427387903 { // diff:12
			return 4
		} else if rnd < 5534023222112865484 { // diff:11
			return 5
		} else if rnd < 6456360425798343065 { // diff:10
			return 6
		} else if rnd < 7378697629483820646 { // diff:9
			return 7
		} else if rnd < 8301034833169298226 { // diff:9
			return 8
		} else if rnd < 9223372036854775807 { // diff:8
			return 9
		} else if rnd < 10145709240540253388 { // diff:7
			return 10
		} else if rnd < 11068046444225730968 { // diff:7
			return 11
		} else if rnd < 11990383647911208549 { // diff:6
			return 12
		} else if rnd < 12912720851596686130 { // diff:5
			return 13
		} else if rnd < 13835058055282163711 { // diff:4
			return 14
		} else if rnd < 14757395258967641291 { // diff:4
			return 15
		} else if rnd < 15679732462653118872 { // diff:3
			return 16
		} else if rnd < 16602069666338596453 { // diff:2
			return 17
		} else if rnd < 17524406870024074034 { // diff:1
			return 18
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 19
		} // num:20
	case 21:
		if rnd < 878416384462359600 { // diff:15; 21÷15=1.4
			return 0
		} else if rnd < 1756832768924719201 { // diff:14
			return 1
		} else if rnd < 2635249153387078802 { // diff:13
			return 2
		} else if rnd < 3513665537849438402 { // diff:13
			return 3
		} else if rnd < 4392081922311798003 { // diff:12
			return 4
		} else if rnd < 5270498306774157604 { // diff:11
			return 5
		} else if rnd < 6148914691236517204 { // diff:11
			return 6
		} else if rnd < 7027331075698876805 { // diff:10
			return 7
		} else if rnd < 7905747460161236406 { // diff:9
			return 8
		} else if rnd < 8784163844623596007 { // diff:8
			return 9
		} else if rnd < 9662580229085955607 { // diff:8
			return 10
		} else if rnd < 10540996613548315208 { // diff:7
			return 11
		} else if rnd < 11419412998010674809 { // diff:6
			return 12
		} else if rnd < 12297829382473034409 { // diff:6
			return 13
		} else if rnd < 13176245766935394010 { // diff:5
			return 14
		} else if rnd < 14054662151397753611 { // diff:4
			return 15
		} else if rnd < 14933078535860113212 { // diff:3
			return 16
		} else if rnd < 15811494920322472812 { // diff:3
			return 17
		} else if rnd < 16689911304784832413 { // diff:2
			return 18
		} else if rnd < 17568327689247192014 { // diff:1
			return 19
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 20
		} // num:21
	case 22:
		if rnd < 838488366986797800 { // diff:15; 22÷15=1.4666666666666666
			return 0
		} else if rnd < 1676976733973595601 { // diff:14
			return 1
		} else if rnd < 2515465100960393402 { // diff:13
			return 2
		} else if rnd < 3353953467947191202 { // diff:13
			return 3
		} else if rnd < 4192441834933989003 { // diff:12
			return 4
		} else if rnd < 5030930201920786804 { // diff:11
			return 5
		} else if rnd < 5869418568907584604 { // diff:11
			return 6
		} else if rnd < 6707906935894382405 { // diff:10
			return 7
		} else if rnd < 7546395302881180206 { // diff:9
			return 8
		} else if rnd < 8384883669867978006 { // diff:9
			return 9
		} else if rnd < 9223372036854775807 { // diff:8
			return 10
		} else if rnd < 10061860403841573608 { // diff:7
			return 11
		} else if rnd < 10900348770828371408 { // diff:7
			return 12
		} else if rnd < 11738837137815169209 { // diff:6
			return 13
		} else if rnd < 12577325504801967010 { // diff:5
			return 14
		} else if rnd < 13415813871788764810 { // diff:5
			return 15
		} else if rnd < 14254302238775562611 { // diff:4
			return 16
		} else if rnd < 15092790605762360412 { // diff:3
			return 17
		} else if rnd < 15931278972749158212 { // diff:3
			return 18
		} else if rnd < 16769767339735956013 { // diff:2
			return 19
		} else if rnd < 17608255706722753814 { // diff:1
			return 20
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 21
		} // num:22
	case 23:
		if rnd < 802032351030850070 { // diff:5; 23÷5=4.6
			return 0
		} else if rnd < 1604064702061700140 { // diff:5
			return 1
		} else if rnd < 2406097053092550210 { // diff:5
			return 2
		} else if rnd < 3208129404123400280 { // diff:5
			return 3
		} else if rnd < 4010161755154250351 { // diff:4
			return 4
		} else if rnd < 4812194106185100421 { // diff:4
			return 5
		} else if rnd < 5614226457215950491 { // diff:4
			return 6
		} else if rnd < 6416258808246800561 { // diff:4
			return 7
		} else if rnd < 7218291159277650631 { // diff:4
			return 8
		} else if rnd < 8020323510308500702 { // diff:3
			return 9
		} else if rnd < 8822355861339350772 { // diff:3
			return 10
		} else if rnd < 9624388212370200842 { // diff:3
			return 11
		} else if rnd < 10426420563401050912 { // diff:3
			return 12
		} else if rnd < 11228452914431900983 { // diff:2
			return 13
		} else if rnd < 12030485265462751053 { // diff:2
			return 14
		} else if rnd < 12832517616493601123 { // diff:2
			return 15
		} else if rnd < 13634549967524451193 { // diff:2
			return 16
		} else if rnd < 14436582318555301263 { // diff:2
			return 17
		} else if rnd < 15238614669586151334 { // diff:1
			return 18
		} else if rnd < 16040647020617001404 { // diff:1
			return 19
		} else if rnd < 16842679371647851474 { // diff:1
			return 20
		} else if rnd < 17644711722678701544 { // diff:1
			return 21
		} else { // < 18446744073709551615; calc:18446744073709551610
			return 22
		} // num:23
	case 24:
		if rnd < 768614336404564650 { // diff:15; 24÷15=1.6
			return 0
		} else if rnd < 1537228672809129301 { // diff:14
			return 1
		} else if rnd < 2305843009213693951 { // diff:14
			return 2
		} else if rnd < 3074457345618258602 { // diff:13
			return 3
		} else if rnd < 3843071682022823253 { // diff:12
			return 4
		} else if rnd < 4611686018427387903 { // diff:12
			return 5
		} else if rnd < 5380300354831952554 { // diff:11
			return 6
		} else if rnd < 6148914691236517204 { // diff:11
			return 7
		} else if rnd < 6917529027641081855 { // diff:10
			return 8
		} else if rnd < 7686143364045646506 { // diff:9
			return 9
		} else if rnd < 8454757700450211156 { // diff:9
			return 10
		} else if rnd < 9223372036854775807 { // diff:8
			return 11
		} else if rnd < 9991986373259340458 { // diff:7
			return 12
		} else if rnd < 10760600709663905108 { // diff:7
			return 13
		} else if rnd < 11529215046068469759 { // diff:6
			return 14
		} else if rnd < 12297829382473034410 { // diff:5
			return 15
		} else if rnd < 13066443718877599060 { // diff:5
			return 16
		} else if rnd < 13835058055282163711 { // diff:4
			return 17
		} else if rnd < 14603672391686728361 { // diff:4
			return 18
		} else if rnd < 15372286728091293012 { // diff:3
			return 19
		} else if rnd < 16140901064495857663 { // diff:2
			return 20
		} else if rnd < 16909515400900422313 { // diff:2
			return 21
		} else if rnd < 17678129737304986964 { // diff:1
			return 22
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 23
		} // num:24
	case 25:
		if rnd < 737869762948382064 { // diff:15; 25÷15=1.6666666666666667
			return 0
		} else if rnd < 1475739525896764129 { // diff:14
			return 1
		} else if rnd < 2213609288845146193 { // diff:14
			return 2
		} else if rnd < 2951479051793528258 { // diff:13
			return 3
		} else if rnd < 3689348814741910322 { // diff:13
			return 4
		} else if rnd < 4427218577690292387 { // diff:12
			return 5
		} else if rnd < 5165088340638674452 { // diff:11
			return 6
		} else if rnd < 5902958103587056516 { // diff:11
			return 7
		} else if rnd < 6640827866535438581 { // diff:10
			return 8
		} else if rnd < 7378697629483820645 { // diff:10
			return 9
		} else if rnd < 8116567392432202710 { // diff:9
			return 10
		} else if rnd < 8854437155380584775 { // diff:8
			return 11
		} else if rnd < 9592306918328966839 { // diff:8
			return 12
		} else if rnd < 10330176681277348904 { // diff:7
			return 13
		} else if rnd < 11068046444225730969 { // diff:6
			return 14
		} else if rnd < 11805916207174113033 { // diff:6
			return 15
		} else if rnd < 12543785970122495098 { // diff:5
			return 16
		} else if rnd < 13281655733070877162 { // diff:5
			return 17
		} else if rnd < 14019525496019259227 { // diff:4
			return 18
		} else if rnd < 14757395258967641291 { // diff:4
			return 19
		} else if rnd < 15495265021916023356 { // diff:3
			return 20
		} else if rnd < 16233134784864405421 { // diff:2
			return 21
		} else if rnd < 16971004547812787485 { // diff:2
			return 22
		} else if rnd < 17708874310761169550 { // diff:1
			return 23
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 24
		} // num:25
	case 26:
		if rnd < 709490156681136600 { // diff:15; 26÷15=1.7333333333333334
			return 0
		} else if rnd < 1418980313362273201 { // diff:14
			return 1
		} else if rnd < 2128470470043409801 { // diff:14
			return 2
		} else if rnd < 2837960626724546402 { // diff:13
			return 3
		} else if rnd < 3547450783405683002 { // diff:13
			return 4
		} else if rnd < 4256940940086819603 { // diff:12
			return 5
		} else if rnd < 4966431096767956204 { // diff:11
			return 6
		} else if rnd < 5675921253449092804 { // diff:11
			return 7
		} else if rnd < 6385411410130229405 { // diff:10
			return 8
		} else if rnd < 7094901566811366005 { // diff:10
			return 9
		} else if rnd < 7804391723492502606 { // diff:9
			return 10
		} else if rnd < 8513881880173639206 { // diff:9
			return 11
		} else if rnd < 9223372036854775807 { // diff:8
			return 12
		} else if rnd < 9932862193535912408 { // diff:7
			return 13
		} else if rnd < 10642352350217049008 { // diff:7
			return 14
		} else if rnd < 11351842506898185609 { // diff:6
			return 15
		} else if rnd < 12061332663579322209 { // diff:6
			return 16
		} else if rnd < 12770822820260458810 { // diff:5
			return 17
		} else if rnd < 13480312976941595410 { // diff:5
			return 18
		} else if rnd < 14189803133622732011 { // diff:4
			return 19
		} else if rnd < 14899293290303868612 { // diff:3
			return 20
		} else if rnd < 15608783446985005212 { // diff:3
			return 21
		} else if rnd < 16318273603666141813 { // diff:2
			return 22
		} else if rnd < 17027763760347278413 { // diff:2
			return 23
		} else if rnd < 17737253917028415014 { // diff:1
			return 24
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 25
		} // num:26
	case 27:
		if rnd < 683212743470724133 { // diff:24; 27÷24=1.125
			return 0
		} else if rnd < 1366425486941448267 { // diff:23
			return 1
		} else if rnd < 2049638230412172401 { // diff:22
			return 2
		} else if rnd < 2732850973882896535 { // diff:21
			return 3
		} else if rnd < 3416063717353620669 { // diff:20
			return 4
		} else if rnd < 4099276460824344803 { // diff:19
			return 5
		} else if rnd < 4782489204295068937 { // diff:18
			return 6
		} else if rnd < 5465701947765793071 { // diff:17
			return 7
		} else if rnd < 6148914691236517204 { // diff:17
			return 8
		} else if rnd < 6832127434707241338 { // diff:16
			return 9
		} else if rnd < 7515340178177965472 { // diff:15
			return 10
		} else if rnd < 8198552921648689606 { // diff:14
			return 11
		} else if rnd < 8881765665119413740 { // diff:13
			return 12
		} else if rnd < 9564978408590137874 { // diff:12
			return 13
		} else if rnd < 10248191152060862008 { // diff:11
			return 14
		} else if rnd < 10931403895531586142 { // diff:10
			return 15
		} else if rnd < 11614616639002310276 { // diff:9
			return 16
		} else if rnd < 12297829382473034409 { // diff:9
			return 17
		} else if rnd < 12981042125943758543 { // diff:8
			return 18
		} else if rnd < 13664254869414482677 { // diff:7
			return 19
		} else if rnd < 14347467612885206811 { // diff:6
			return 20
		} else if rnd < 15030680356355930945 { // diff:5
			return 21
		} else if rnd < 15713893099826655079 { // diff:4
			return 22
		} else if rnd < 16397105843297379213 { // diff:3
			return 23
		} else if rnd < 17080318586768103347 { // diff:2
			return 24
		} else if rnd < 17763531330238827481 { // diff:1
			return 25
		} else { // < 18446744073709551615; calc:18446744073709551591
			return 26
		} // num:27
	case 28:
		if rnd < 658812288346769700 { // diff:15; 28÷15=1.8666666666666667
			return 0
		} else if rnd < 1317624576693539401 { // diff:14
			return 1
		} else if rnd < 1976436865040309101 { // diff:14
			return 2
		} else if rnd < 2635249153387078802 { // diff:13
			return 3
		} else if rnd < 3294061441733848502 { // diff:13
			return 4
		} else if rnd < 3952873730080618203 { // diff:12
			return 5
		} else if rnd < 4611686018427387903 { // diff:12
			return 6
		} else if rnd < 5270498306774157604 { // diff:11
			return 7
		} else if rnd < 5929310595120927304 { // diff:11
			return 8
		} else if rnd < 6588122883467697005 { // diff:10
			return 9
		} else if rnd < 7246935171814466705 { // diff:10
			return 10
		} else if rnd < 7905747460161236406 { // diff:9
			return 11
		} else if rnd < 8564559748508006106 { // diff:9
			return 12
		} else if rnd < 9223372036854775807 { // diff:8
			return 13
		} else if rnd < 9882184325201545508 { // diff:7
			return 14
		} else if rnd < 10540996613548315208 { // diff:7
			return 15
		} else if rnd < 11199808901895084909 { // diff:6
			return 16
		} else if rnd < 11858621190241854609 { // diff:6
			return 17
		} else if rnd < 12517433478588624310 { // diff:5
			return 18
		} else if rnd < 13176245766935394010 { // diff:5
			return 19
		} else if rnd < 13835058055282163711 { // diff:4
			return 20
		} else if rnd < 14493870343628933411 { // diff:4
			return 21
		} else if rnd < 15152682631975703112 { // diff:3
			return 22
		} else if rnd < 15811494920322472812 { // diff:3
			return 23
		} else if rnd < 16470307208669242513 { // diff:2
			return 24
		} else if rnd < 17129119497016012213 { // diff:2
			return 25
		} else if rnd < 17787931785362781914 { // diff:1
			return 26
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 27
		} // num:28
	case 29:
		if rnd < 636094623231363848 { // diff:23; 29÷23=1.2608695652173914
			return 0
		} else if rnd < 1272189246462727697 { // diff:22
			return 1
		} else if rnd < 1908283869694091546 { // diff:21
			return 2
		} else if rnd < 2544378492925455395 { // diff:20
			return 3
		} else if rnd < 3180473116156819243 { // diff:20
			return 4
		} else if rnd < 3816567739388183092 { // diff:19
			return 5
		} else if rnd < 4452662362619546941 { // diff:18
			return 6
		} else if rnd < 5088756985850910790 { // diff:17
			return 7
		} else if rnd < 5724851609082274639 { // diff:16
			return 8
		} else if rnd < 6360946232313638487 { // diff:16
			return 9
		} else if rnd < 6997040855545002336 { // diff:15
			return 10
		} else if rnd < 7633135478776366185 { // diff:14
			return 11
		} else if rnd < 8269230102007730034 { // diff:13
			return 12
		} else if rnd < 8905324725239093883 { // diff:12
			return 13
		} else if rnd < 9541419348470457731 { // diff:12
			return 14
		} else if rnd < 10177513971701821580 { // diff:11
			return 15
		} else if rnd < 10813608594933185429 { // diff:10
			return 16
		} else if rnd < 11449703218164549278 { // diff:9
			return 17
		} else if rnd < 12085797841395913127 { // diff:8
			return 18
		} else if rnd < 12721892464627276975 { // diff:8
			return 19
		} else if rnd < 13357987087858640824 { // diff:7
			return 20
		} else if rnd < 13994081711090004673 { // diff:6
			return 21
		} else if rnd < 14630176334321368522 { // diff:5
			return 22
		} else if rnd < 15266270957552732371 { // diff:4
			return 23
		} else if rnd < 15902365580784096219 { // diff:4
			return 24
		} else if rnd < 16538460204015460068 { // diff:3
			return 25
		} else if rnd < 17174554827246823917 { // diff:2
			return 26
		} else if rnd < 17810649450478187766 { // diff:1
			return 27
		} else { // < 18446744073709551615; calc:18446744073709551592
			return 28
		} // num:29
	case 30:
		if rnd < 614891469123651720 { // diff:15; 30÷15=2
			return 0
		} else if rnd < 1229782938247303440 { // diff:15
			return 1
		} else if rnd < 1844674407370955161 { // diff:14
			return 2
		} else if rnd < 2459565876494606881 { // diff:14
			return 3
		} else if rnd < 3074457345618258602 { // diff:13
			return 4
		} else if rnd < 3689348814741910322 { // diff:13
			return 5
		} else if rnd < 4304240283865562043 { // diff:12
			return 6
		} else if rnd < 4919131752989213763 { // diff:12
			return 7
		} else if rnd < 5534023222112865484 { // diff:11
			return 8
		} else if rnd < 6148914691236517204 { // diff:11
			return 9
		} else if rnd < 6763806160360168925 { // diff:10
			return 10
		} else if rnd < 7378697629483820645 { // diff:10
			return 11
		} else if rnd < 7993589098607472366 { // diff:9
			return 12
		} else if rnd < 8608480567731124086 { // diff:9
			return 13
		} else if rnd < 9223372036854775807 { // diff:8
			return 14
		} else if rnd < 9838263505978427527 { // diff:8
			return 15
		} else if rnd < 10453154975102079248 { // diff:7
			return 16
		} else if rnd < 11068046444225730968 { // diff:7
			return 17
		} else if rnd < 11682937913349382689 { // diff:6
			return 18
		} else if rnd < 12297829382473034409 { // diff:6
			return 19
		} else if rnd < 12912720851596686130 { // diff:5
			return 20
		} else if rnd < 13527612320720337850 { // diff:5
			return 21
		} else if rnd < 14142503789843989571 { // diff:4
			return 22
		} else if rnd < 14757395258967641291 { // diff:4
			return 23
		} else if rnd < 15372286728091293012 { // diff:3
			return 24
		} else if rnd < 15987178197214944732 { // diff:3
			return 25
		} else if rnd < 16602069666338596453 { // diff:2
			return 26
		} else if rnd < 17216961135462248173 { // diff:2
			return 27
		} else if rnd < 17831852604585899894 { // diff:1
			return 28
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 29
		} // num:30
	case 31:
		if rnd < 595056260442243600 { // diff:15; 31÷15=2.066666666666667
			return 0
		} else if rnd < 1190112520884487200 { // diff:15
			return 1
		} else if rnd < 1785168781326730801 { // diff:14
			return 2
		} else if rnd < 2380225041768974401 { // diff:14
			return 3
		} else if rnd < 2975281302211218002 { // diff:13
			return 4
		} else if rnd < 3570337562653461602 { // diff:13
			return 5
		} else if rnd < 4165393823095705203 { // diff:12
			return 6
		} else if rnd < 4760450083537948803 { // diff:12
			return 7
		} else if rnd < 5355506343980192404 { // diff:11
			return 8
		} else if rnd < 5950562604422436004 { // diff:11
			return 9
		} else if rnd < 6545618864864679605 { // diff:10
			return 10
		} else if rnd < 7140675125306923205 { // diff:10
			return 11
		} else if rnd < 7735731385749166806 { // diff:9
			return 12
		} else if rnd < 8330787646191410406 { // diff:9
			return 13
		} else if rnd < 8925843906633654007 { // diff:8
			return 14
		} else if rnd < 9520900167075897607 { // diff:8
			return 15
		} else if rnd < 10115956427518141208 { // diff:7
			return 16
		} else if rnd < 10711012687960384808 { // diff:7
			return 17
		} else if rnd < 11306068948402628409 { // diff:6
			return 18
		} else if rnd < 11901125208844872009 { // diff:6
			return 19
		} else if rnd < 12496181469287115610 { // diff:5
			return 20
		} else if rnd < 13091237729729359210 { // diff:5
			return 21
		} else if rnd < 13686293990171602811 { // diff:4
			return 22
		} else if rnd < 14281350250613846411 { // diff:4
			return 23
		} else if rnd < 14876406511056090012 { // diff:3
			return 24
		} else if rnd < 15471462771498333612 { // diff:3
			return 25
		} else if rnd < 16066519031940577213 { // diff:2
			return 26
		} else if rnd < 16661575292382820813 { // diff:2
			return 27
		} else if rnd < 17256631552825064414 { // diff:1
			return 28
		} else if rnd < 17851687813267308014 { // diff:1
			return 29
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 30
		} // num:31
	case 32:
		if rnd < 576460752303423487 { // diff:31; 32÷31=1.032258064516129
			return 0
		} else if rnd < 1152921504606846975 { // diff:30
			return 1
		} else if rnd < 1729382256910270463 { // diff:29
			return 2
		} else if rnd < 2305843009213693951 { // diff:28
			return 3
		} else if rnd < 2882303761517117439 { // diff:27
			return 4
		} else if rnd < 3458764513820540927 { // diff:26
			return 5
		} else if rnd < 4035225266123964415 { // diff:25
			return 6
		} else if rnd < 4611686018427387903 { // diff:24
			return 7
		} else if rnd < 5188146770730811391 { // diff:23
			return 8
		} else if rnd < 5764607523034234879 { // diff:22
			return 9
		} else if rnd < 6341068275337658367 { // diff:21
			return 10
		} else if rnd < 6917529027641081855 { // diff:20
			return 11
		} else if rnd < 7493989779944505343 { // diff:19
			return 12
		} else if rnd < 8070450532247928831 { // diff:18
			return 13
		} else if rnd < 8646911284551352319 { // diff:17
			return 14
		} else if rnd < 9223372036854775807 { // diff:16
			return 15
		} else if rnd < 9799832789158199295 { // diff:15
			return 16
		} else if rnd < 10376293541461622783 { // diff:14
			return 17
		} else if rnd < 10952754293765046271 { // diff:13
			return 18
		} else if rnd < 11529215046068469759 { // diff:12
			return 19
		} else if rnd < 12105675798371893247 { // diff:11
			return 20
		} else if rnd < 12682136550675316735 { // diff:10
			return 21
		} else if rnd < 13258597302978740223 { // diff:9
			return 22
		} else if rnd < 13835058055282163711 { // diff:8
			return 23
		} else if rnd < 14411518807585587199 { // diff:7
			return 24
		} else if rnd < 14987979559889010687 { // diff:6
			return 25
		} else if rnd < 15564440312192434175 { // diff:5
			return 26
		} else if rnd < 16140901064495857663 { // diff:4
			return 27
		} else if rnd < 16717361816799281151 { // diff:3
			return 28
		} else if rnd < 17293822569102704639 { // diff:2
			return 29
		} else if rnd < 17870283321406128127 { // diff:1
			return 30
		} else { // < 18446744073709551615; calc:18446744073709551584
			return 31
		} // num:32
	case 33:
		if rnd < 558992244657865200 { // diff:15; 33÷15=2.2
			return 0
		} else if rnd < 1117984489315730400 { // diff:15
			return 1
		} else if rnd < 1676976733973595601 { // diff:14
			return 2
		} else if rnd < 2235968978631460801 { // diff:14
			return 3
		} else if rnd < 2794961223289326002 { // diff:13
			return 4
		} else if rnd < 3353953467947191202 { // diff:13
			return 5
		} else if rnd < 3912945712605056403 { // diff:12
			return 6
		} else if rnd < 4471937957262921603 { // diff:12
			return 7
		} else if rnd < 5030930201920786804 { // diff:11
			return 8
		} else if rnd < 5589922446578652004 { // diff:11
			return 9
		} else if rnd < 6148914691236517204 { // diff:11
			return 10
		} else if rnd < 6707906935894382405 { // diff:10
			return 11
		} else if rnd < 7266899180552247605 { // diff:10
			return 12
		} else if rnd < 7825891425210112806 { // diff:9
			return 13
		} else if rnd < 8384883669867978006 { // diff:9
			return 14
		} else if rnd < 8943875914525843207 { // diff:8
			return 15
		} else if rnd < 9502868159183708407 { // diff:8
			return 16
		} else if rnd < 10061860403841573608 { // diff:7
			return 17
		} else if rnd < 10620852648499438808 { // diff:7
			return 18
		} else if rnd < 11179844893157304009 { // diff:6
			return 19
		} else if rnd < 11738837137815169209 { // diff:6
			return 20
		} else if rnd < 12297829382473034410 { // diff:5
			return 21
		} else if rnd < 12856821627130899610 { // diff:5
			return 22
		} else if rnd < 13415813871788764810 { // diff:5
			return 23
		} else if rnd < 13974806116446630011 { // diff:4
			return 24
		} else if rnd < 14533798361104495211 { // diff:4
			return 25
		} else if rnd < 15092790605762360412 { // diff:3
			return 26
		} else if rnd < 15651782850420225612 { // diff:3
			return 27
		} else if rnd < 16210775095078090813 { // diff:2
			return 28
		} else if rnd < 16769767339735956013 { // diff:2
			return 29
		} else if rnd < 17328759584393821214 { // diff:1
			return 30
		} else if rnd < 17887751829051686414 { // diff:1
			return 31
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 32
		} // num:33
	case 34:
		if rnd < 542551296285575047 { // diff:17; 34÷17=2
			return 0
		} else if rnd < 1085102592571150094 { // diff:17
			return 1
		} else if rnd < 1627653888856725142 { // diff:16
			return 2
		} else if rnd < 2170205185142300189 { // diff:16
			return 3
		} else if rnd < 2712756481427875237 { // diff:15
			return 4
		} else if rnd < 3255307777713450284 { // diff:15
			return 5
		} else if rnd < 3797859073999025332 { // diff:14
			return 6
		} else if rnd < 4340410370284600379 { // diff:14
			return 7
		} else if rnd < 4882961666570175427 { // diff:13
			return 8
		} else if rnd < 5425512962855750474 { // diff:13
			return 9
		} else if rnd < 5968064259141325522 { // diff:12
			return 10
		} else if rnd < 6510615555426900569 { // diff:12
			return 11
		} else if rnd < 7053166851712475617 { // diff:11
			return 12
		} else if rnd < 7595718147998050664 { // diff:11
			return 13
		} else if rnd < 8138269444283625712 { // diff:10
			return 14
		} else if rnd < 8680820740569200759 { // diff:10
			return 15
		} else if rnd < 9223372036854775807 { // diff:9
			return 16
		} else if rnd < 9765923333140350854 { // diff:9
			return 17
		} else if rnd < 10308474629425925902 { // diff:8
			return 18
		} else if rnd < 10851025925711500949 { // diff:8
			return 19
		} else if rnd < 11393577221997075997 { // diff:7
			return 20
		} else if rnd < 11936128518282651044 { // diff:7
			return 21
		} else if rnd < 12478679814568226092 { // diff:6
			return 22
		} else if rnd < 13021231110853801139 { // diff:6
			return 23
		} else if rnd < 13563782407139376187 { // diff:5
			return 24
		} else if rnd < 14106333703424951234 { // diff:5
			return 25
		} else if rnd < 14648884999710526282 { // diff:4
			return 26
		} else if rnd < 15191436295996101329 { // diff:4
			return 27
		} else if rnd < 15733987592281676377 { // diff:3
			return 28
		} else if rnd < 16276538888567251424 { // diff:3
			return 29
		} else if rnd < 16819090184852826472 { // diff:2
			return 30
		} else if rnd < 17361641481138401519 { // diff:2
			return 31
		} else if rnd < 17904192777423976567 { // diff:1
			return 32
		} else { // < 18446744073709551615; calc:18446744073709551598
			return 33
		} // num:34
	case 35:
		if rnd < 527049830677415760 { // diff:15; 35÷15=2.3333333333333335
			return 0
		} else if rnd < 1054099661354831520 { // diff:15
			return 1
		} else if rnd < 1581149492032247281 { // diff:14
			return 2
		} else if rnd < 2108199322709663041 { // diff:14
			return 3
		} else if rnd < 2635249153387078802 { // diff:13
			return 4
		} else if rnd < 3162298984064494562 { // diff:13
			return 5
		} else if rnd < 3689348814741910322 { // diff:13
			return 6
		} else if rnd < 4216398645419326083 { // diff:12
			return 7
		} else if rnd < 4743448476096741843 { // diff:12
			return 8
		} else if rnd < 5270498306774157604 { // diff:11
			return 9
		} else if rnd < 5797548137451573364 { // diff:11
			return 10
		} else if rnd < 6324597968128989125 { // diff:10
			return 11
		} else if rnd < 6851647798806404885 { // diff:10
			return 12
		} else if rnd < 7378697629483820645 { // diff:10
			return 13
		} else if rnd < 7905747460161236406 { // diff:9
			return 14
		} else if rnd < 8432797290838652166 { // diff:9
			return 15
		} else if rnd < 8959847121516067927 { // diff:8
			return 16
		} else if rnd < 9486896952193483687 { // diff:8
			return 17
		} else if rnd < 10013946782870899448 { // diff:7
			return 18
		} else if rnd < 10540996613548315208 { // diff:7
			return 19
		} else if rnd < 11068046444225730968 { // diff:7
			return 20
		} else if rnd < 11595096274903146729 { // diff:6
			return 21
		} else if rnd < 12122146105580562489 { // diff:6
			return 22
		} else if rnd < 12649195936257978250 { // diff:5
			return 23
		} else if rnd < 13176245766935394010 { // diff:5
			return 24
		} else if rnd < 13703295597612809771 { // diff:4
			return 25
		} else if rnd < 14230345428290225531 { // diff:4
			return 26
		} else if rnd < 14757395258967641292 { // diff:3
			return 27
		} else if rnd < 15284445089645057052 { // diff:3
			return 28
		} else if rnd < 15811494920322472812 { // diff:3
			return 29
		} else if rnd < 16338544750999888573 { // diff:2
			return 30
		} else if rnd < 16865594581677304333 { // diff:2
			return 31
		} else if rnd < 17392644412354720094 { // diff:1
			return 32
		} else if rnd < 17919694243032135854 { // diff:1
			return 33
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 34
		} // num:35
	case 36:
		if rnd < 512409557603043100 { // diff:15; 36÷15=2.4
			return 0
		} else if rnd < 1024819115206086200 { // diff:15
			return 1
		} else if rnd < 1537228672809129301 { // diff:14
			return 2
		} else if rnd < 2049638230412172401 { // diff:14
			return 3
		} else if rnd < 2562047788015215502 { // diff:13
			return 4
		} else if rnd < 3074457345618258602 { // diff:13
			return 5
		} else if rnd < 3586866903221301702 { // diff:13
			return 6
		} else if rnd < 4099276460824344803 { // diff:12
			return 7
		} else if rnd < 4611686018427387903 { // diff:12
			return 8
		} else if rnd < 5124095576030431004 { // diff:11
			return 9
		} else if rnd < 5636505133633474104 { // diff:11
			return 10
		} else if rnd < 6148914691236517204 { // diff:11
			return 11
		} else if rnd < 6661324248839560305 { // diff:10
			return 12
		} else if rnd < 7173733806442603405 { // diff:10
			return 13
		} else if rnd < 7686143364045646506 { // diff:9
			return 14
		} else if rnd < 8198552921648689606 { // diff:9
			return 15
		} else if rnd < 8710962479251732707 { // diff:8
			return 16
		} else if rnd < 9223372036854775807 { // diff:8
			return 17
		} else if rnd < 9735781594457818907 { // diff:8
			return 18
		} else if rnd < 10248191152060862008 { // diff:7
			return 19
		} else if rnd < 10760600709663905108 { // diff:7
			return 20
		} else if rnd < 11273010267266948209 { // diff:6
			return 21
		} else if rnd < 11785419824869991309 { // diff:6
			return 22
		} else if rnd < 12297829382473034410 { // diff:5
			return 23
		} else if rnd < 12810238940076077510 { // diff:5
			return 24
		} else if rnd < 13322648497679120610 { // diff:5
			return 25
		} else if rnd < 13835058055282163711 { // diff:4
			return 26
		} else if rnd < 14347467612885206811 { // diff:4
			return 27
		} else if rnd < 14859877170488249912 { // diff:3
			return 28
		} else if rnd < 15372286728091293012 { // diff:3
			return 29
		} else if rnd < 15884696285694336112 { // diff:3
			return 30
		} else if rnd < 16397105843297379213 { // diff:2
			return 31
		} else if rnd < 16909515400900422313 { // diff:2
			return 32
		} else if rnd < 17421924958503465414 { // diff:1
			return 33
		} else if rnd < 17934334516106508514 { // diff:1
			return 34
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 35
		} // num:36
	case 37:
		if rnd < 498560650640798692 { // diff:11; 37÷11=3.3636363636363638
			return 0
		} else if rnd < 997121301281597384 { // diff:11
			return 1
		} else if rnd < 1495681951922396076 { // diff:11
			return 2
		} else if rnd < 1994242602563194769 { // diff:10
			return 3
		} else if rnd < 2492803253203993461 { // diff:10
			return 4
		} else if rnd < 2991363903844792153 { // diff:10
			return 5
		} else if rnd < 3489924554485590846 { // diff:9
			return 6
		} else if rnd < 3988485205126389538 { // diff:9
			return 7
		} else if rnd < 4487045855767188230 { // diff:9
			return 8
		} else if rnd < 4985606506407986922 { // diff:9
			return 9
		} else if rnd < 5484167157048785615 { // diff:8
			return 10
		} else if rnd < 5982727807689584307 { // diff:8
			return 11
		} else if rnd < 6481288458330382999 { // diff:8
			return 12
		} else if rnd < 6979849108971181692 { // diff:7
			return 13
		} else if rnd < 7478409759611980384 { // diff:7
			return 14
		} else if rnd < 7976970410252779076 { // diff:7
			return 15
		} else if rnd < 8475531060893577769 { // diff:6
			return 16
		} else if rnd < 8974091711534376461 { // diff:6
			return 17
		} else if rnd < 9472652362175175153 { // diff:6
			return 18
		} else if rnd < 9971213012815973845 { // diff:6
			return 19
		} else if rnd < 10469773663456772538 { // diff:5
			return 20
		} else if rnd < 10968334314097571230 { // diff:5
			return 21
		} else if rnd < 11466894964738369922 { // diff:5
			return 22
		} else if rnd < 11965455615379168615 { // diff:4
			return 23
		} else if rnd < 12464016266019967307 { // diff:4
			return 24
		} else if rnd < 12962576916660765999 { // diff:4
			return 25
		} else if rnd < 13461137567301564692 { // diff:3
			return 26
		} else if rnd < 13959698217942363384 { // diff:3
			return 27
		} else if rnd < 14458258868583162076 { // diff:3
			return 28
		} else if rnd < 14956819519223960768 { // diff:3
			return 29
		} else if rnd < 15455380169864759461 { // diff:2
			return 30
		} else if rnd < 15953940820505558153 { // diff:2
			return 31
		} else if rnd < 16452501471146356845 { // diff:2
			return 32
		} else if rnd < 16951062121787155538 { // diff:1
			return 33
		} else if rnd < 17449622772427954230 { // diff:1
			return 34
		} else if rnd < 17948183423068752922 { // diff:1
			return 35
		} else { // < 18446744073709551615; calc:18446744073709551604
			return 36
		} // num:37
	case 38:
		if rnd < 485440633518672410 { // diff:35; 38÷35=1.0857142857142856
			return 0
		} else if rnd < 970881267037344821 { // diff:34
			return 1
		} else if rnd < 1456321900556017232 { // diff:33
			return 2
		} else if rnd < 1941762534074689643 { // diff:32
			return 3
		} else if rnd < 2427203167593362054 { // diff:31
			return 4
		} else if rnd < 2912643801112034465 { // diff:30
			return 5
		} else if rnd < 3398084434630706876 { // diff:29
			return 6
		} else if rnd < 3883525068149379287 { // diff:28
			return 7
		} else if rnd < 4368965701668051698 { // diff:27
			return 8
		} else if rnd < 4854406335186724109 { // diff:26
			return 9
		} else if rnd < 5339846968705396520 { // diff:25
			return 10
		} else if rnd < 5825287602224068931 { // diff:24
			return 11
		} else if rnd < 6310728235742741341 { // diff:24
			return 12
		} else if rnd < 6796168869261413752 { // diff:23
			return 13
		} else if rnd < 7281609502780086163 { // diff:22
			return 14
		} else if rnd < 7767050136298758574 { // diff:21
			return 15
		} else if rnd < 8252490769817430985 { // diff:20
			return 16
		} else if rnd < 8737931403336103396 { // diff:19
			return 17
		} else if rnd < 9223372036854775807 { // diff:18
			return 18
		} else if rnd < 9708812670373448218 { // diff:17
			return 19
		} else if rnd < 10194253303892120629 { // diff:16
			return 20
		} else if rnd < 10679693937410793040 { // diff:15
			return 21
		} else if rnd < 11165134570929465451 { // diff:14
			return 22
		} else if rnd < 11650575204448137862 { // diff:13
			return 23
		} else if rnd < 12136015837966810273 { // diff:12
			return 24
		} else if rnd < 12621456471485482683 { // diff:12
			return 25
		} else if rnd < 13106897105004155094 { // diff:11
			return 26
		} else if rnd < 13592337738522827505 { // diff:10
			return 27
		} else if rnd < 14077778372041499916 { // diff:9
			return 28
		} else if rnd < 14563219005560172327 { // diff:8
			return 29
		} else if rnd < 15048659639078844738 { // diff:7
			return 30
		} else if rnd < 15534100272597517149 { // diff:6
			return 31
		} else if rnd < 16019540906116189560 { // diff:5
			return 32
		} else if rnd < 16504981539634861971 { // diff:4
			return 33
		} else if rnd < 16990422173153534382 { // diff:3
			return 34
		} else if rnd < 17475862806672206793 { // diff:2
			return 35
		} else if rnd < 17961303440190879204 { // diff:1
			return 36
		} else { // < 18446744073709551615; calc:18446744073709551580
			return 37
		} // num:38
	case 39:
		if rnd < 472993437787424400 { // diff:15; 39÷15=2.6
			return 0
		} else if rnd < 945986875574848800 { // diff:15
			return 1
		} else if rnd < 1418980313362273201 { // diff:14
			return 2
		} else if rnd < 1891973751149697601 { // diff:14
			return 3
		} else if rnd < 2364967188937122001 { // diff:14
			return 4
		} else if rnd < 2837960626724546402 { // diff:13
			return 5
		} else if rnd < 3310954064511970802 { // diff:13
			return 6
		} else if rnd < 3783947502299395203 { // diff:12
			return 7
		} else if rnd < 4256940940086819603 { // diff:12
			return 8
		} else if rnd < 4729934377874244003 { // diff:12
			return 9
		} else if rnd < 5202927815661668404 { // diff:11
			return 10
		} else if rnd < 5675921253449092804 { // diff:11
			return 11
		} else if rnd < 6148914691236517204 { // diff:11
			return 12
		} else if rnd < 6621908129023941605 { // diff:10
			return 13
		} else if rnd < 7094901566811366005 { // diff:10
			return 14
		} else if rnd < 7567895004598790406 { // diff:9
			return 15
		} else if rnd < 8040888442386214806 { // diff:9
			return 16
		} else if rnd < 8513881880173639206 { // diff:9
			return 17
		} else if rnd < 8986875317961063607 { // diff:8
			return 18
		} else if rnd < 9459868755748488007 { // diff:8
			return 19
		} else if rnd < 9932862193535912408 { // diff:7
			return 20
		} else if rnd < 10405855631323336808 { // diff:7
			return 21
		} else if rnd < 10878849069110761208 { // diff:7
			return 22
		} else if rnd < 11351842506898185609 { // diff:6
			return 23
		} else if rnd < 11824835944685610009 { // diff:6
			return 24
		} else if rnd < 12297829382473034409 { // diff:6
			return 25
		} else if rnd < 12770822820260458810 { // diff:5
			return 26
		} else if rnd < 13243816258047883210 { // diff:5
			return 27
		} else if rnd < 13716809695835307611 { // diff:4
			return 28
		} else if rnd < 14189803133622732011 { // diff:4
			return 29
		} else if rnd < 14662796571410156411 { // diff:4
			return 30
		} else if rnd < 15135790009197580812 { // diff:3
			return 31
		} else if rnd < 15608783446985005212 { // diff:3
			return 32
		} else if rnd < 16081776884772429613 { // diff:2
			return 33
		} else if rnd < 16554770322559854013 { // diff:2
			return 34
		} else if rnd < 17027763760347278413 { // diff:2
			return 35
		} else if rnd < 17500757198134702814 { // diff:1
			return 36
		} else if rnd < 17973750635922127214 { // diff:1
			return 37
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 38
		} // num:39
	case 40:
		if rnd < 461168601842738790 { // diff:15; 40÷15=2.6666666666666665
			return 0
		} else if rnd < 922337203685477580 { // diff:15
			return 1
		} else if rnd < 1383505805528216371 { // diff:14
			return 2
		} else if rnd < 1844674407370955161 { // diff:14
			return 3
		} else if rnd < 2305843009213693951 { // diff:14
			return 4
		} else if rnd < 2767011611056432742 { // diff:13
			return 5
		} else if rnd < 3228180212899171532 { // diff:13
			return 6
		} else if rnd < 3689348814741910322 { // diff:13
			return 7
		} else if rnd < 4150517416584649113 { // diff:12
			return 8
		} else if rnd < 4611686018427387903 { // diff:12
			return 9
		} else if rnd < 5072854620270126694 { // diff:11
			return 10
		} else if rnd < 5534023222112865484 { // diff:11
			return 11
		} else if rnd < 5995191823955604274 { // diff:11
			return 12
		} else if rnd < 6456360425798343065 { // diff:10
			return 13
		} else if rnd < 6917529027641081855 { // diff:10
			return 14
		} else if rnd < 7378697629483820646 { // diff:9
			return 15
		} else if rnd < 7839866231326559436 { // diff:9
			return 16
		} else if rnd < 8301034833169298226 { // diff:9
			return 17
		} else if rnd < 8762203435012037017 { // diff:8
			return 18
		} else if rnd < 9223372036854775807 { // diff:8
			return 19
		} else if rnd < 9684540638697514597 { // diff:8
			return 20
		} else if rnd < 10145709240540253388 { // diff:7
			return 21
		} else if rnd < 10606877842382992178 { // diff:7
			return 22
		} else if rnd < 11068046444225730968 { // diff:7
			return 23
		} else if rnd < 11529215046068469759 { // diff:6
			return 24
		} else if rnd < 11990383647911208549 { // diff:6
			return 25
		} else if rnd < 12451552249753947340 { // diff:5
			return 26
		} else if rnd < 12912720851596686130 { // diff:5
			return 27
		} else if rnd < 13373889453439424920 { // diff:5
			return 28
		} else if rnd < 13835058055282163711 { // diff:4
			return 29
		} else if rnd < 14296226657124902501 { // diff:4
			return 30
		} else if rnd < 14757395258967641291 { // diff:4
			return 31
		} else if rnd < 15218563860810380082 { // diff:3
			return 32
		} else if rnd < 15679732462653118872 { // diff:3
			return 33
		} else if rnd < 16140901064495857663 { // diff:2
			return 34
		} else if rnd < 16602069666338596453 { // diff:2
			return 35
		} else if rnd < 17063238268181335243 { // diff:2
			return 36
		} else if rnd < 17524406870024074034 { // diff:1
			return 37
		} else if rnd < 17985575471866812824 { // diff:1
			return 38
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 39
		} // num:40
	case 41:
		if rnd < 449920587163647600 { // diff:15; 41÷15=2.7333333333333334
			return 0
		} else if rnd < 899841174327295200 { // diff:15
			return 1
		} else if rnd < 1349761761490942801 { // diff:14
			return 2
		} else if rnd < 1799682348654590401 { // diff:14
			return 3
		} else if rnd < 2249602935818238001 { // diff:14
			return 4
		} else if rnd < 2699523522981885602 { // diff:13
			return 5
		} else if rnd < 3149444110145533202 { // diff:13
			return 6
		} else if rnd < 3599364697309180802 { // diff:13
			return 7
		} else if rnd < 4049285284472828403 { // diff:12
			return 8
		} else if rnd < 4499205871636476003 { // diff:12
			return 9
		} else if rnd < 4949126458800123604 { // diff:11
			return 10
		} else if rnd < 5399047045963771204 { // diff:11
			return 11
		} else if rnd < 5848967633127418804 { // diff:11
			return 12
		} else if rnd < 6298888220291066405 { // diff:10
			return 13
		} else if rnd < 6748808807454714005 { // diff:10
			return 14
		} else if rnd < 7198729394618361605 { // diff:10
			return 15
		} else if rnd < 7648649981782009206 { // diff:9
			return 16
		} else if rnd < 8098570568945656806 { // diff:9
			return 17
		} else if rnd < 8548491156109304406 { // diff:9
			return 18
		} else if rnd < 8998411743272952007 { // diff:8
			return 19
		} else if rnd < 9448332330436599607 { // diff:8
			return 20
		} else if rnd < 9898252917600247208 { // diff:7
			return 21
		} else if rnd < 10348173504763894808 { // diff:7
			return 22
		} else if rnd < 10798094091927542408 { // diff:7
			return 23
		} else if rnd < 11248014679091190009 { // diff:6
			return 24
		} else if rnd < 11697935266254837609 { // diff:6
			return 25
		} else if rnd < 12147855853418485209 { // diff:6
			return 26
		} else if rnd < 12597776440582132810 { // diff:5
			return 27
		} else if rnd < 13047697027745780410 { // diff:5
			return 28
		} else if rnd < 13497617614909428010 { // diff:5
			return 29
		} else if rnd < 13947538202073075611 { // diff:4
			return 30
		} else if rnd < 14397458789236723211 { // diff:4
			return 31
		} else if rnd < 14847379376400370812 { // diff:3
			return 32
		} else if rnd < 15297299963564018412 { // diff:3
			return 33
		} else if rnd < 15747220550727666012 { // diff:3
			return 34
		} else if rnd < 16197141137891313613 { // diff:2
			return 35
		} else if rnd < 16647061725054961213 { // diff:2
			return 36
		} else if rnd < 17096982312218608813 { // diff:2
			return 37
		} else if rnd < 17546902899382256414 { // diff:1
			return 38
		} else if rnd < 17996823486545904014 { // diff:1
			return 39
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 40
		} // num:41
	case 42:
		if rnd < 439208192231179800 { // diff:15; 42÷15=2.8
			return 0
		} else if rnd < 878416384462359600 { // diff:15
			return 1
		} else if rnd < 1317624576693539401 { // diff:14
			return 2
		} else if rnd < 1756832768924719201 { // diff:14
			return 3
		} else if rnd < 2196040961155899001 { // diff:14
			return 4
		} else if rnd < 2635249153387078802 { // diff:13
			return 5
		} else if rnd < 3074457345618258602 { // diff:13
			return 6
		} else if rnd < 3513665537849438402 { // diff:13
			return 7
		} else if rnd < 3952873730080618203 { // diff:12
			return 8
		} else if rnd < 4392081922311798003 { // diff:12
			return 9
		} else if rnd < 4831290114542977803 { // diff:12
			return 10
		} else if rnd < 5270498306774157604 { // diff:11
			return 11
		} else if rnd < 5709706499005337404 { // diff:11
			return 12
		} else if rnd < 6148914691236517204 { // diff:11
			return 13
		} else if rnd < 6588122883467697005 { // diff:10
			return 14
		} else if rnd < 7027331075698876805 { // diff:10
			return 15
		} else if rnd < 7466539267930056606 { // diff:9
			return 16
		} else if rnd < 7905747460161236406 { // diff:9
			return 17
		} else if rnd < 8344955652392416206 { // diff:9
			return 18
		} else if rnd < 8784163844623596007 { // diff:8
			return 19
		} else if rnd < 9223372036854775807 { // diff:8
			return 20
		} else if rnd < 9662580229085955607 { // diff:8
			return 21
		} else if rnd < 10101788421317135408 { // diff:7
			return 22
		} else if rnd < 10540996613548315208 { // diff:7
			return 23
		} else if rnd < 10980204805779495008 { // diff:7
			return 24
		} else if rnd < 11419412998010674809 { // diff:6
			return 25
		} else if rnd < 11858621190241854609 { // diff:6
			return 26
		} else if rnd < 12297829382473034409 { // diff:6
			return 27
		} else if rnd < 12737037574704214210 { // diff:5
			return 28
		} else if rnd < 13176245766935394010 { // diff:5
			return 29
		} else if rnd < 13615453959166573811 { // diff:4
			return 30
		} else if rnd < 14054662151397753611 { // diff:4
			return 31
		} else if rnd < 14493870343628933411 { // diff:4
			return 32
		} else if rnd < 14933078535860113212 { // diff:3
			return 33
		} else if rnd < 15372286728091293012 { // diff:3
			return 34
		} else if rnd < 15811494920322472812 { // diff:3
			return 35
		} else if rnd < 16250703112553652613 { // diff:2
			return 36
		} else if rnd < 16689911304784832413 { // diff:2
			return 37
		} else if rnd < 17129119497016012213 { // diff:2
			return 38
		} else if rnd < 17568327689247192014 { // diff:1
			return 39
		} else if rnd < 18007535881478371814 { // diff:1
			return 40
		} else { // < 18446744073709551615; calc:18446744073709551600
			return 41
		} // num:42
	case 43:
		if rnd < 428994048225803525 { // diff:40; 43÷40=1.075
			return 0
		} else if rnd < 857988096451607051 { // diff:39
			return 1
		} else if rnd < 1286982144677410577 { // diff:38
			return 2
		} else if rnd < 1715976192903214103 { // diff:37
			return 3
		} else if rnd < 2144970241129017629 { // diff:36
			return 4
		} else if rnd < 2573964289354821155 { // diff:35
			return 5
		} else if rnd < 3002958337580624681 { // diff:34
			return 6
		} else if rnd < 3431952385806428207 { // diff:33
			return 7
		} else if rnd < 3860946434032231733 { // diff:32
			return 8
		} else if rnd < 4289940482258035259 { // diff:31
			return 9
		} else if rnd < 4718934530483838785 { // diff:30
			return 10
		} else if rnd < 5147928578709642311 { // diff:29
			return 11
		} else if rnd < 5576922626935445837 { // diff:28
			return 12
		} else if rnd < 6005916675161249363 { // diff:27
			return 13
		} else if rnd < 6434910723387052888 { // diff:27
			return 14
		} else if rnd < 6863904771612856414 { // diff:26
			return 15
		} else if rnd < 7292898819838659940 { // diff:25
			return 16
		} else if rnd < 7721892868064463466 { // diff:24
			return 17
		} else if rnd < 8150886916290266992 { // diff:23
			return 18
		} else if rnd < 8579880964516070518 { // diff:22
			return 19
		} else if rnd < 9008875012741874044 { // diff:21
			return 20
		} else if rnd < 9437869060967677570 { // diff:20
			return 21
		} else if rnd < 9866863109193481096 { // diff:19
			return 22
		} else if rnd < 10295857157419284622 { // diff:18
			return 23
		} else if rnd < 10724851205645088148 { // diff:17
			return 24
		} else if rnd < 11153845253870891674 { // diff:16
			return 25
		} else if rnd < 11582839302096695200 { // diff:15
			return 26
		} else if rnd < 12011833350322498726 { // diff:14
			return 27
		} else if rnd < 12440827398548302251 { // diff:14
			return 28
		} else if rnd < 12869821446774105777 { // diff:13
			return 29
		} else if rnd < 13298815494999909303 { // diff:12
			return 30
		} else if rnd < 13727809543225712829 { // diff:11
			return 31
		} else if rnd < 14156803591451516355 { // diff:10
			return 32
		} else if rnd < 14585797639677319881 { // diff:9
			return 33
		} else if rnd < 15014791687903123407 { // diff:8
			return 34
		} else if rnd < 15443785736128926933 { // diff:7
			return 35
		} else if rnd < 15872779784354730459 { // diff:6
			return 36
		} else if rnd < 16301773832580533985 { // diff:5
			return 37
		} else if rnd < 16730767880806337511 { // diff:4
			return 38
		} else if rnd < 17159761929032141037 { // diff:3
			return 39
		} else if rnd < 17588755977257944563 { // diff:2
			return 40
		} else if rnd < 18017750025483748089 { // diff:1
			return 41
		} else { // < 18446744073709551615; calc:18446744073709551575
			return 42
		} // num:43
	}
	return -1
}
