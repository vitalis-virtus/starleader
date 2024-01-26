Build js game file:
`$Env:GOOS = 'js'
$Env:GOARCH = 'wasm'
go build -o game.wasm github.com/vitalis-virtus/starleader    
Remove-Item Env:GOOS
Remove-Item Env:GOARCH`

`$goroot = go env GOROOT
cp $goroot\misc\wasm\wasm_exec.js .`