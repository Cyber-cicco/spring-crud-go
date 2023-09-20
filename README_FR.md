# Spring-go

## Présentation

Spring-go est une application en ligne de commande aidant à la mise en place et la gestion de projet pour les applications fullstack Spring-boot / Angular basée sur des API REST. Elle s'appuie sur le respect du design pattern controller - service - repository jpa - mapper - dto.

  On part du principe suivant : un développeur back a pour travail d'écrire de la logique métier, et un développeur front de mettre en place une UI pour que les utilisateurs puissent interagir avec le back. Tout temps que le développeur back passe à écrire du boilerplate pour permettre les opérations CRUD, et tout le temps que le développeur front passe à de la logique du back est du temps perdu. Et les stacks Spring-Boot / Angular sont très friands de ce type de code.

Le but de cette application est de tenter de remedier aux problèmes de ce stack technique.

## Problématiques que l'on cherche à régler

### La verbosité des opérations des CRUD dans Spring Boot
 * L'utilisation de design patterns basés sur les repository JPA, les mappers et les Dtos amène à **l'écriture d'une grande quantité de code répétitif et difficile à abstraire sans inutilement complexifier le code** (notamment en ce qui concerne les mappers). Certains framework de mapping comme Mapstruct tentent de remédier à cela, mais ils deviennent vite complexes à utiliser lorsque l'on souhaite implémenter du mapping pour les champs complexes. Notre application propose ainsi la possibilité de **générer du boilerplate (et uniquement du boilerplate)**, proposant une forme de template pour un controller, un service, un repository, un mapper et un dto pour chaque entité d'un projet.

### La verbosité de JPA dans la création des entités
 * L'initialisation d'un projet peut parfois **prendre beaucoup de temps sur Spring**, notamment du fait que l'on ai besoin d'écrire des entités JPA et que la configurations de celles-ci est souvent très verbeuse. Spring-go propose de pouvoir faire cela de façon beaucoup plus rapide en tappant le nom de l'entité et de ses champs en ligne de commande pour générer des fichiers de configuration permettant ensuite de générer les entités JPA, un peu à la façon d'un **JPA Buddy en ligne de commande**

### L'inconsistance des designs patterns dans un gros projet Spring boot
 * Il est assez facile pour un gros projet d'avoir plusieurs personnes arrivant chacunes avec leurs habitudes concernant le code. Certains vont suffixé leurs controllers par "Controller", d'autres par "Ctrl", etc. Certains vont créer les mappers à la main, d'autres vont utiliser des framework de mapping. Certains aiment avoir des interfaces définissant des opérations nécessaires pour chaque service, d'autres vont oublier des les implémenter. Pour cela, ce programme propose un **fichier de configuration et des templates permettant d'assurer que chaque élément généré par le CLI de suivre certaines règles de design concernant les suffixes des classes, les interfaces qu'elles implémentent, etc.**

### Le manque d'outils d'écriture de Java pour les éditeurs de texte classiques
 * Toutes les IDEs ne proposent pas nécessairement d'utilitaire de génération de classe, notamment les éditeurs de texte boostés aux plugins comme VSCode, NeoVim et Vim. Le but de ce projet est d'offir un programme qui permet de **créer divers types de classe à partir d'une ligne de commande.**
  Par défaut, il y a des templates pour les classes, les interfaces, les records, les enums, les controllers, les repository, les services, les dtos, les entités, les mappers et les exceptions. Il est possible d'en rajouter d'autres en changeant quelques lignes de code, et de personaliser celles existantes en changeant simplement le contenu des templates. **A noter que ce type de template existe déjà dans IntelliJ**

