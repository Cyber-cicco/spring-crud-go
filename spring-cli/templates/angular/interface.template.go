package angular

var INTERFACE_TEMPLATE = `export interface {%class_name%} {
{%attributes%}
}
`
var INTERFACE_ATTRIBUTE_TEMPLATE = `  {%attribute_name%}: {%attribute_type%};
`
