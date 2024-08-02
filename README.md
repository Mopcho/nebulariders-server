# nebulariders-server

## Project Workflow

- In order to merge your branch, use REBASE as it keeps the history cleaner

```
git checkout feature
git rebase main
```

Run the project:

```
cd server
go run main.go player.go
```

## PlayerActor { x, y, health, basic_attack }

- The initial x and y will be randomized for a random spawn point on the map
- On init the player model will query the database for his items
- Initial health will be calculated by his items
- Basic attack will be calculated by his items

## Items

- Items will be saved in a folder called /items, the server will load them in memmory when starting

## Authentitaction and Authrorization

- Token based authentiactions, we will attach the token to the query params of the socket /ws?token={token}

Login: /api/auth/login { username, password }
Register: /api/auth/register { username, password, email }

- For now use memory to store the users
- Use PCKE

## Client to Server Messages:

- AttackMeesage { type: "attack", attackType, enemyToAttack }
- PositionMessage { type: "position", x, y }
- JoinGameMessage { type: "join_game" }

## Server to client Messages:

- UpdatePositionMessage { type: "update_position", x, y } - Validation Failure
- DeathMessage { type: "death", by }
- EntityKilledMessage { type: "entity_killed", rewards }
- EntityDeathMessage { type: "entity_death", entity }
- PumpDataMessage { type: "pump", entities: id,pos[] }
- WorldStateMessage { type: "world_state_message", entities: Entity[] }
- PlayerJoinedMessage { type: "player_joined", entity: Entity }
