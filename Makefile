image:
	docker build . -t therealfakemoot/genesis:latest
demo: image
	docker run -d --name=genesis -p 8080:8080 therealfakemoot/genesis:latest
clean:
	docker stop genesis
	docker rm genesis
