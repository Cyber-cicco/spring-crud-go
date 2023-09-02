# Spring-go

## Présentation

Spring-go est une application en ligne de commande aidant à la mise en place et la gestion de projet pour les applications fullstack Spring-boot / Angular basée sur des API REST. Elle s'appuie sur le respect du design pattern controller - service - repository jpa - mapper - dto

  On part du principe suivant : un développeur back a pour travail d'écrire de la logique métier, et un développeur front de mettre en place une UI pour que les utilisateurs puissent interagir avec le back. Tout temps que le développeur back passe à écrire du boilerplate pour permettre les opérations CRUD, et tout le temps que le développeur front passe à définir de la logique métier déjà présente dans le back est du temps perdu. Et les stacks Spring-Boot / Angular font très certainement partis des pires pour ce qui est de faire écrire du code inutile à ses développeurs

Le but de cette application est de tenter de remedier aux problèmes de ce stack technique.

## Problématiques que l'on cherche à régler

### La verbosité des opérations des CRUD dans Spring Boot
 * L'utilisation de design patterns basés sur les repository JPA, les mappers et les Dtos amène à **l'écriture d'une grande quantité de code répétitif et difficile à abstraire sans inutilement complexifier le code** (notamment en ce qui concerne les mappers). Certains framework de mapping comme Mapstruct tentent de remédier à cela, mais ils deviennent vite complexes à utiliser lorsque l'on souhaite des règles particulières. Notre application propose ainsi la possibilité de **générer du boilerplate (et uniquement du boilerplate)**, proposant une forme de template pour un controller, un service, un repository, un mapper et un dto pour chaque entité d'un projet.

### La verbosité de JPA dans la création des entités
 * L'initialisation d'un projet peut parfois **prendre beaucoup de temps sur Spring**, notamment du fait que l'on ai besoin d'écrire des entités JPA et que la configurations de celles-ci est souvent très verbeuse. Spring-go propose de pouvoir faire cela de façon beaucoup plus rapide en tappant le nom de l'entité et de ses champs en ligne de commande pour générer des fichiers de configuration permettant ensuite de générer les entités JPA, un peu à la façon d'un **JPA Buddy en ligne de commande**

### L'inconsistance des designs patterns dans un gros projet Spring boot
 * Il est assez facile pour un gros projet d'avoir plusieurs personnes débarquant chacune avec leurs habitudes concernant le code. Certains vont suffixé leurs controllers par "Controller", d'autres par "Ctrl", etc. Certains vont créer les mappers à la main, d'autres vont utiliser des framework de mapping. Certains aiment avoir des interfaces définissant des opérations nécessaires pour chaque service, d'autres vont oublier des les implémenter. Pour cela, ce programme propose un **fichier de configuration et des templates permettant d'assurer que chaque élément généré par le CLI de suivre certaines règles de design concernant les suffixes des classes, les interfaces qu'elles implémentent, etc.**

### Le manque d'outils d'écriture de Java pour les éditeurs de texte classiques
 * Toutes les IDEs ne proposent pas nécessairement d'utilitaire de génération de classe, notamment les éditeurs de texte boostés aux plugins comme VSCode, NeoVim et Vim. Le but de ce projet est d'offir un programme qui permet de **créer divers types de classe à partir d'une ligne de commande.**
  Par défaut, il y a des templates pour les classes, les interfaces, les records, les enums, les controllers, les repository, les services, les dtos, les entités, les mappers et les exceptions. Il est possible d'en rajouter d'autres en changeant quelques lignes de code, et de personaliser celles existantes en changeant simplement le contenu des templates.

### La difficulté de la cohérence des types entre back et front

 * Lorsque l'on utilise deux langages différents pour le back et le front, il peut devenir difficile d'assurer la cohérence des types entre les deux entités, et il est toujours nécessaire de générer deux fois les mêmes types / structures / classes dans deux langages différents. Ici, on propose la possibilité de **générer des interfaces typescript à partir des Dtos du back**

 ### La redondance entre controllers cotés back et appels http coté front

 * Lorsque l'on a défini un endpoint coté back, il n'y a dans Angular pas trente façon de faire des appels au back : on crée un service utilisant le HttpClient du framework pour faire les requêtes. L'agencement du code de ce fichier est entièrement fonction du code coté back, ce qui veut dire qu'il peut être généré automatiquement en parsant le code Java des controllers. C'est donc ce que propose ce CLI : **générer automatiquement toutes les requêtes HTTP vers le back dans des services Angular**
 
 ### La redondance des formulaires de changement des entités:

 * De la même façon que les appels HTTP sont entièrement dépendants de la logique des controllers, les formulaires sont aussi très dépendants de la forme des DTOs et de la logique de validation de ceux-ci. Si, à l'heure actuelle, le CLI ne permet pas encore de générer ces formulaires, il devrait être possible de **créer des composants en fonction des Dtos et des annotations de validation présents sur ceux-ci.**

