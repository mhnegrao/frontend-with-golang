# Go Gin Svelte SPA Embed Example

```sh
git clone https://git.web8.is/example/go-gin-embed-svelte-example.git
cd go-gin-embed-svelte-example

cd frontend
npm i # install dependency
npm i -g vite # install vite globally
vite build # build svelte into static output

cd ..
go mod tidy
go run .
```