### La difficulté de la cohérence des types entre back et front

 * Lorsque l'on utilise deux langages différents pour le back et le front, il peut devenir difficile d'assurer la cohérence des types entre les deux entités, et il est toujours nécessaire de générer deux fois les mêmes types / structures / classes dans deux langages différents. Ici, on propose la possibilité de **générer des interfaces typescript à partir des Dtos du back**

 ### La redondance entre controllers cotés back et appels http coté front

 * Lorsque l'on a défini un endpoint coté back, il n'y a dans Angular pas trente façon de faire des appels au back : on crée un service utilisant le HttpClient du framework pour faire les requêtes. L'agencement du code de ce fichier est entièrement fonction du code coté back, ce qui veut dire qu'il peut être généré automatiquement en parsant le code Java des controllers. C'est donc ce que propose ce CLI : **générer automatiquement toutes les requêtes HTTP vers le back dans des services Angular**
 
 ### La redondance des formulaires de changement des entités:

 * De la même façon que les appels HTTP sont entièrement dépendants de la logique des controllers, les formulaires sont aussi très dépendants de la forme des DTOs et de la logique de validation de ceux-ci. Si, à l'heure actuelle, le CLI ne permet pas encore de générer ces formulaires, il devrait être possible de **créer des composants en fonction des Dtos et des annotations de validation présents sur ceux-ci.**

## Comment utiliser l'application

### Installation:

Pour une installation sans volonté de modifier le code source, il est maintenant possible de se 
saisir de la relase dans le dépôt GitHub et de lancer le script d'installation en suivant les instructions du fichier INSTALL.

L'installation fonctionne et a été testé dans un environnement Linux. Il existe également un binaire fonctionnel pour Windows dans une autre release, mais cela nécessite d'avoir installé GitBash, et demande d'exécuter le script via GitBash 

Sinon, vous pouvez cloner le repo, aller dans le dossier spring-cli, et lancer :

```bash
go install
```

