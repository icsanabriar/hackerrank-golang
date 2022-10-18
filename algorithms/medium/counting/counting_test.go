// Copyright 2020 Iván Camilo Sanabria
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

// TestFullCountingSortFirstGivenCase implements the test given as first example on hackerrank.
func TestFullCountingSortFirstGivenCase(t *testing.T) {

	n := int32(4)
	systemIn := "0 a\n1 b\n0 c\n1 d\n"
	reader := bufio.NewReader(strings.NewReader(systemIn))

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	expected := "- c - d \n"

	countSort(stdout, readInput(reader, n))
	result := readTestFile()

	if result != expected {
		t.Errorf("Full counting sort first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestFullCountingSortSecondGivenCase implements the test given as second example on hackerrank.
func TestFullCountingSortSecondGivenCase(t *testing.T) {

	n := int32(20)
	systemIn := "0 ab\n6 cd\n0 ef\n6 gh\n4 ij\n0 ab\n6 cd\n0 ef\n6 gh\n0 ij\n4 that\n3 be\n0 to\n1 be\n5 question\n1 or\n2 not\n4 is\n2 to\n4 the"
	reader := bufio.NewReader(strings.NewReader(systemIn))

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	expected := "- - - - - to be or not to be - that is the question - - - - \n"

	countSort(stdout, readInput(reader, n))
	result := readTestFile()

	if result != expected {
		t.Errorf("Full counting sort second case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestFullCountingSortHiddenGivenCase implements the test given as hidden example on hackerrank.
func TestFullCountingSortHiddenGivenCase(t *testing.T) {

	n := int32(1000)
	systemIn := "24 oz\n36 xr\n43 cu\n44 oq\n55 qf\n2 oz\n96 bk\n1 tv\n2 jy\n27 wy\n30 kj\n11 mn\n85 mp\n3 wu\n" +
		"54 zq\n80 ee\n29 rv\n46 wi\n13 zu\n37 rv\n29 sc\n59 on\n76 ts\n86 wf\n49 gg\n9 yc\n2 gr\n46 ny\n56 ws\n" +
		"15 jr\n74 lc\n50 ag\n77 vk\n16 er\n84 cf\n82 dn\n12 ss\n4 pe\n35 pm\n28 qo\n47 ym\n54 lr\n34 bx\n18 jm\n" +
		"20 kn\n86 av\n17 rx\n35 zb\n14 il\n10 zo\n43 lk\n58 fk\n0 xs\n91 cn\n4 qq\n95 fw\n75 eo\n83 pr\n65 tc\n" +
		"11 vd\n32 kr\n55 so\n24 dk\n51 br\n38 db\n19 sr\n91 ir\n76 fe\n60 hr\n0 jz\n44 ju\n26 td\n70 ki\n62 zd\n" +
		"42 tb\n60 cg\n71 zj\n31 bw\n57 ti\n69 we\n28 ur\n24 pd\n24 tj\n65 dk\n94 cc\n27 az\n33 pm\n42 dp\n57 pz\n" +
		"49 dt\n73 ia\n17 li\n69 rv\n39 hm\n41 vq\n4 br\n73 xo\n70 nr\n45 hv\n75 yl\n58 ah\n39 fv\n48 tw\n45 eb\n" +
		"94 zy\n57 vo\n42 sq\n41 kn\n1 io\n13 iz\n29 yu\n21 tn\n91 ov\n75 jq\n91 yl\n46 br\n86 vc\n84 mm\n30 ls\n" +
		"78 rf\n99 uj\n24 nk\n15 ln\n76 nn\n83 jw\n71 hm\n72 hm\n46 pz\n10 ms\n40 un\n96 vf\n62 cn\n87 hj\n6 kz\n" +
		"41 bt\n78 qa\n97 yo\n26 qc\n1 gr\n61 eq\n7 gh\n73 dt\n7 zb\n88 at\n67 wq\n35 lo\n81 jw\n17 de\n24 vf\n" +
		"75 eh\n79 mo\n68 ol\n41 gs\n31 pa\n22 ji\n80 nu\n82 jl\n95 ln\n15 fy\n57 xd\n91 jb\n85 pv\n84 ra\n73 qg\n" +
		"73 lf\n51 py\n84 br\n68 nd\n13 ve\n98 xb\n24 xo\n0 sy\n15 dl\n28 oy\n94 el\n54 mf\n47 fu\n9 fb\n34 zg\n" +
		"99 tb\n56 ho\n74 wy\n88 gb\n72 uh\n66 su\n83 qf\n59 zy\n69 ye\n35 tj\n63 du\n25 cw\n73 og\n62 bv\n6 gp\n" +
		"92 ux\n16 eg\n27 vo\n19 vz\n94 ly\n17 ip\n23 bf\n28 wh\n62 tm\n87 jp\n87 la\n0 yv\n57 hq\n54 rd\n61 ga\n" +
		"32 ss\n37 sn\n18 ft\n83 hm\n50 es\n72 mz\n41 ka\n0 ww\n30 wl\n62 dg\n4 ov\n85 ch\n58 fm\n41 bq\n81 rm\n" +
		"92 ho\n22 sy\n49 gh\n19 zg\n25 fv\n58 mh\n25 tu\n5 km\n74 ob\n73 rv\n12 ia\n0 zp\n23 pt\n62 cw\n98 jc\n" +
		"13 hp\n47 oz\n85 ye\n46 qi\n40 qf\n60 iy\n3 ie\n85 ab\n40 ro\n33 ym\n43 rg\n26 yc\n88 vu\n85 os\n54 el\n" +
		"1 un\n10 ah\n93 fe\n78 zo\n78 xo\n14 hj\n20 ht\n88 px\n79 vv\n70 rn\n20 tc\n0 oh\n5 xh\n57 kt\n45 wb\n" +
		"22 jt\n24 ke\n67 fc\n33 uk\n59 jz\n42 nk\n16 gb\n91 ba\n48 nl\n86 gz\n89 ae\n73 fa\n19 yi\n95 kt\n61 jv\n" +
		"45 ue\n29 jq\n8 mz\n99 yr\n81 mp\n53 jw\n48 kb\n6 tj\n67 xt\n40 ou\n6 bc\n38 sa\n11 xk\n1 bq\n78 pf\n69 vr\n" +
		"19 ex\n52 ij\n59 cl\n95 mj\n13 bk\n30 pe\n65 yf\n54 nh\n80 tl\n11 xp\n15 fx\n71 on\n18 rk\n12 vf\n54 dc\n" +
		"59 ni\n99 su\n89 bl\n2 rm\n77 eu\n3 ur\n0 fu\n90 mh\n93 oc\n22 vu\n60 ax\n59 zk\n78 kl\n68 np\n29 nw\n" +
		"98 yj\n88 qh\n36 vo\n40 za\n75 fv\n10 mu\n53 sp\n39 nb\n42 qe\n31 cv\n48 bq\n27 is\n88 qo\n27 un\n22 hp\n" +
		"42 ui\n35 jb\n46 nm\n15 ed\n56 rn\n11 tn\n64 ro\n29 yd\n6 rj\n69 va\n53 cb\n13 tp\n23 ck\n59 bl\n45 co\n" +
		"7 iw\n23 pv\n6 ec\n75 sp\n80 zm\n77 mk\n44 zo\n88 ha\n42 om\n41 po\n34 on\n5 le\n78 fm\n27 rv\n9 bk\n49 up\n" +
		"75 nj\n20 hf\n41 wp\n18 lv\n91 aq\n95 np\n35 hu\n12 km\n29 qh\n19 zt\n25 ez\n33 us\n53 bo\n76 ki\n61 jd\n" +
		"8 rp\n87 hb\n88 ok\n70 ay\n64 sn\n16 ty\n53 si\n18 om\n12 rs\n12 yv\n80 ft\n80 ll\n22 sj\n75 nc\n69 xx\n" +
		"54 qx\n16 rz\n67 mr\n74 wb\n17 wf\n64 rk\n49 hq\n6 jw\n74 fa\n8 gl\n7 lx\n40 oo\n29 yk\n65 nc\n37 vm\n" +
		"97 ir\n80 au\n2 ho\n30 xi\n91 rr\n9 jo\n89 zb\n51 dl\n98 vj\n81 cr\n59 dc\n7 rq\n17 sm\n64 ur\n36 de\n" +
		"37 rh\n15 yt\n59 ah\n89 at\n82 yw\n73 ww\n34 ar\n46 xa\n92 fv\n33 oq\n88 ol\n85 cn\n29 sz\n44 cg\n33 no\n" +
		"30 bv\n97 vs\n0 ou\n18 cg\n33 pm\n34 wn\n12 vd\n31 el\n4 ah\n69 ke\n98 lo\n69 hw\n4 ks\n23 ca\n72 yb\n" +
		"72 ie\n57 qt\n66 sj\n74 vu\n45 hn\n76 gf\n26 oy\n92 kv\n53 ji\n66 sn\n84 pu\n60 kq\n37 ul\n54 gi\n93 gv\n" +
		"12 pl\n6 by\n90 ux\n30 gj\n38 gl\n93 ja\n87 fu\n36 dk\n76 zz\n73 db\n79 rd\n11 qi\n63 oq\n22 fu\n57 wn\n" +
		"99 cs\n99 hy\n69 wo\n20 zo\n69 bv\n43 di\n38 iv\n32 mn\n57 dv\n43 qk\n9 op\n71 zd\n34 ry\n67 ft\n25 nq\n" +
		"52 fe\n0 pq\n43 lf\n80 bx\n91 df\n59 ku\n91 vp\n31 so\n60 hc\n4 qd\n30 jp\n27 cz\n25 zh\n39 kj\n18 jh\n" +
		"8 ot\n16 fa\n76 cg\n17 wv\n38 yw\n45 ei\n11 ek\n48 rh\n83 bx\n62 ev\n36 cw\n7 iz\n95 jk\n61 ku\n89 gc\n" +
		"17 jy\n52 qs\n43 ud\n26 bq\n21 bq\n66 uc\n4 ou\n20 mj\n3 kx\n55 vo\n42 my\n77 em\n95 uu\n4 ew\n50 sf\n" +
		"31 og\n35 bz\n67 ci\n73 gj\n2 sd\n79 do\n5 vj\n85 hq\n47 ma\n6 uy\n44 ee\n58 ny\n63 wb\n48 io\n90 qo\n" +
		"76 yh\n53 kn\n8 cx\n90 yk\n2 wv\n87 ht\n67 rk\n4 mx\n1 mc\n41 kj\n50 wi\n81 ov\n91 xo\n52 ac\n26 jt\n" +
		"40 wf\n63 od\n91 sg\n87 jt\n75 ww\n61 fm\n16 dr\n71 mp\n50 mh\n46 pu\n69 dx\n17 oo\n53 mg\n3 sd\n74 uo\n" +
		"0 qy\n39 to\n89 hb\n6 xm\n11 me\n47 lz\n6 hl\n43 on\n56 bd\n71 ig\n90 hh\n99 jn\n12 ll\n80 rv\n7 cc\n" +
		"41 ol\n91 yk\n92 ms\n62 bk\n33 iy\n15 hp\n52 kw\n72 zo\n20 rl\n57 yy\n62 br\n34 nt\n29 au\n28 da\n3 ot\n" +
		"50 oo\n44 bt\n97 fe\n45 vl\n86 kw\n97 pv\n9 vu\n23 rj\n83 ee\n69 ty\n89 cv\n11 mq\n39 pq\n93 lx\n25 oj\n" +
		"18 gh\n71 ey\n40 xz\n55 xd\n64 bl\n21 bo\n15 gf\n98 sj\n10 oe\n2 dw\n28 rv\n31 lv\n40 yg\n72 cp\n63 dn\n" +
		"45 eu\n64 sc\n63 ib\n87 ye\n43 kd\n53 jj\n41 vs\n85 kh\n39 ki\n46 zq\n32 py\n54 nn\n61 cc\n35 iy\n52 yf\n" +
		"27 ya\n33 iv\n34 jx\n71 wx\n29 hj\n85 gm\n91 hc\n43 bb\n57 oq\n91 bw\n57 re\n26 hc\n13 no\n32 nm\n37 un\n" +
		"44 qa\n64 jx\n99 cx\n29 wj\n37 zk\n69 ur\n61 gc\n91 wn\n33 pz\n11 el\n73 ad\n93 zk\n66 kl\n54 ky\n32 jv\n" +
		"0 uz\n10 la\n97 kk\n27 co\n96 nh\n75 ze\n11 ud\n14 iv\n34 co\n80 zg\n21 qh\n58 ww\n6 cs\n61 kb\n74 mf\n" +
		"5 zh\n69 wq\n62 en\n84 ad\n17 qv\n63 xu\n93 zh\n2 cd\n99 nr\n2 hr\n62 vo\n5 rt\n52 go\n54 lj\n57 bd\n" +
		"0 eb\n19 qs\n53 im\n72 fu\n28 qi\n77 gd\n28 qy\n18 ej\n7 ja\n77 bd\n48 on\n76 pi\n93 wr\n94 im\n11 xz\n" +
		"70 bm\n53 di\n36 zx\n15 xl\n18 es\n1 yu\n49 xj\n16 rm\n82 vk\n80 py\n94 by\n2 qo\n40 ru\n53 rl\n74 gv\n" +
		"30 so\n52 ps\n82 gu\n53 rs\n35 uj\n41 os\n39 qi\n61 ed\n34 am\n89 ia\n65 ho\n90 jp\n62 vn\n91 dx\n60 ud\n" +
		"69 up\n33 qb\n6 pl\n16 nt\n29 be\n7 lx\n29 co\n5 dq\n89 hz\n26 fx\n21 os\n25 pc\n25 aj\n8 vr\n56 vh\n18 is\n" +
		"60 va\n10 dt\n4 aq\n62 pz\n27 gr\n25 md\n87 gu\n75 zr\n68 sh\n30 if\n95 fj\n33 pj\n74 xl\n11 nn\n2 tr\n" +
		"82 wk\n69 ie\n25 ly\n61 ne\n92 dd\n47 jl\n46 ys\n5 er\n7 de\n70 hs\n59 lh\n46 xu\n22 vx\n21 jn\n48 ek\n" +
		"48 op\n69 hm\n45 qi\n92 ut\n69 ta\n97 wf\n84 me\n96 wt\n90 ds\n90 rg\n22 zi\n32 oj\n10 im\n15 yl\n60 hu\n" +
		"56 gw\n83 kc\n36 yv\n23 bl\n77 jn\n7 gy\n77 ug\n58 fz\n44 pi\n92 yb\n11 os\n92 kl\n1 no\n85 vy\n75 zd\n" +
		"37 hj\n45 yp\n54 ep\n52 hi\n95 aq\n41 ls\n62 wc\n91 la\n56 xp\n31 cx\n32 lt\n2 zv\n12 nz\n44 wu\n51 ul\n" +
		"66 bd\n29 lv\n92 sd\n24 av\n77 tc\n4 ai\n74 hc\n59 ts\n87 hz\n46 ng\n60 vn\n81 gz\n58 yc\n15 mh\n74 yj\n" +
		"57 iy\n56 qf\n89 fb\n80 is\n73 za\n85 it\n81 vf\n74 pi\n88 vd\n23 xg\n40 ts\n27 xa\n67 zh\n30 zf\n48 qk\n" +
		"54 sf\n9 un\n31 pz\n4 vm\n88 py\n80 hf\n77 ls\n78 uw\n73 rt\n70 or\n33 te\n13 nn\n75 ug\n0 fo\n42 du\n" +
		"55 uz\n68 fc\n87 vn\n13 mr\n98 xo\n80 fv\n97 bk\n92 hq\n46 tz\n41 ho\n79 ls\n38 ft\n19 fr\n8 dk\n37 rw\n" +
		"15 kp\n30 lg\n42 ut\n78 ht\n70 bu\n22 yw\n36 io\n35 fk\n52 tw\n78 bq\n56 wx\n7 cr\n92 zh\n4 ux\n24 cq\n" +
		"33 lw\n38 pf\n88 go\n49 mb\n46 rk\n65 sh\n75 jg\n86 ng\n14 qd\n39 wd\n84 sp\n40 zo\n2 ev\n68 sj\n46 hk\n" +
		"13 li\n16 rf\n16 gk\n84 rn\n72 ul\n7 zp\n31 cg\n90 zl\n63 pk\n69 qi\n95 bz\n24 py\n69 eu\n31 wo\n4 me\n" +
		"85 jy\n92 ga\n1 wf\n99 xx\n90 dz\n63 fn\n11 xu\n89 hd\n61 tl\n89 ql\n24 up\n92 cj\n93 hf\n57 bw\n82 fj\n" +
		"65 gs\n40 hv\n54 kv\n47 nv\n45 tz\n17 we\n20 wj\n7 ll"

	reader := bufio.NewReader(strings.NewReader(systemIn))

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	expected := "- - - - - - - - - pq qy uz eb fo - - - - - mc yu no wf - - - - - sd wv dw cd hr qo tr zv ev - - - " +
		"kx sd ot - - - - - - qd ou ew mx aq ai vm ux me - - - vj zh rt dq er - - - - - - - - uy xm hl cs pl - - - " +
		"- - iz cc ja lx de gy cr zp ll - - - ot cx vr dk - - - - op vu un - - - - oe la dt im - - - - - - ek me mq" +
		" el ud xz nn os xu - - - - - - - - ll nz - - - - - - no nn mr li - - iv qd - - - - - - - hp gf xl yl mh kp" +
		" - - - - - fa dr rm nt rf gk - - - - - - wv jy oo qv we - - - - - - jh gh ej es is - - - - - - qs fr - - -" +
		" - - mj rl wj - bq bo qh os jn - - - - - - - vx zi yw - - - - - rj bl xg - - - - - - - - av cq py up - - -" +
		" - nq zh oj pc aj md ly - - - - bq jt hc fx - - - - - - cz ya co gr xa - - - - da rv qi qy - - - - - - - -" +
		" - au hj wj be co lv - - - - - - - jp so if zf lg - - - - so og lv cx pz cg wo - - mn py nm jv oj lt - - -" +
		" - - - - iy iv pz qb pj te lw - - - - - ry nt jx co am - - - - - - bz iy uj fk - - - - cw zx yv io - - - -" +
		" - un zk hj rw - - - iv yw ft pf - - - kj to pq ki qi wd - - - - - - wf xz yg ru ts zo hv - - - - - - - - " +
		"kj ol vs os ls ho - - - - - - - my du ut - - - di qk lf ud on kd bb - - - - ee bt qa pi wu - - - - - - ei " +
		"vl eu qi yp tz - - - - - - - pu zq ys xu ng tz rk hk - - - ma lz jl nv - - - - rh io on ek op qk - - - - -" +
		" xj mb - - sf wi mh oo - - - ul - fe qs ac kw yf go ps hi tw - - - - - - kn mg jj im di rl rs - - - - - - " +
		"- - - nn ky lj ep sf kv - - vo xd uz - - - bd vh gw xp qf wx - - - - - - - - dv yy oq re bd iy bw - - - - " +
		"ny ww fz yc - - - - - - - - - ku lh ts - - - - - hc ud va hu vn - - - - ku fm cc gc kb ed ne tl - - - - - " +
		"- ev bk br en vo vn pz wc - - wb od dn ib xu pk fn - - - - bl sc jx - - - - ho sh gs - - - uc kl bd - - - " +
		"- ft ci rk zh - - - sh fc sj - - - - - - - - - bv dx ty ur wq up ie hm ta qi eu - - - - bm hs or bu - - - " +
		"zd mp ig ey wx - - - - - zo cp fu ul - - - - - - - - - - gj ad za rt - - - - - - uo mf gv xl hc yj pi - - " +
		"- - - - - - ww ze zr zd ug jg - - - - - - cg yh pi - - - em gd bd jn ug tc ls - - - - - - - uw ht bq - - -" +
		" do ls - - - - - - - bx rv zg py is hf fv - - - - ov gz vf - - - vk gu wk fj - - - - bx ee kc - - - - - ad" +
		" me sp rn - - - - - - - hq kh gm vy it jy - - - - kw ng - - - - - ht jt ye gu hz vn - - - - - - - - - vd p" +
		"y go - - - - gc hb cv ia hz fb hd ql - - qo yk hh jp ds rg zl dz - - - - - - - - df vp xo sg yk hc bw wn d" +
		"x la - - - - ms dd ut yb kl sd hq zh ga cj - - - - lx zk zh wr hf - - - - im by - - - - - jk uu fj aq bz -" +
		" - nh wt - - - fe pv kk wf bk - - - - - sj xo - - - - - - jn cx nr xx \n"

	countSort(stdout, readInput(reader, n))
	result := readTestFile()

	if result != expected {
		t.Errorf("Full counting sort hidden case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// readTestFile is responsible of reading the output of the program written in the given writer.
func readTestFile() string {

	text := ""

	file, err := os.Open(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	defer func(stdout *os.File) {
		err := stdout.Close()
		if err != nil {
			fmt.Println("Error reading output path file.")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = text + scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return text
}
