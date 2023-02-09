echo "[!] build for mac...."
export GOOS="darwin"
export GOARCH="amd64"
go build -v -a -o ./build/gdbc_darwin -ldflags '-s -w' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" &> /dev/null
echo "[+] build for mac done!!!"

echo "[!] build for linux-amd64...."
export GOOS="linux"
export GOARCH="amd64"
go build -v -a  -o ./build/gdbc_linux_amd64 -ldflags '-s -w' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" &> /dev/null
echo "[+] build for mac done!!!"

echo "[!] build for windows...."
export GOOS="windows"
export GOARCH="amd64"
go build -v -a -o ./build/gdbc_windows.exe -ldflags '-s -w' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" &> /dev/null
echo "[+] build for mac done!!!"

