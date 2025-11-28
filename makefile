# b = back-end
b:

# fi = front-install
fi: 
	bun install --cwd ./front-end

f: fi
	bun --cwd ./front-end dev

# bb = build back-end docker image
bb: 
	docker build -t api ./back-end

# bf = build front-end docker image
bf: 
	docker build --network=host -t front ./front-end
	
run: bb bf
	docker-compose up