## Comment utiliser l'application

### Installation:

Actuellement, il faut copier le dossier springCli, go.mod, go.sum et spring-parameters.json à la racine d'un projet Spring boot. Ensuite, il faut lancer le fichier .cmd dans le dossier springCli/cmd/ pour pouvoir lancer l'application. Le binaire ne fonctionnera potentiellement pas sur Windows, il faut donc mieux avoir go installé et lancer la commande
```bash
go build
```
pour créer le binaire adapté à votre OS.
Si go n'est pas installé, vous pouvez utiliser ce lien pour l'installer :
[https://go.dev/doc/install](https://go.dev/doc/install)

Ensuite, il suffit d'exécuter le fichier cmd de cette façon : 

```bash
./cmd
```

### Configuration
La configuration par défaut se trouve dans spring-parameters.json.
Il contient les paramètres suivants:
*  base-package : le package de base du projet, qui préfixe tous les autres

* erease-files : détermine si, lorsque l'on souhaite créer un fichier dont le nom existe déjà, il est écrasé ou non

* ts-interface-folder : dossier dans lequel se trouveront stockés les interfaces typescript. Il est conseillé de le configurer vers un projet Angular (ou typescript d'une manière générale)

* ts-service-folder :  La même chose mais pour les services de requêtes HTTP

* nom-package : contient des options pour la gestion des classes d'un certain type

* package : sous package de la classe d'un certain type

* package-policy : deux options possibles : appended, et dans ce cas, quand on précise un package supplémentaire pour une entité (par exemple : "administration.BonDeCommande"), le package supplémentaire sera rajouté après le package de base.
Sinon, on peut choisir l'option "base", qui permet de d'ignorer le package supplémentaire

* suffix : permet de déterminer le suffix des classes d'un certain type. Par exemple, "Transformer" pour les classes de type mapper.

### Générer une classe simple



 ## Ce que ce projet est / veut être

 ### Une API pour générer n'importe quel boilerplate pour ce stack

  * Si les templates demandents encore de recompiler le projet pour être changé, le but de ce projet est de fournir des templates et des fichiers de configurations totalement personnalisables et non liées à des bibliothèques particulières.

### Un utilitaire fournissant des composants réutilisables pour créer une des plugins pour des éditeurs de texte

 * Actuellement, les options viables pour écrire du Java de façon efficiente se limitent à IntelliJ. Le  LSP fournit par Eclipse n'est clairement pas suffisant pour permettre à des éditeurs de texte de rendre agréable l'écriture d'un langage aussi verbeux que Java, et cela rend tout dev Java s'étant habitué à l'expérience de l'analyseur de code d'IntelliJ complétement dépendant de JetBrains
 
### Un exercice
* Ce projet a bien plus été créé pour me permettre de monter en compétence dans l'apprentissage de nouveaux langages, la création d'algorithme et l'analyse de code que pour être réellement le plus efficace possible. Il existe surement de très bonnes librairies d'utilitaires de manipulation de chaine de caractère, de templating ou de création de fichiers et dossiers. Mais du fait du but pédagogique du projet, j'ai préféré créer presque tout à la main.
 
 ## Ce que ce projet n'est pas / ne veut pas être

 ### Un meilleur stack
  * Si nous n'étions pas déjà au courant qu'un stack Spring-Boot / Angular n'était pas forcément optimal d'un point de vue de l'expérience de développement, nous n'aurions pas créé un utilitaire de génération de code en premier lieu. Ce projet est fait pour les personnes coincées dans ce stack, quelqu'en soit la raison.

### Jhipster / Wordpress
 * L'outil a pour but de générer du boilerplate personnalisable en fonction des librairies d'un projet, de ses conventions de nommage, de l'organisation de ses packages, etc (même s'il y a une grosse dépendance à Lombok). D'un point de vue général, on ne souhaite intégrer aucune forme de logique métier dans les classes, et aucune forme de mise à jour permettant de créer automatiquement de la configuration de sécurité ou quelque chose du même acabit n'est prévue. Le but est d'en faire une forme d'API de génération de boilerplate pour un projet Spring Boot / Angular, pas un outil qui crée un site internet remplis de code non controllé en tappant "springgo create facebook" dans le terminal.

