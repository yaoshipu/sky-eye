
# 安装程序依赖
install:
	go get -u github.com/labstack/echo/...

# 构建后端服务
build:
	cd src; go install -v ./skyeye.paypal.com/...; cd ..

# 启动后端服务
dev:
	./bin/skyeye

# 安装配置文件
config:
	mkdir -p ~/.skyeye
	cp src/skyeye.paypal.com/skyeye/skyeye.conf ~/.skyeye
