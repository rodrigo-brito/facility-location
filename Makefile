build:
	go build -i

test: build
	./facility-location data/ap10_2.txt -b 90963539.4763
	./facility-location data/ap20_2.txt -b 91507336.6084
	./facility-location data/ap30_2.txt -b 83576754.7126
	./facility-location data/ap40_2.txt -b 80537210.1476
	./facility-location data/ap50_2.txt -b 71261044.6135
	./facility-location data/ap60_2.txt -b 64790967.0386
	./facility-location data/ap70_2.txt -b 74451085.601
	./facility-location data/ap80_2.txt -b 70713485.9938
	./facility-location data/ap90_2.txt -b 69223173.9238
	./facility-location data/ap100_2.txt -b 67584119.7648