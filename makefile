.PHONY: b f bb bf fi back-end front-end

# b = back-end
back-end:
	@go -C ./back-end run ./cmd

b: back-end

# fi = front-install
front-end_install_dependencies:
	@bun install --cwd ./front-end

fi: front-end_install_dependencies

front-end: front-end_install_dependencies
	@bun --cwd ./front-end dev

f: front-end

build_back-end_docker_image:
	@podman build --network=host -t back ./back-end

bb: build_back-end_docker_image

# bf = build front-end docker image
build_front-end_docker_image:
	@podman build --network=host -t front ./front-end

bf: build_front-end_docker_image

run: build_back-end_docker_image build_front-end_docker_image
	@podman-compose up -d
