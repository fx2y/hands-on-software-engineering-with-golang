package dependency_test

import (
	"github.com/fx2y/hands-on-software-engineering-with-golang/dependency"
	mock_dependency "github.com/fx2y/hands-on-software-engineering-with-golang/dependency/mock"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestDependencyCollector(t *testing.T) {
	// Create a controller to manage all our mock objects and make sure that all expectations were met before completing
	// the test
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Obtain a mock instance that implement API and associate it with the controller.
	api := mock_dependency.NewMockAPI(ctrl)

	// Verify not just the call list but the call order as well!
	gomock.InOrder(
		api.EXPECT().ListDependencies("proj0").Return([]string{"proj1", "res1"}, nil),
		api.EXPECT().DependencyType("proj1").Return(dependency.DepTypeProject, nil),
		api.EXPECT().DependencyType("res1").Return(dependency.DepTypeResource, nil),
		api.EXPECT().ListDependencies("proj1").Return([]string{"res1", "res2"}, nil),
		api.EXPECT().DependencyType("res2").Return(dependency.DepTypeResource, nil),
	)

	collector := dependency.NewCollector(api)
	depList, err := collector.AllDependencies("proj0")
	if err != nil {
		t.Fatal(err)
	}

	if exp := []string{"proj1", "res1", "res2"}; !reflect.DeepEqual(depList, exp) {
		t.Fatalf("expected dependency list to be:\n%v\ngot:\n%v", exp, depList)
	}
}
