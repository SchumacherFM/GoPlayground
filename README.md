# Go Playground

Various weird code examples.

## Benchmark Results by package `db`

```sql 
CREATE TABLE `catalog_product_entity_varchar` (
  `value_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Value ID',
  `entity_type_id` int(10) unsigned NOT NULL COMMENT 'Entity Type ID',
  `attribute_id` smallint(5) unsigned NOT NULL COMMENT 'Attribute ID',
  `store_id` smallint(5) unsigned NOT NULL COMMENT 'Store ID',
  `entity_id` int(10) unsigned NOT NULL COMMENT 'Entity ID',
  `value` varchar(255) DEFAULT NULL COMMENT 'Value',
  PRIMARY KEY (`value_id`),
  UNIQUE KEY `UNQ_CAT_PRD_ENTT_VCHR_ENTT_ID_ATTR_ID_STORE_ID` (`entity_id`,`attribute_id`,`store_id`),
  KEY `IDX_CATALOG_PRODUCT_ENTITY_VARCHAR_ATTRIBUTE_ID` (`attribute_id`),
  KEY `IDX_CATALOG_PRODUCT_ENTITY_VARCHAR_STORE_ID` (`store_id`),
  KEY `IDX_CATALOG_PRODUCT_ENTITY_VARCHAR_ENTITY_ID` (`entity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

1773 rows, example data:

```
1	10	96	0	16	Nokia 2610 Phone
2	10	102	0	16	20
3	10	481	0	16	nokia-2610-phone
4	10	103	0	16	NULL
5	10	105	0	16	Offering advanced media and calling bla ...
6	10	106	0	16	/n/o/nokia-2610-phone-2.jpg
7	10	109	0	16	/n/o/nokia-2610-phone-2.jpg
8	10	493	0	16	/n/o/nokia-2610-phone-2.jpg
9	10	96	0	17	BlackBerry 8100 Pearl
```

```
$ go test -v -run=XX -bench=. -benchmem -count=1 .
goos: darwin
goarch: amd64
pkg: github.com/SchumacherFM/GoPlayground/db
BenchmarkMapStringScan-4   	           1000	   2111936 ns/op	  422529 B/op	   20911 allocs/op
BenchmarkStrStrScan-4      	           1000	   2105879 ns/op	  422395 B/op	   20910 allocs/op
BenchmarkRowMapString-4    	            500	   3349476 ns/op	 1528189 B/op	   36858 allocs/op
BenchmarkSQLx/append-4     	            500	   3934703 ns/op	  535713 B/op	   22676 allocs/op
BenchmarkSQLx/select-4     	            300	   4190569 ns/op	  593669 B/op	   24450 allocs/op
BenchmarkCSFWdbr/convertAssign-4        500	   3129286 ns/op	  535301 B/op	   22674 allocs/op
BenchmarkCSFWdbr/special-convert-4     1000	   2039783 ns/op	  510957 B/op	   13809 allocs/op <-- !!!!
BenchmarkSqlStruct-4                    300	   5366031 ns/op	  961026 B/op	   29766 allocs/op
BenchmarkSqlBoiler/all-4                300	   4257519 ns/op	  791203 B/op	   26223 allocs/op
BenchmarkKnq_xo/all-4                   500	   3322086 ns/op	  535404 B/op	   22674 allocs/op
PASS
ok  	github.com/SchumacherFM/GoPlayground/db	20.731s```
