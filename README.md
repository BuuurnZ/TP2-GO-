# TP2: Système de Notifications et Logging

## Description
Application en ligne de commande qui simule l'envoi de notifications via différents canaux et archive un historique de tous les envois réussis.

## Fonctionnalités
- **Notificateurs** : Email, SMS (avec validation), Push
- **Validation SMS** : Les numéros doivent commencer par "06"
- **Archivage** : Stockage en mémoire des notifications réussies
- **Gestion d'erreurs** : Affichage des erreurs sans arrêter le programme

## Structure du code
- `main.go` : Point d'entrée et logique principale
- `notifiers.go` : Implémentation des notificateurs
- `storer.go` : Système d'archivage en mémoire

## Exécution
```bash
go run .
```

## Résultat attendu
Le programme affiche :
1. Les notifications envoyées avec succès
2. Les erreurs d'envoi
3. L'historique complet des notifications archivées
