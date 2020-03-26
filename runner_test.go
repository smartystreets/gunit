package gunit

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

/**************************************************************************/
/**************************************************************************/

func TestRunnerEndsFatallyIfFixtureIsIncompatible(t *testing.T) {
	t.Parallel()

	test := Setup(false)
	ensureEmbeddedFixture(new(FixtureWithoutEmbeddedGunitFixture), test.fakeT)
	assertions.New(t).So(test.fixture.Failed(), should.BeTrue)
}

type FixtureWithoutEmbeddedGunitFixture struct {
	Fixture string /* should be: *gunit.Fixture */
}

/**************************************************************************/
/**************************************************************************/

func TestMarkedAsSkippedIfNoTestCases(t *testing.T) {
	t.Parallel()

	Run(new(FixtureWithNoTestCases), t, Options.AllSequential())
}

type FixtureWithNoTestCases struct{ *Fixture }

/**************************************************************************/
/**************************************************************************/

func TestRunnerFixtureWithSetupAndTeardown(t *testing.T) {
	t.Parallel()

	defer assertSetupTeardownInvocationsInCorrectOrder(t)
	Run(new(RunnerFixtureSetupTeardown), t, Options.AllSequential())
}
func assertSetupTeardownInvocationsInCorrectOrder(t *testing.T) {
	expectedInvocations := []string{
		"Setup", "Test3", "Teardown",
		"Setup", "Test1", "Teardown",
		// Test2 and Test4 are always skipped
	}
	if testing.Short() {
		expectedInvocations = expectedInvocations[3:]
	}
	assertions.New(t).So(invocations_A, should.Resemble, expectedInvocations)
}

var invocations_A []string

type RunnerFixtureSetupTeardown struct{ *Fixture }

func (this *RunnerFixtureSetupTeardown) Setup()     { invocations_A = append(invocations_A, "Setup") }
func (this *RunnerFixtureSetupTeardown) Teardown()  { invocations_A = append(invocations_A, "Teardown") }
func (this *RunnerFixtureSetupTeardown) Test1()     { invocations_A = append(invocations_A, "Test1") }
func (this *RunnerFixtureSetupTeardown) SkipTest2() { invocations_A = append(invocations_A, "Test2") }
func (this *RunnerFixtureSetupTeardown) LongTest3() { invocations_A = append(invocations_A, "Test3") }
func (this *RunnerFixtureSetupTeardown) SkipLongTest4() {
	invocations_A = append(invocations_A, "Test4")
}

/**************************************************************************/
/**************************************************************************/

func TestRunnerFixture(t *testing.T) {
	t.Parallel()

	defer assertInvocationsInCorrectOrder(t)
	Run(new(RunnerFixturePlain), t, Options.AllSequential())
}
func assertInvocationsInCorrectOrder(t *testing.T) {
	expectedInvocations := []string{"Test3", "Test1"} // Test2 and Test4 are always skipped
	if testing.Short() {
		expectedInvocations = expectedInvocations[1:]
	}
	assertions.New(t).So(invocations_B, should.Resemble, expectedInvocations)
}

var invocations_B []string

type RunnerFixturePlain struct{ *Fixture }

func (this *RunnerFixturePlain) Test1()         { invocations_B = append(invocations_B, "Test1") }
func (this *RunnerFixturePlain) SkipTest2()     { invocations_B = append(invocations_B, "Test2") }
func (this *RunnerFixturePlain) LongTest3()     { invocations_B = append(invocations_B, "Test3") }
func (this *RunnerFixturePlain) SkipLongTest4() { invocations_B = append(invocations_B, "Test4") }

/**************************************************************************/
/**************************************************************************/

func TestRunnerFixtureWithFocus(t *testing.T) {
	t.Parallel()

	defer assertFocusIsOnlyInvocation(t)
	Run(new(RunnerFixtureFocus), t, Options.AllSequential())
}
func assertFocusIsOnlyInvocation(t *testing.T) {
	assertions.New(t).So(invocations_C, should.Resemble, []string{"Test3"})
}

var invocations_C []string

type RunnerFixtureFocus struct{ *Fixture }

func (this *RunnerFixtureFocus) Test1()      { invocations_C = append(invocations_C, "Test1") }
func (this *RunnerFixtureFocus) Test2()      { invocations_C = append(invocations_C, "Test2") }
func (this *RunnerFixtureFocus) FocusTest3() { invocations_C = append(invocations_C, "Test3") }
func (this *RunnerFixtureFocus) Test4()      { invocations_C = append(invocations_C, "Test4") }

/**************************************************************************/
/**************************************************************************/

func TestRunnerFixtureWithFocusLong(t *testing.T) {
	t.Parallel()

	defer assertFocusLongIsOnlyInvocation(t)
	Run(new(RunnerFixtureFocusLong), t, Options.AllSequential())
}
func assertFocusLongIsOnlyInvocation(t *testing.T) {
	expected := []string{"Test3"}
	if testing.Short() {
		expected = nil
	}
	assertions.New(t).So(invocations_D, should.Resemble, expected)
}

var invocations_D []string

type RunnerFixtureFocusLong struct{ *Fixture }

func (this *RunnerFixtureFocusLong) Test1()          { invocations_D = append(invocations_D, "Test1") }
func (this *RunnerFixtureFocusLong) Test2()          { invocations_D = append(invocations_D, "Test2") }
func (this *RunnerFixtureFocusLong) FocusLongTest3() { invocations_D = append(invocations_D, "Test3") }
func (this *RunnerFixtureFocusLong) Test4()          { invocations_D = append(invocations_D, "Test4") }

/**************************************************************************/
/**************************************************************************/

func TestRunnerFixtureWithOnlyOneFocus(t *testing.T) {
	t.Parallel()

	defer assertSingleFocusIsOnlyInvocation(t)
	Run(new(RunnerFixtureWithOnlyOneFocus), t, Options.AllSequential())
}
func assertSingleFocusIsOnlyInvocation(t *testing.T) {
	assertions.New(t).So(invocations_E, should.Resemble, []string{"Test1"})
}

var invocations_E []string

type RunnerFixtureWithOnlyOneFocus struct{ *Fixture }

func (this *RunnerFixtureWithOnlyOneFocus) FocusTest1() {
	invocations_E = append(invocations_E, "Test1")
}

/**************************************************************************/
/**************************************************************************/
