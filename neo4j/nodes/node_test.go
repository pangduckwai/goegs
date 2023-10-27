package nodes

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	d0 := []uint32{
		0x05000009,
		0x030899FF,
		0x03CDFFB0,
		0x02AABBCC,
	}
	d1 := []uint32{
		0x05000009,
		0x030899FF,
		0x03CDFFB0,
		0x02AABBCC,
	}

	sort.Slice(d0, func(i, j int) bool { return d0[i] < d0[j] })
	fmt.Printf("%x\n", d0)

	sort.Slice(d1, func(i, j int) bool { return Less(d1[i], 0, d1[j], 0) })
	fmt.Printf("%x\n", d1)
}

func TestCalc(t *testing.T) {
	v := 327680
	fmt.Printf("Division 65536: %v\nRight shift 16: %v\n", v/65536, v>>16)
}

/*
 72.98
913.67 (turn as int)
913.67 (turn as string)

Build tree
==========
MERGE (n:N {v:"TESTING-3", r:1, d:0}) SET n:T RETURN elementId(n)  -->  "4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:0"
MERGE (t:T {v: "TESTING-3"}) ON CREATE SET t:N, t.r=1, t.d=0 RETURN elementId(t) as id, t{.*} as prop
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:0" CREATE (c:N {r:1, d:67108864})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:0" CREATE (c:N {r:1, d:68157440})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:0" CREATE (c:N {r:1, d:69206016})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:1" CREATE (c:N {r:1, d:68157441})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:1" CREATE (c:N {r:1, d:69206017})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:2" CREATE (c:N {r:1, d:67108865})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:2" CREATE (c:N {r:1, d:69206017})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:3" CREATE (c:N {r:1, d:67108865})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:3" CREATE (c:N {r:1, d:68157441})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:4" CREATE (c:N {r:1, d:69206018})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:5" CREATE (c:N {r:1, d:68157442})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:6" CREATE (c:N {r:1, d:69206018})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:7" CREATE (c:N {r:1, d:67108866})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:8" CREATE (c:N {r:1, d:68157442})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;
MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:9" CREATE (c:N {r:1, d:67108866})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1;

MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:13" CREATE (c:N {r:1, d:135266304})<-[:C]-(p)
WITH c MATCH (c)<-[:C*]-(b)
WITH c, b,
CASE
	WHEN apoc.bitwise.op(b.d, "&", 7) = 0 THEN 1
	WHEN b.w IS NULL THEN null
	ELSE 0
END AS w,
CASE
	WHEN apoc.bitwise.op(c.d, "&", 7) = 0 THEN 1
	ELSE null
END AS v
SET b.r = b.r + 1, b.w = coalesce(b.w, 0) + w, c.w = v
RETURN elementId(c)

MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:13" AND NOT EXISTS (MATCH (:N {r:1, d:136314880})<-[:C]-(p))
CREATE (c:N {r:1, d:136314880})<-[:C]-(p)
WITH c MATCH (c)<-[:C*]-(b)
WITH c, b,
CASE
	WHEN apoc.bitwise.op(b.d, "&", 7) = 2 THEN 1
	WHEN b.w IS NULL THEN null
	ELSE 0
END AS w,
CASE
	WHEN apoc.bitwise.op(c.d, "&", 7) = 2 THEN 1
	ELSE null
END AS v
SET b.r = b.r + 1, b.w = coalesce(b.w, 0) + w, c.w = v
RETURN elementId(c)

MATCH (p:N) WHERE elementId(p)="4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:13" CREATE (c:N {r:1, d:134217728})<-[:C]-(p)
WITH c MATCH (c)<-[:C*]-(b)
WITH c, b,
CASE
	WHEN apoc.bitwise.op(b.d, "&", 7) = 1 THEN 1
	WHEN b.w IS NULL THEN null
	ELSE 0
END AS w,
CASE
	WHEN apoc.bitwise.op(c.d, "&", 7) = 1 THEN 1
	ELSE null
END AS v
SET b.r = b.r + 1, b.w = coalesce(b.w, 0) + w, c.w = v
RETURN elementId(c)
*/
// TestBuildNodes
func TestBuildNodes(t *testing.T) {
	r := Init("TESTING-3")
	fmt.Printf("└┬──── %v\n", r.R)

	fmt.Printf(" ├┬─── %v\n", New(r.R, 0, 1, []uint8{0}, 0))
	fmt.Printf(" │├┬── %v\n", New(r.R, 1, 1, []uint8{1}, 0))
	fmt.Printf(" ││└── %v\n", New(r.R, 2, 1, []uint8{2}, 0))
	fmt.Printf(" │└┬── %v\n", New(r.R, 1, 1, []uint8{2}, 0))
	fmt.Printf(" │ └── %v\n", New(r.R, 2, 1, []uint8{1}, 0))

	fmt.Printf(" ├┬─── %v\n", New(r.R, 0, 1, []uint8{1}, 0))
	fmt.Printf(" │├┬── %v\n", New(r.R, 1, 1, []uint8{0}, 0))
	fmt.Printf(" ││└── %v\n", New(r.R, 2, 1, []uint8{2}, 0))
	fmt.Printf(" │└┬── %v\n", New(r.R, 1, 1, []uint8{2}, 0))
	fmt.Printf(" │ └── %v\n", New(r.R, 2, 1, []uint8{0}, 0))

	fmt.Printf(" └┬─── %v\n", New(r.R, 0, 1, []uint8{2}, 0))
	fmt.Printf("  ├┬── %v\n", New(r.R, 1, 1, []uint8{0}, 0))
	fmt.Printf("  │└── %v\n", New(r.R, 2, 1, []uint8{1}, 0))
	fmt.Printf("  └┬── %v\n", New(r.R, 1, 1, []uint8{1}, 0))
	fmt.Printf("   └── %v\n", New(r.R, 2, 1, []uint8{0}, 0))

	fmt.Printf("       %v\n", New(r.R, 0, 2, []uint8{1}, 0))
	fmt.Printf("       %v\n", New(r.R, 0, 2, []uint8{2}, 0))
	fmt.Printf("       %v\n", New(r.R, 0, 2, []uint8{0}, 0))
}

