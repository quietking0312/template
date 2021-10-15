import os
import time
os.environ["CGO_ENABLED"] = "0"
os.environ["GOARCH"] = "amd64"

# darwin, windows, linux
os.environ["GOOS"] = "linux"


has = os.popen("git log -n1 --format=format:%H")

LDFLAGS = f"-X 'main.buildTime={time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(time.time()))}'" \
          f"-X 'main.version=1.1.{time.strftime('%Y%m%d', time.localtime(time.time()))}'" \
          f"-X 'main.commit={has.read()}'"

os.system(f"go build -ldflags \"{LDFLAGS}\"")

