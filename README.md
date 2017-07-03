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
4	10	103	0	16	Nokia 2610
5	10	105	0	16	Offering advanced media and calling bla ...
6	10	106	0	16	/n/o/nokia-2610-phone-2.jpg
7	10	109	0	16	/n/o/nokia-2610-phone-2.jpg
8	10	493	0	16	/n/o/nokia-2610-phone-2.jpg
9	10	96	0	17	BlackBerry 8100 Pearl
```

```
$ go test -v -run=XX -bench=. -benchmem .
goos: darwin
goarch: amd64
pkg: github.com/SchumacherFM/GoPlayground/db
BenchmarkMapStringScan-4   	    1000	   2127262 ns/op	  422528 B/op	   20912 allocs/op
BenchmarkStrStrScan-4      	    1000	   2339049 ns/op	  422402 B/op	   20911 allocs/op
BenchmarkRowMapString-4    	     500	   3352882 ns/op	 1528198 B/op	   36859 allocs/op
BenchmarkSQLx/append-4     	     300	   3949653 ns/op	  535795 B/op	   22677 allocs/op
BenchmarkSQLx/select-4     	     300	   4163452 ns/op	  593664 B/op	   24451 allocs/op
BenchmarkCSFWdbr-4         	     500	   3232287 ns/op	  563782 B/op	   24448 allocs/op
BenchmarkSqlStruct-4       	     300	   5388080 ns/op	  961058 B/op	   29767 allocs/op
BenchmarkSqlBoiler/all-4   	     300	   4287595 ns/op	  791223 B/op	   26225 allocs/op
PASS
ok  	github.com/SchumacherFM/GoPlayground/db	16.046s
```
