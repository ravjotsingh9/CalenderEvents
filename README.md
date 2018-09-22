### Event Conflict Finder
Utility to read event json data from file and return conflicting events to stdout.

#### Build
```
make deps
```
```
make
```

#### Run
You can provide your own data in the format presented in `testData/data.txt`, Or you can use the given generator to generate some random data. Following steps will generate 10 data entries:
```
cd testData
make gendeps
make testDataGenerator
./generateData 10
cd ..
```
asTo run the conflict finder program, run as below:

```
./ecf_binary

------------ Overlapped Events ------------

EVENT NAME     | START TIME          | END TIME
enabled-sponge | 1994-01-22 19:55:01 | 2011-03-16 04:38:09
giving-calf    | 2008-08-18 00:43:36 | 2012-10-27 03:26:43
sacred-corgi   | 2011-01-28 01:01:41 | 2012-02-23 19:32:40
```

It will display conflicting events.


#### Known bugs
