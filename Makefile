PROJECT=carpark

# These are the values we want to pass for Version and BuildTime
GITTAG=`git describe --tags`
COMMIT=`git rev-parse HEAD`

clean:
	@rm -rf orm/sql
	@rm -rf orm/model

gen: $(shell which db-orm)
	@mkdir -p orm/sql
	@mkdir -p orm/model
	db-orm sql -i orm/yaml -o orm/sql
	db-orm code -i orm/yaml -o orm/model

image: 
	docker build --build-arg VERSION=${GITTAG} --build-arg COMMIT=${COMMIT} .
