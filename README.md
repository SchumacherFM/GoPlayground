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

For more details see the `bench_result.txt` in this repo.

```
$ go test -v -run=XX -bench=. -benchmem -count=1 .
$ benchstat bench_result.txt
name                        time/op
MapStringScan-4             1.76ms ± 4%
StrStrScan-4                1.50ms ± 1%
RowMapString-4              2.95ms ± 0%
SQLx/append-4               3.55ms ± 0%
SQLx/select-4               3.70ms ± 0%
CSFWdbr/convertAssign-4     2.72ms ± 0%
CSFWdbr/convert_byteconv-4  1.70ms ± 1%
CSFWdbr/convert_stdlib-4    1.88ms ± 1%
SqlStruct-4                 4.96ms ± 0%
SqlBoiler/all-4             3.85ms ± 0%
Knq_xo/all-4                2.84ms ± 0%
Reform/all-4                2.99ms ± 0%

name                        alloc/op
MapStringScan-4              423kB ± 0%
StrStrScan-4                 422kB ± 0%
RowMapString-4              1.53MB ± 0%
SQLx/append-4                536kB ± 0%
SQLx/select-4                593kB ± 0%
CSFWdbr/convertAssign-4      535kB ± 0%
CSFWdbr/convert_byteconv-4   511kB ± 0%
CSFWdbr/convert_stdlib-4     535kB ± 0%
SqlStruct-4                  961kB ± 0%
SqlBoiler/all-4              791kB ± 0%
Knq_xo/all-4                 535kB ± 0%
Reform/all-4                 829kB ± 0%

name                        allocs/op
MapStringScan-4              20.9k ± 0%
StrStrScan-4                 20.9k ± 0%
RowMapString-4               36.9k ± 0%
SQLx/append-4                22.7k ± 0%
SQLx/select-4                24.4k ± 0%
CSFWdbr/convertAssign-4      22.7k ± 0%
CSFWdbr/convert_byteconv-4   13.8k ± 0%
CSFWdbr/convert_stdlib-4     22.7k ± 0%
SqlStruct-4                  29.8k ± 0%
SqlBoiler/all-4              26.2k ± 0%
Knq_xo/all-4                 22.7k ± 0%
Reform/all-4                 24.5k ± 0%
```
