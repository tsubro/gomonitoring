# Go Monitoring


> This is a Go Utility for logging metrics
> Idea of this project was to solve a problem of logging metrics in a periodic interval. eg: If your application keeps calling a huge number of external and you dont wan't to log for each api calls instead of that metric logging will useful to know for a given time interval how many api calls succeded and failed.
> I have also include a common context which will store information that needs to be printed with everytime actuals logs gets emitted (like Java MDC logging)
> Code is still in pre-mature state and I will be obliged if folks here can support me in this learning journey :)

Types of Aggregator available:
* Periodic Aggregator
* Workflow Aggregator (future plan)

Types of Log Emitter available:
* Console Emitter
* File Emitter


[![Build Status](http://img.shields.io/travis/badges/badgerbadgerbadger.svg?style=flat-square)](https://travis-ci.org/badges/badgerbadgerbadger) [![Dependency Status](http://img.shields.io/gemnasium/badges/badgerbadgerbadger.svg?style=flat-square)](https://gemnasium.com/badges/badgerbadgerbadger) [![Coverage Status](http://img.shields.io/coveralls/badges/badgerbadgerbadger.svg?style=flat-square)](https://coveralls.io/r/badges/badgerbadgerbadger) [![Code Climate](http://img.shields.io/codeclimate/github/badges/badgerbadgerbadger.svg?style=flat-square)](https://codeclimate.com/github/badges/badgerbadgerbadger) [![Github Issues](http://githubbadges.herokuapp.com/badges/badgerbadgerbadger/issues.svg?style=flat-square)](https://github.com/badges/badgerbadgerbadger/issues) [![Pending Pull-Requests](http://githubbadges.herokuapp.com/badges/badgerbadgerbadger/pulls.svg?style=flat-square)](https://github.com/badges/badgerbadgerbadger/pulls) [![Gem Version](http://img.shields.io/gem/v/badgerbadgerbadger.svg?style=flat-square)](https://rubygems.org/gems/badgerbadgerbadger) [![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org) [![Badges](http://img.shields.io/:badges-9/9-ff6799.svg?style=flat-square)](https://github.com/badges/badgerbadgerbadger)

---
## Sample Code
```
  imports(
    "github.com/tsubro/gomonitoring/monitoring"
    "github.com/tsubro/gomonitoring/emitters"
  )
```

```
func Test() {
	m := monitoring.GetInstance()                 //Get a singleton monitoring object

	m.SetAggregatorFrequency(1)                   //Value in hour
	m.SetEmitterFrequency(1)                      //Value in hour
	m.SetEmitter(emitters.ConsoleEmitterImpl{})   //Setting appropriate emitter
	m.Start()                                     //Start monitoring
}
```
---

## Contributing

> To get started...

### Step 1

- **Option 1**
    - 🍴 Fork this repo!

- **Option 2**
    - 👯 Clone this repo to your local machine 

### Step 2

- **HACK AWAY!** 🔨🔨🔨

### Step 3

- 🔃 Create a new pull request 
---

## License

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)

- **[MIT license](http://opensource.org/licenses/mit-license.php)**
