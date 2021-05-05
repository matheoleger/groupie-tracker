
# groupie-tracker
	# Groupie-Tracker

## :o: Comment ça marche 


## :o: Explication du projet

Nous voici dans le projet Groupie-Tracker, il est joué par groupe de 3 ou 4. De notre coté nous sommes composé de 4 membres (Mathéo.L, Nathy.M, Pierric.C et Enzo.P).
Nous avons pour mission de créer une page Web qui affiche 52 artistes avec toutes leurs données, caractéristiques, ainsi que l'emplacement des concerts de chacun sur une map.
Nous devons faire en sorte que l'utilisateur puisse choisir les artistes qu'il souhaite découvrir ou redécouvrir par l'intermédiaire d'une barre de recherche.
Laissons le suspense s'étendre à son maximum !
On se donne rendez-vous le jour de l'oral pour voir si le contrat est rempli ou non ;).

## :o: répartition du travail au sein du groupe

La répartition du travail a été plus compliqué que sur le précédent projet, nous avons eu des difficultés sur la communication en interne pour la répartition des rôles du à un manque de compréhension du sujet. Une fois cette phase difficle nous nous sommes finalement réparti les rôles de la manière suivante : Nathy a travaillé sur la map, Mathéo a s'est penché sur le search, Pierric lui a préféré s'occuper des filtres et Enzo a travaillé sur la page artistes  . Afin de travailler dans de bonnes conditions nous avons créé un trelo pour suivre les taches à faire et un environnement **gitea**.

*:small_red_triangle:Explication de notre environnement GITEA :*

Notre Répertoire est séparé en 6 branches bien distinct. Il y-a quatre branches nommer : **mleger**, **Enzo**, **Nathy**, **PCOME** elles sont des espaces de travail individuel à chaque membre du groupe. Elle permet au développeur de travailler sur un axe du code (search, map, artistes...) sans créer de conflit avec les autres. Ensuite nous avons décidé de créer une branche **dev** afin de déposer chaque bout de code des membres, elle permet de faire une mise en commun et en parallèle l'assemblage du programme final. Pour finir la branche **master** permet de récupérer le programme final situer dans la branche **dev** après un accord avec les développeur. Cependant seul le propriétaire du répertoire peut "push" sur la **master** afin d' éviter et de bloquer des "push" fait par inadvertance par les autres membres. Cela permet de sécuriser le programme qui s'y trouve et de le garder fonctionnel.

## :o: Organisation du code 

Le projet contient 3 dossiers a la racine : src/handlers, static, templates et un readme. Le dossier templates comporte tout les fichiers **html**, le dossier static comporte tout les **css** et **js** et un dossier img et vidéo, pour finir le dossier src/handlers contient tout les fichiers **go**  . Cette organisation permet de se retrouver plus rapidement dans le projet et de retrouver n'importe quel fichier sans accros.

## :o: Element à améliorer :

On aurait pu améliorer pas mal de chose :
- **Premièrement**, L'utilisation de la map google a son maximum, comme par exemple utiliser les infoWindow pour le markeur.

- **Deuxièmement**, 
- **Troisièmement**,


## :o: Outils pour le projet :
- **Visual Studio Code**
- **Trello** (https://trello.com/b/QCa9CDPc/groupie-t)
- **gitea** (https://git.ytrack.learn.ynov.com/MLEGER/groupie-tracker)