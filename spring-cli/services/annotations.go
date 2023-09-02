package services

var MTM_MAP = map[uint64]string{}
var MANY_TO_MANY = `@ManyToMany
    @JoinTable(name="{%class_name%}_{%target_class_name%}",
            joinColumns = @JoinColumn(name = "{%class_name%}_id", referencedColumnName = "id"),
            inverseJoinColumns = @JoinColumn(name = "{%target_class_name%}_id", referencedColumnName = "id")
    )
    `
var MANY_TO_ONE = `@ManyToOne
    @JoinColumn(name = "{%class_name%}_id")
    `

var ONE_TO_MANY = `@OneToMany(mappedBy = "{%class_name%}")
    `
