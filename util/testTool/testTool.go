package testTool

import (
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	Equal int = 1 + iota
	NotEqual
	AlmostEqual
	NotAlmostEqual
	Resemble
	NotResemble
	PointTo
	NotPointTo
	BeNil
	NotBeNil
	BeTrue
	BeFalse
	BeZeroValue
	NotBeZeroValue

	BeGreaterThan
	BeGreaterThanOrEqualTo
	BeLessThan
	BeLessThanOrEqualTo
	BeBetween
	NotBeBetween
	BeBetweenOrEqual
	NotBeBetweenOrEqual
	Contain
	NotContain
	ContainKey
	NotContainKey
	BeIn
	NotBeIn
	BeEmpty
	NotBeEmpty
	HaveLength
	StartWith
	NotStartWith
	EndWith
	NotEndWith
	BeBlank
	NotBeBlank
	ContainSubstring
	NotContainSubstring
	Panic
	NotPanic
	PanicWith
	NotPanicWith
	HaveSameTypeAs
	NotHaveSameTypeAs
	Implement
	NotImplement
	HappenBefore
	HappenOnOrBefore
	HappenAfter
	HappenOnOrAfter
	HappenBetween
	HappenOnOrBetween
	NotHappenOnOrBetween
	HappenWithin
	NotHappenWithin
	BeChronological

	BeError
)

var methodMap = map[int]func(actual interface{}, expected ...interface{}) string{

	Equal:          should.Equal,
	NotEqual:       should.NotEqual,
	AlmostEqual:    should.AlmostEqual,
	NotAlmostEqual: should.NotAlmostEqual,
	Resemble:       should.Resemble,
	NotResemble:    should.NotResemble,
	PointTo:        should.PointTo,
	NotPointTo:     should.NotPointTo,
	BeNil:          should.BeNil,
	NotBeNil:       should.NotBeNil,
	BeTrue:         should.BeTrue,
	BeFalse:        should.BeFalse,
	BeZeroValue:    should.BeZeroValue,
	NotBeZeroValue: should.NotBeZeroValue,

	BeGreaterThan:          should.BeGreaterThan,
	BeGreaterThanOrEqualTo: should.BeGreaterThanOrEqualTo,
	BeLessThan:             should.BeLessThan,
	BeLessThanOrEqualTo:    should.BeLessThanOrEqualTo,
	BeBetween:              should.BeBetween,
	NotBeBetween:           should.NotBeBetween,
	BeBetweenOrEqual:       should.BeBetweenOrEqual,
	NotBeBetweenOrEqual:    should.NotBeBetweenOrEqual,
	Contain:                should.Contain,
	NotContain:             should.NotContain,
	ContainKey:             should.ContainKey,
	NotContainKey:          should.NotContainKey,
	BeIn:                   should.BeIn,
	NotBeIn:                should.NotBeIn,
	BeEmpty:                should.BeEmpty,
	NotBeEmpty:             should.NotBeEmpty,
	HaveLength:             should.HaveLength,
	StartWith:              should.StartWith,
	NotStartWith:           should.NotStartWith,
	EndWith:                should.EndWith,
	NotEndWith:             should.NotEndWith,
	BeBlank:                should.BeBlank,
	NotBeBlank:             should.NotBeBlank,
	ContainSubstring:       should.ContainSubstring,
	NotContainSubstring:    should.NotContainSubstring,
	Panic:                  should.Panic,
	NotPanic:               should.NotPanic,
	PanicWith:              should.PanicWith,
	NotPanicWith:           should.NotPanicWith,
	HaveSameTypeAs:         should.HaveSameTypeAs,
	NotHaveSameTypeAs:      should.NotHaveSameTypeAs,
	Implement:              should.Implement,
	NotImplement:           should.NotImplement,
	HappenBefore:           should.HappenBefore,
	HappenOnOrBefore:       should.HappenOnOrBefore,
	HappenAfter:            should.HappenAfter,
	HappenOnOrAfter:        should.HappenOnOrAfter,
	HappenBetween:          should.HappenBetween,
	HappenOnOrBetween:      should.HappenOnOrBetween,
	NotHappenOnOrBetween:   should.NotHappenOnOrBetween,
	HappenWithin:           should.HappenWithin,
	NotHappenWithin:        should.NotHappenWithin,
	BeChronological:        should.BeChronological,

	BeError: should.BeError,
}

func Test(t *testing.T, actual interface{}, assertion int, expected ...interface{}) {

	Convey("test", t, func() {
		So(actual, methodMap[assertion], expected...)
	})
}

func EqualTest(t *testing.T, actual interface{}, expected ...interface{}) {
	Convey("testassssss", t, func() {
		So(actual, methodMap[Equal], expected...)
	})
}

func NotEqualTest(t *testing.T, actual interface{}, expected ...interface{}) {
	Convey("test", t, func() {
		So(actual, methodMap[NotEqual], expected...)
	})
}
