import os
import shutil
import time
os.environ["CGO_ENABLED"] = "0"
os.environ["GOARCH"] = "amd64"

# darwin, windows, linux
os.environ["GOOS"] = "linux"


has = os.popen("git log -n1 --format=format:%H")

LDFLAGS = f"-X 'main.buildTime={time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(time.time()))}'" \
          f"-X 'main.version=1.1.{time.strftime('%Y%m%d', time.localtime(time.time()))}'" \
          f"-X 'main.commit={has.read()}'"

os.system(f"go build -ldflags \"-w -s {LDFLAGS}\"")

if not os.path.isdir("bin"):
    os.makedirs("bin")

if os.path.isfile("server"):
    if os.path.isfile("bin/server"):
        os.remove("bin/server")
    shutil.move("server", "bin/server")
    shutil.copy("server.toml", "bin/server.toml")
    shutil.copytree("sql", "bin/sql")


