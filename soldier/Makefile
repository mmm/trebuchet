default: sorcerer.pb.go

sorcerer.pb.go : ../sorcerer/sorcerer.proto
	@echo "---" $@ "---"
	protoc $< -I ../sorcerer --go_out=plugins=grpc:.

clean:
	@echo "---" $@ "---"
	rm -f *.pb.go
