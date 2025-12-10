.PHONY: b f bb bf fi back-end front-end

back-end:
	@go -C ./back-end run ./cmd

# b = back-end
b: back-end

front-end_install_dependencies:
	@bun install --cwd ./front-end

# fi = front-install
fi: front-end_install_dependencies

front-end: front-end_install_dependencies
	@bun --cwd ./front-end dev

# f = front-end
f: front-end

build_back-end_docker_image:
	@podman build --network=host -t back ./back-end

# build back-end
bb: build_back-end_docker_image

# drb = docker run back-end
drb: build_back-end_docker_image
	@podman run -p 2306:2306 back

build_front-end_docker_image:
	@podman build --network=host -t front ./front-end

# bf = build front-end docker image
bf: build_front-end_docker_image

# drf = docker run front-end
drf: build_front-end_docker_image
	@podman run -p 2207:2207 front

run: build_back-end_docker_image build_front-end_docker_image
	@podman-compose up -d