/*
└┬──── [          0/          1]0{0} 0, 0, 0,0,0,   0|   0≡ TEMP: {r:1, d:0, v:0}
 ├┬─── [          0/          0]0{1} 0, 0, 0,0,0,   0|   0± TEMP: {r:1, d:67108864, v:0}
 │├┬── [          0/          0]1{1} 1, 0, 0,0,0,   0|   0± TEMP: {r:1, d:68157441, v:0}
 ││└── [          0/          0]2{1} 2, 0, 0,0,0,   0|   0± TEMP: {r:1, d:69206018, v:0}
 │└┬── [          0/          0]1{1} 2, 0, 0,0,0,   0|   0± TEMP: {r:1, d:69206017, v:0}
 │ └── [          0/          0]2{1} 1, 0, 0,0,0,   0|   0± TEMP: {r:1, d:68157442, v:0}
 ├┬─── [          0/          0]0{1} 1, 0, 0,0,0,   0|   0± TEMP: {r:1, d:68157440, v:0}
 │├┬── [          0/          0]1{1} 0, 0, 0,0,0,   0|   0± TEMP: {r:1, d:67108865, v:0}
 ││└── [          0/          0]2{1} 2, 0, 0,0,0,   0|   0± TEMP: {r:1, d:69206018, v:0}
 │└┬── [          0/          0]1{1} 2, 0, 0,0,0,   0|   0± TEMP: {r:1, d:69206017, v:0}
 │ └── [          0/          0]2{1} 0, 0, 0,0,0,   0|   0± TEMP: {r:1, d:67108866, v:0}
 └┬─── [          0/          0]0{1} 2, 0, 0,0,0,   0|   0± TEMP: {r:1, d:69206016, v:0}
  ├┬── [          0/          0]1{1} 0, 0, 0,0,0,   0|   0± TEMP: {r:1, d:67108865, v:0}
  │└── [          0/          0]2{1} 1, 0, 0,0,0,   0|   0± TEMP: {r:1, d:68157442, v:0}
  └┬── [          0/          0]1{1} 1, 0, 0,0,0,   0|   0± TEMP: {r:1, d:68157441, v:0}
   └── [          0/          0]2{1} 0, 0, 0,0,0,   0|   0± TEMP: {r:1, d:67108866, v:0}
       [          0/          0]0{2} 1, 0, 0,0,0,   0|   0± TEMP: {r:1, d:135266304, v:0}
       [          0/          0]0{2} 2, 0, 0,0,0,   0|   0± TEMP: {r:1, d:136314880, v:0}
       [          0/          0]0{2} 0, 0, 0,0,0,   0|   0± TEMP: {r:1, d:134217728, v:0}
*/
