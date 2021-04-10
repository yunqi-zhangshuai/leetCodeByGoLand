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

var methodMap = []func(actual interface{}, expected ...interface{}) string{
	Equal: should.Equal,
	should.NotEqual,
	should.AlmostEqual,
	should.NotAlmostEqual,
	should.Resemble,
	should.NotResemble,
	should.PointTo,
	should.NotPointTo,
	should.BeNil,
	should.NotBeNil,
	should.BeTrue,
	should.BeFalse,
	should.BeZeroValue,
	should.NotBeZeroValue,

	should.BeGreaterThan,
	should.BeGreaterThanOrEqualTo,
	should.BeLessThan,
	should.BeLessThanOrEqualTo,
	should.BeBetween,
	should.NotBeBetween,
	should.BeBetweenOrEqual,
	should.NotBeBetweenOrEqual,
	should.Contain,
	should.NotContain,
	should.ContainKey,
	should.NotContainKey,
	should.BeIn,
	should.NotBeIn,
	should.BeEmpty,
	should.NotBeEmpty,
	should.HaveLength,
	should.StartWith,
	should.NotStartWith,
	should.EndWith,
	should.NotEndWith,
	should.BeBlank,
	should.NotBeBlank,
	should.ContainSubstring,
	should.NotContainSubstring,
	should.Panic,
	should.NotPanic,
	should.PanicWith,
	should.NotPanicWith,
	should.HaveSameTypeAs,
	should.NotHaveSameTypeAs,
	should.Implement,
	should.NotImplement,
	should.HappenBefore,
	should.HappenOnOrBefore,
	should.HappenAfter,
	should.HappenOnOrAfter,
	should.HappenBetween,
	should.HappenOnOrBetween,
	should.NotHappenOnOrBetween,
	should.HappenWithin,
	should.NotHappenWithin,
	should.BeChronological,

	should.BeError,
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