pour créer le binaire adapté à votre OS.
Si go n'est pas installé, vous pouvez utiliser ce lien pour l'installer :
[https://go.dev/doc/install](https://go.dev/doc/install)

Ensuite, il suffit de créer un alias pointant vers le fichier ~/go/bin/spring-cli pour pouvoir commencer à utiliser l'outil en ligne de commande

### Initialiser un projet 

**En supposant que l'alias que vous avez choisi est springgo,** lancez, à la racine d'un projet Spring Boot, la commande suivante:

```bash
springgo init
```

Cela créera un dossier templates/ contenant des templates au format texte de chaque élément d'une 
classe dans un dossier qui lui correspond. Par exemple, un Repository possédera un dossier templates/Repository/ dans lequel on va trouver une fichier ClassType.txt contenant juste le mot 
```
interface
```
Le changer en 
```
class
```
fera en sorte que les Repository créés ne soient plus des interfaces, mais des
classes. Plus d'informations dans la partie **Personnalisation**.

Initialiser un projet crée également un dossier jpa/ , dans lequel se trouveront 
tous les fichiers de configuration générés par la commande jpa. Plus 
d'information dans la partie **Générer des classes de CRUD** 

Enfin, initialiser un projet crée un fichier spring-parameters.json. Plus d'informations dans la partie :

### Configuration
La configuration par défaut se trouve dans spring-parameters.json.
Il contient les paramètres suivants:
*  base-package : le package de base du projet, qui préfixe tous les autres

* erease-files : détermine si, lorsque l'on souhaite créer un fichier dont le nom existe déjà, il est écrasé ou non **IL EST FORTEMENT CONSEILLÉ DE LE GARDER À FALSE**. En effet la commande ./cmd project pourrait rééinitialiser des classes entières si vous avez oublié de faire le ménage dans vos fichiers du dossier JPA en voulant générer de nouvelles classes (plus de précision là dessus dans la partie "génération de projet")

* ts-interface-folder : dossier dans lequel se trouveront stockés les interfaces typescript. Il est conseillé de le configurer vers un projet Angular (ou typescript d'une manière générale)

* ts-service-folder :  La même chose mais pour les services de requêtes HTTP

* nom-package : contient des options pour la gestion des classes d'un certain type

* package : sous package de la classe d'un certain type

* package-policy : deux options possibles : appended, et dans ce cas, quand on précise un package supplémentaire pour une entité (par exemple : "administration.BonDeCommande"), le package supplémentaire sera rajouté après le package de base.
Sinon, on peut choisir l'option "base", qui permet de d'ignorer le package supplémentaire

* suffix : permet de déterminer le suffix des classes d'un certain type. Par exemple, "Transformer" pour les classes de type mapper.

### Générer une classe simple

```bash
./cmd class -c Foo
```

Permet de générer une classe simple :

```java
package com.example.springgo;


public class Foo {

}
```
Celle-ci sera placée dans le package précisée dans [spring-parameters.json](spring-parameters.json)

On peut également générer une classe de cette façon : 

```bash
./cmd class -c bar.Foo
```
Ainsi, la classe sera placée dans le package bar, ajouté au package de base précisé dans le fichier de configuration.

#### Astuce :

Il est possible de changer le paramètre "package" de "default-package" dans 
[spring-parameters.json](spring-parameters.json) pour ne pas avoir à repréciser
le package à chaque création de classe si vous travailler tout le temps 
dans le même package

#### Options possibles :

L'option -t permet de préciser un type de classe particulier, de cette façon :

Générer un controller :
```bash
./cmd class -c Foo -t ctrl
```
```java
package com.example.springgo.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;    


@RestController
@RequiredArgsConstructor
@RequestMapping("/foo")
public class FooController {

}
```
Générer un service :
```bash
./cmd class -c Foo -t srv
```
```java
package com.example.springgo.service;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class FooService {

}

```
Générer une entité :
```bash
./cmd class -c Foo -t ent
```
```java
package com.example.springgo.entites;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.Builder;
import lombok.NoArgsConstructor;    


@NoArgsConstructor
@AllArgsConstructor
@Data
@Entity
@Builder
public class Foo {

}
```
Générer un mapper :
```bash
./cmd class -c Foo -t map
```
```java
package com.example.springgo.dto;


import org.springframework.stereotype.Component;


@Component
public class FooTransformer {

}
```
Générer un Dto :
```bash
./cmd class -c Foo -t dto
```
```java
package com.example.springgo.dto;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;


@AllArgsConstructor
@NoArgsConstructor
@Data
@Builder
public class FooDto {

}
```
Générer un Repo :
```bash
./cmd class -c Foo -t repo
```
```java
package com.example.springgo.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;


public interface FooRepository extends JpaRepository<Foo, Long>  {

}

```
Générer une Exception :
```bash
./cmd class -c Foo -t ex
```
```java
package com.example.springgo.exception;


public class FooException extends RuntimeException {

}

```
Générer un Enum :
```bash
./cmd class -c Foo -t enum
```
```java
package com.example.springgo;


public enum Foo {

}
```
Générer une interface :
```bash
./cmd class -c Foo -t int
```
```java
package com.example.springgo;


public interface Foo {

}
```
Générer un record :
```bash
./cmd class -c Foo -t rec
```
```java
package com.example.springgo;


public record Foo() {

}

```
Générer une annotation :
```bash
./cmd class -c Foo -t ano
```
```java
package com.example.springgo;


public @interface Foo {

}
```

### Générer des classes de CRUD pour de nouvelles entités :

Cela se fait en deux étapes

#### 1 Générer les fichiers de configuration des entités JPA

```bash
./cmd jpa -c Foo -f "bar nbBuzz dateBro"
```
```json
{
    "name": "Foo",
    "package": "com.example.springgo",
    "fields": [
        {
            "name": "bar",
            "type": "String",
            "options": {
                "Annotations": []
            }
        },
        {
            "name": "nbBuzz",
            "type": "Integer",
            "options": {
                "Annotations": []
            }
        },
        {
            "name": "dateBro",
            "type": "LocalDate",
            "options": {
                "Annotations": []
            }
        }
    ]
}
```

L'utilisation basique de cette option de ligne de commande permet de préciser le nom de la classe avec l'option -c, et les noms des fields avec l'option -f. Par défaut, l'application va essayer d'inférer le type du champ à partir de son nom.
Il est possible d'observer les règles de l'inférence de type dans [ce fichier](./spring-cli/utils/type-inferer.go)

Cependant, si ce genre de comportement ne convient pas, il est possible de typer à la main le field en utilisant la syntaxe suivante :

\<nom_field>:<type_field>

Exemple : 
```
 ./cmd jpa -c Foo -f "bar nbBuzz:Long dateBro"
```
Ici, nbBuzz sera de type Long

On peut également préciser certaines annotations en utilisant la syntaxe suivante:

\<field>@\<annotation>

il existe différentes annotations possibles:

@mtm : Many to Many

@mto : Many to One

@otm : One to Many

Il est possible d'aller vérifier la logique de création d'annotation dans [ce fichier](./spring-cli/services/java-classes/shared.go), et également d'en rajouter d'autres si vous le souhaitez

Enfin, dernière option, il est possible de préciser que le nom du champ doit prendre le type du nom du champ en mettant sa première lettre en majuscule. Si cela peut paraitre très spécifique, cela permet en fait de préciser que le champ est une entité. Par exemple:

```bash
 ./cmd jpa -c Foo -f "*bar nbBuzz dateBro "
```
va générer :
```json
{
    "name": "Foo",
    "package": "com.example.springgo",
    "fields": [
        {
            "name": "bar",
            "type": "Bar",
            "options": {
                "Annotations": []
            }
        },
        {
            "name": "nbBuzz",
            "type": "Integer",
            "options": {
                "Annotations": []
            }
        },
        {
            "name": "dateBro",
            "type": "LocalDate",
            "options": {
                "Annotations": []
            }
        }
    ]
}
```

Ce fichier de configuration se trouvera dans le dossier springCli/jpa/

#### 2 Générer le projet

```
 ./cmd project
```

Cette commande va générer le projet à partir des fichiers de configuration des entités JPA.
Pout ces deux commandes :

```bash
./cmd jpa -c Foo -f "*bar@mto nbPoint dateCreation"
./cmd jpa -c Bar -f "*foo@otm titre dateEcheance"
```

![projet généré](./img/project.png)

#### Supprimer ses d'entités JPA

Pour supprimer les fichiers de configuration de entités JPA et ne donc par régénérer les classes chaque fois, lancez la commande suivante : 

```bash
./cmd jpa -clear
```

### Générer des interfaces et des services le module HTTP de Angular

La commande est extrêmement simple :

```bash
./cmd ng
```
Cela va générer, dans le dossier précisé dans spring-parameters.json, les interfaces et les services correspondants aux controllers et Dtos du back.

Par exemple, pour le projet créé précédement, nous pouvons obtenir les résultats suivants:

interface de Foo :
```typescript
import { Foo } from './foo'

export interface Bar {
  id: number;
  foo: Foo;
  titre: string;
  dateEcheance: Date;

}
```
interface de Bar :
```typescript
import { Bar } from './bar'

export interface Foo {
  id: number;
  bar: Bar;
  nbPoint: number;
  dateCreation: Date;

}
```
Service d'appel au controller de Bar :
```typescript
import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Bar } from '../models/bar'


@Injectable({
  providedIn: 'root'
})
export class BarHttpService {

  private URL_ = environment.baseUrl + "/bar";

  constructor(private http:HttpClient) {}


  get(){
    return this.http.get<Bar[]>(this.URL_)
  }


  post(dto : Bar){
    return this.http.post<Bar>(this.URL_, dto)
  }


  put(dto : Bar){
    return this.http.put<Bar>(this.URL_, dto)
  }


  delete(dto : Bar){
    return this.http.delete<Bar>(this.URL_, dto)
  }

}
```
Service d'appel au controller de Foo :
```typescript
import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Foo } from '../models/foo'


@Injectable({
  providedIn: 'root'
})
export class FooHttpService {

  private URL_ = environment.baseUrl + "/foo";

  constructor(private http:HttpClient) {}


  get(){
    return this.http.get<Foo[]>(this.URL_)
  }


  post(dto : Foo){
    return this.http.post<Foo>(this.URL_, dto)
  }


  put(dto : Foo){
    return this.http.put<Foo>(this.URL_, dto)
  }


  delete(dto : Foo){
    return this.http.delete<Foo>(this.URL_, dto)
  }

}

```


 ## Ce que ce projet est / veut être

 ### Une API pour générer n'importe quel boilerplate pour ce stack

  * Si les templates demandents encore de recompiler le projet pour être changé, le but de ce projet est de fournir des templates et des fichiers de configurations totalement personnalisables et non liées à des bibliothèques particulières.

### Un utilitaire fournissant des composants réutilisables pour créer une des plugins pour des éditeurs de texte

 * Actuellement, les options viables pour écrire du Java de façon efficiente se limitent à IntelliJ. Le  LSP fournit par Eclipse n'est clairement pas suffisant pour permettre à des éditeurs de texte de rendre agréable l'écriture d'un langage aussi verbeux que Java, et cela rend tout dev Java s'étant habitué à l'expérience de l'analyseur de code d'IntelliJ complétement dépendant de JetBrains
 
### Un exercice
* Ce projet a bien plus été créé pour me permettre de monter en compétence dans l'apprentissage de nouveaux langages, la création d'algorithme et l'analyse de code que pour être réellement le plus efficace possible. Il existe surement de très bonnes librairies d'utilitaires de manipulation de chaine de caractère, de templating ou de création de fichiers et dossiers. Mais du fait du but pédagogique du projet, j'ai préféré créer presque tout à la main.
 
 ## Ce que ce projet n'est pas / ne veut pas être

 ### Un meilleur stack
  * Le but n'est pas de recréer un stack sans les difficultés de Spring Boot / Angular. Il est possible que les Server Component de React, Next.js, ou Go / HTMX soient des technologies plus adaptées au développement web moderne. Toujours est-il que Spring Boot / Angular existe, a des avantages pour accompagener ses difficultés, et qu'il est possible de
  rendre l'expérience de développement dans ces conditions plus agréable

### Jhipster / Wordpress
 * L'outil a pour but de générer du boilerplate personnalisable en fonction des librairies d'un projet, de ses conventions de nommage, de l'organisation de ses packages, etc (même s'il y a une grosse dépendance à Lombok). D'un point de vue général, on ne souhaite intégrer aucune forme de logique métier dans les classes, et aucune forme de mise à jour permettant de créer automatiquement de la configuration de sécurité ou quelque chose du même acabit n'est prévue. Le but est d'en faire une forme d'API de génération de boilerplate pour un projet Spring Boot / Angular, pas un outil qui crée un site internet remplis de code non controllé en tappant "springgo create facebook" dans le terminal.

## Personalisation

### Personnaliser les templates
Le script est fait pour être personnalisable en fonction de vos besoins particuliers pour un projet. Si vous utilisez javax au lieu de jakarta, vous devrez personnaliser la génération de classe Java pour la faire correspondre à votre environnement. Pour cela, chaque forme de classe possède un dossier qui lui correspond dans le dossier templates/ généré par la commande init.

Voici un résumé du rôle de chaque fichier :

* Annotations.txt : Correspond aux annotations de la classe
* Body.txt : Correspond au contenu de la classe, après sa déclaration
* ClassType.txt : Correspond au type de la classe : interface, class, record, etc.
* Extends : permet de définir une classe parente.
* Implements : permet de définir des interfaces implémentées
* Imports : permet de définir les imports de base de la classe

### Personnaliser la logique du script

La logique de création des classes se trouve dans spring-cli/services/class_writer.go
La documentation de chaque fonction arrive bientôt
