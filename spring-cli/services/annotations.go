package services

var MTM_MAP = map[uint64]string{}

var MANY_TO_ONE = `@ManyToOne
    @JoinColumn(name = "{%target_class_name%}_id")
    `

var ONE_TO_MANY = `@OneToMany(mappedBy = "{%class_name%}")
    `
