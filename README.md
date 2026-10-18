# go-cloud-api

Nous allons construire dans ce lab un pipeline dont 
l'objectif est qu'à chaque fois que nous pousses du code sur GitHub, le pipeline puisse :
1. Tester ton application Go.
2. Construire une image Docker avec Buildx Bake.
3. Analyser les vulnérabilités avec Trivy.
4. Publier l'image sur GitHub Container Registry (GHCR), uniquement si les contrôles réussissent.
Nous allons également distinguer le rôle du Makefile, du Dockerfile, de docker-bake.hcl et du fichier GitHub Actions. C'est leur collaboration qui est importante.

1. Architecture du projet
go-cloud-api/

main.go — code de l'application
go.mod — module et dépendances Go
Dockerfile — construction de l'image
docker-bake.hcl — configuration Buildx Bake
Makefile — raccourcis pour les tâches courantes
.github/workflows/ci.yml — pipeline CI/CD
.trivyignore — exceptions documentées, si nécessaire


Nous utiliserons un seul dépôt et une image nommée ghcr.io/arsenedev/go-cloud-api. 
