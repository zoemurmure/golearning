module github.com/zoemurmure/helloworld

go 1.17

require (
	github.com/zoemurmure/calculator v0.0.0
	github.com/zoemurmure/geometry v0.0.0
	github.com/zoemurmure/store v0.0.0
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace github.com/zoemurmure/calculator => ../calculator
replace github.com/zoemurmure/geometry => ../geometry
replace github.com/zoemurmure/store => ../store
