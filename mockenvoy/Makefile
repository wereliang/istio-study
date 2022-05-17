BuildVersion = v2
Registry = 10.80.73.169:5000
LDFlags = "-X 'main.BuildVersion=$(BuildVersion)'"
Image = $(Registry)/xxx/mockenvoy:$(BuildVersion)

all:
	docker build --build-arg LDFLAGS=$(LDFlags) -t $(Image) .

push:
	docker push $(Image)
