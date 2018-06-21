build:
	go build -i

test: build
	./facility-location data/ap10_2.txt -b 90963539.4763
	./facility-location data/ap10_4.txt -b 95079629.9069
	./facility-location data/ap10_6.txt -b 95161467.5800
	./facility-location data/ap10_8.txt -b 95161467.5800
	./facility-location data/ap20_2.txt -b 91507336.6084
	./facility-location data/ap20_4.txt -b 96673177.8591
	./facility-location data/ap20_6.txt -b 98181949.7142
	./facility-location data/ap20_8.txt -b 98181949.7142
	./facility-location data/ap30_2.txt -b 83576754.7126
	./facility-location data/ap30_4.txt -b 91209293.1037
	./facility-location data/ap30_6.txt -b 95053569.4196
	./facility-location data/ap30_8.txt -b 98520500.5833
	./facility-location data/ap40_2.txt -b 80537210.1476
	./facility-location data/ap40_4.txt -b 88042046.0930
	./facility-location data/ap40_6.txt -b 94523086.4474
	./facility-location data/ap40_8.txt -b 99180263.1029
	./facility-location data/ap50_2.txt -b 71261044.6135
	./facility-location data/ap50_4.txt -b 80325464.9026
	./facility-location data/ap50_6.txt -b 89389885.1918
	./facility-location data/ap50_8.txt -b 95205946.9629
	./facility-location data/ap60_2.txt -b 64790967.0386
	./facility-location data/ap60_4.txt -b 73074656.5752
	./facility-location data/ap60_6.txt -b 80673823.6599
	./facility-location data/ap60_8.txt -b 87285162.0911
	./facility-location data/ap70_2.txt -b 74451085.6010
	./facility-location data/ap70_4.txt -b 83272519.6325
	./facility-location data/ap70_6.txt -b 91741977.1993
	./facility-location data/ap70_8.txt -b 97109841.9802
	./facility-location data/ap80_2.txt -b 70713485.9938
	./facility-location data/ap80_4.txt -b 79704787.9527
	./facility-location data/ap80_6.txt -b 88418089.7530
	./facility-location data/ap80_8.txt -b 95798238.8757
	./facility-location data/ap90_2.txt -b 69223173.9238
	./facility-location data/ap90_4.txt -b 78931357.8060
	./facility-location data/ap90_6.txt -b 87012179.9499
	./facility-location data/ap90_8.txt -b 92780846.2542
	./facility-location data/ap100_2.txt -b 67584119.7648
	./facility-location data/ap100_4.txt -b 77545112.3194
	./facility-location data/ap100_6.txt -b 86371515.8647
	./facility-location data/ap100_8.txt -b 93184508.9692
	./facility-location data/ap100_8.txt -b 62774425.0872 -a 1
	./facility-location data/ap100_8.txt -b 62774425.0872 -a 5
	./facility-location data/ap100_8.txt -b 62774425.0872 -a 10
	./facility-location data/ap100_8.txt -b 62774425.0872 -a 15
	./facility-location data/ap100_8.txt -b 62774425.0872 -a 20
	./facility-location data/ap100_8.txt -b 62774425.0872 -a 30
	./facility-location data/ap100_8.txt -b 62774425.0872 -a 40
	./facility-location data/ap100_8.txt -b 62774425.0872 -a 50
