你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# colorgo
`colorgo` is a wrapper to `go` command that colorizes output from `go build` and `go test`.
![screenshot 1](http://songgao.github.com/colorgo/images/screenshot1.png)

# Installation
```
go get -u github.com/songgao/colorgo
```

# Usage
```bash
colorgo build
```

# alias
`colorgo` changes nothing to sub-commands other than `go build`. So you can optionally define alias in your shell conf so that `go build` always prints colorized error message:

bash: `~/.bashrc`
```
alias go=colorgo
```

fish-shell: `~/.config/fish/config.fish`
```
alias go colorgo
```

# License
[BSD 3-Clause License](http://opensource.org/licenses/BSD-3-Clause)
