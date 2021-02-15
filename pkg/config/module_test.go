package config

import (
	"github.com/DoNewsCode/std/pkg/container"
	"github.com/DoNewsCode/std/pkg/contract"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func setup() *cobra.Command {
	os.Remove("./mock/module_test.yaml")
	os.Remove("./mock/module_test.json")
	var cont container.Container
	var mod = Module{Container: &cont}
	cont.AddModule(MockModule(func() []contract.ExportedConfig {
		return []contract.ExportedConfig{
			{
				"foo",
				map[string]interface{}{
					"foo": "bar",
				},
				"A mock config",
			},
		}
	}))
	cont.AddModule(Module{Container: &cont})
	rootCmd := &cobra.Command{
		Use: "root",
	}
	mod.ProvideCommand(rootCmd)
	return rootCmd
}

type MockModule func() []contract.ExportedConfig

func (m MockModule) ProvideConfig() []contract.ExportedConfig {
	return m()
}

func TestModule_ProvideCommand(t *testing.T) {
	rootCmd := setup()
	cases := []struct {
		name string
		args []string
	}{
		{
			"new yaml",
			[]string{"exportConfig", "--outputFile", "mock/module_test.yaml"},
		},
		{
			"old yaml",
			[]string{"exportConfig", "--outputFile", "mock/module_test.yaml"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			rootCmd.SetArgs(c.args)
			rootCmd.Execute()
			testTarget, _ := ioutil.ReadFile("mock/module_test.yaml")
			expected, _ := ioutil.ReadFile("mock/module_test_expected.yaml")
			assert.Equal(t, expected, testTarget)
		})
	}
	cases = []struct {
		name string
		args []string
	}{
		{
			"new json",
			[]string{"exportConfig", "--outputFile", "mock/module_test.json", "--style", "json"},
		},
		{
			"old json",
			[]string{"exportConfig", "--outputFile", "mock/module_test.json", "--style", "json"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rootCmd.SetArgs(c.args)
			rootCmd.Execute()
			testTarget, _ := ioutil.ReadFile("mock/module_test.json")
			expected, _ := ioutil.ReadFile("mock/module_test_expected.json")
			assert.Equal(t, expected, testTarget)
		})
	}
}