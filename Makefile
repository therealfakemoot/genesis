image:
	docker build . -t therealfakemoot/genesis:latest
demo: clean
	docker run -d --name=genesis -p 8888:8888 therealfakemoot/genesis:latest
clean:
	docker stop genesis
	docker rm genesis
