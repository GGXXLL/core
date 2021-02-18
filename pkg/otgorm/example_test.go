package otgorm_test

import (
	"fmt"

	"github.com/DoNewsCode/std/pkg/core"
	"github.com/DoNewsCode/std/pkg/otgorm"
	"gorm.io/gorm"
)

func Example() {
	// Suppress log output by gorm in this test.
	c := core.New(core.WithInline("log.level", "error"))
	c.AddCoreDependencies()
	c.AddDependencyFunc(otgorm.ProvideDatabase)
	c.Invoke(func(db *gorm.DB) {
		fmt.Println(db.Name())
	})
	// Output:
	// mysql
}
