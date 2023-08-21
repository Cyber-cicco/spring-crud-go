package config


var JPA_CONFIG_CREATION_ARG = "jpa"

var USAGE_MESSAGE = 
`
-----------------------
spring-go - utilisation
-----------------------

/!\ Attention : le projet est en pre-alpha et certaines fonctionalités pourraient ne pas être implémentés ou avoir changé.

commandes:
    jpa [...args] : permet de créer le fichier de configuration d'une entité dans le dossier spring-cli/jpa. Prend pour arguments la classe et le nom des champs de l'entité

        -c : le nom de la classe. Il est possible de le préfixer par un nom de package en séparant le package et la classe par un point.
        Exemple : monpackage.MaClasse
        A noter qu'il n'est pas nécessaire de préciser le package racine du projet pour chaque classe, cela est configuré dans le fichier de configuration principal

        -f "<fields>" : permet de préciser le nom de chaque fields de l'entité JPA, en les séparant par des espaces. Le nom du field n'accepte que des caractères alphaNumériques
        Ces fields voient leur type inféré via des règles posées dans le fichier type-inferer. Elles sont customisables en changeant simplement la logique de ce fichier, et n'ont pas vocation à être définitives ni exhaustives
        Il est possible de poser des préfixes et des suffixes à ces entités:
        <field>List permet de signifier qu'il s'agit d'un type liste.
        *<field> permet de signifier que le nom de la variable doit être également pris pour type.
        <field>:<type> permet de préciser le type explicitement après le nom du field.
        <field>@<annotation> permet de préciser l'existence d'une annotation
        Annotations possibles:
            - otm : annotation OneToMany
            - mto : annotation ManyToOne
            - mtm : annotation ManyToMany
            - enum : annotation Enumerated

        -clear : permet de supprimer tous les fichiers du répertoire jpa


    project : permet de créer le projet spring en fonction des spécifications précisées dans spring-parameters.json et des fichiers de configuration des entités JPA dans le dossier jpa. Va donc créer un repository, un service, un controller, un mapper et un DTO pour chaque entité. Pour cela, va utiliser un template présent dans le package java-classes pour chaque entité, et va s'en servir pour créer chaque classe selon des règles particulières. Les templates sont customisables, ce qui permet de les adapter à sa façon de faire.
    
    init : permet de créer le fichier de configuration du projet s'il n'existe pas déjà. Devrait également être capable de créer la structure du projet spring avec cela, grâce à des templates. Par défaut, ces templates sont paramétrés pour aller avec un projet maven et spring boot avec des dépendances spécifiques, mais il est possible de changer les templates. D'autres options devraient également pouvoir exister pour update les templates en fonction des changements que l'on a fait sur ces fichiers.
        
        - p : précise le package de base du projet

    class : permet de créer une classe de base, un peu à la façon de l'utilitaire permettant de créer le boilerplate des classes de base dans IntelliJ. Prend une option pour le nom de la classe, et prend une option pour le type de la classe

        -c : le nom de la classe. Il est possible de le préfixer par un nom de package en séparant le package et la classe par un point.
        Exemple : monpackage.MaClasse

        -t : le type de la classe. Paramètre optionnel. Si jamais il n'est pas renseigné, cela permet juste de créer une classe java classique.

        Liste des options possibles :

            ctrl : crée un controller
            
            srv : crée un service
            
            ent : crée une entité

            map : crée un mapper

            dto : crée un dto

            repo : crée un repository

            ex : crée une exception

`
