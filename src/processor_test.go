package dbanon

import (
	"strings"
	"testing"
)

type TestProvider struct{}

func NewTestProvider() *TestProvider {
	p := &TestProvider{}

	return p
}

func (p TestProvider) Get(s string) string {
	return s
}

func TestProcessLine(t *testing.T) {
	config, _ := NewConfig("magento2")
	provider := NewTestProvider()
	processor := NewLineProcessor(config, provider)

	r1 := processor.ProcessLine("foobar")
	if r1 != "foobar" {
		t.Errorf("Got %s wanted foobar", r1)
	}

	r2 := processor.ProcessLine("INSERT INTO `admin_user` (`firstname`) VALUES ('bob');")
	if strings.Contains(r2, "bob") {
		t.Error("Got bob wanted no bob")
	}

	r3 := processor.ProcessLine("INSERT INTO `admin_user` (`user_id`) VALUES (1337);")
	if !strings.Contains(r3, "1337") {
		t.Error("Got no 1337 wanted 1337")
	}

	for _, e := range processor.Config.Eav {
		if e.Name == "customer" {
			e.Attributes["1"] = "firstname"
		}
	}

	r4 := processor.ProcessLine("INSERT INTO `customer_entity_varchar` (`attribute_id`, `value`) VALUES (1, 'bob');")
	if strings.Contains(r4, "bob") {
		t.Error("Got bob wanted no bob")
	}
}