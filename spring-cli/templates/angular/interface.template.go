package angular

var INTERFACE_TEMPLATE = `{%imports%}
export interface {%class_name%} {
{%attributes%}
}
`
var INTERFACE_ATTRIBUTE_TEMPLATE = `  {%attribute_name%}: {%attribute_type%};
`
var ENTITY_IMPORT_TEMPLATE = `import { {%new_import%} } from './{%file_import%}'
`
