default: docker-images
apps = sorcerer soldier

protobufs:
	@echo "---" $@ "---"
	@for app in $(apps); do \
	  $(MAKE) -C $$app; \
	done

docker-images: protobufs
	docker-compose build

clean:
	@echo "---" $@ "---"
	docker-compose down
	@for app in $(apps); do \
	  $(MAKE) -C $$app $@; \
	done
