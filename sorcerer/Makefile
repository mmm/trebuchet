default: sorcerer.pb.go

%.pb.go : %.proto
	@echo "---" $@ "---"
	protoc $< --go_out=plugins=grpc:.

clean:
	@echo "---" $@ "---"
	rm -f *.pb.go
