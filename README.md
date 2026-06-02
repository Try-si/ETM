# Description

ETM (Ebien Tiled Moteur) est un helper & framework de base pour les jeux utilisant le moteur Ebiten et les maps Tiled.

## Installation

```bash
go get github.com/Try-si/ETM
```

## Utilisation

architécture :

```
 Maps/
    Maps/ *
        overworld.tmx // pas obligatoire (c'est un exemple)
        vos map tiled (.tmx)
    Json/ *
        overworld.json // pas obligatoire (c'est un exemple)
        vos json de config de map
    Elements.json
    Maps.json
 Textures/
    vos textures
config.json
main.go

* = les nom des fichiers sont les mêmes
```

main.go :

```go
package main

import (
    "github.com/Try-si/ETM"
)

func main() {
    ETM.Init(func(deltaTime float64) error {
        // votre code ici
        return nil
    }, "config.json")
    
    ETM.GameLoop()
}
```

config.json :

```json
{
    "ScreenWidth": 800,          // largeur de l'écran
    "ScreenHeight": 600,         // hauteur de l'écran
    "Title": "Test",             // titre de la fenêtre
    "Map": "Overworld",          // map actuelle/de base

    "SpritePath": "Textures",    // chemin vers les sprites
    "MapsPath": "Maps/Maps.json" // chemin vers les maps
}
```

Maps.json : 

```json
{
    "Maps": ["Overworld"],
    "JsonMap": "Json",
    "ImgMap": "Maps",
    "Elements": "Elements.json"
}
```

Elements.json :

```json
{
    "Elements": {
        "Player": {
            "Image": "Player.png",
            "Size": [32, 32],
            "Rotation": 0,
            "Layer": 5,
            "Box": [0, 0, 0, 0] // witdh, height (si il est == a 0 alors c'est un cercle et witdh = rayon), box pos x, box pos y
        }
    }
}
```

Exemple map json (overworld.json) :

```json
{
    "Map": "Overworld",         // nom de la map
    "CellSize": 1,              // taille de la cellule
    "Unité": 32,                // taille d'une unité en pixels
    "Cam": {
        "Zoom": 1.0,
        "Offset": [0.0, 0.0]
    },

    "Elements": [
        {
            "Name": "Player",  // nom de l'élément dans Elements.json
            "Pos": [0.0, 0.0], // position de l'élément
            "MetaData": {
                "Nom de la variable": "valeur de la variable"
            }
        }
    ]       // éléments dans la map
}
```