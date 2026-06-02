package ETEngine

import (
	"log"

	assets "github.com/Try-si/ETM/Assets"
	ETEStruct "github.com/Try-si/ETM/ETEStruct"
	ETMhelper "github.com/Try-si/ETM/Helper"
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

var (
	Gam    *ETEStruct.Game // game instance
	Assets assets.Assets   // assets instance
)

// if configPath is not .json, use configPath instead of config.json
func Init(update func(float64) error, configPath string) {
	Gam = &ETEStruct.Game{}

	if configPath[len(configPath)-5:] != ".json" {
		Gam.Conf = ETMhelper.StringToStruct[ETEStruct.Config](configPath)
	} else {
		Gam.Conf = ETMhelper.Jsontostruct[ETEStruct.Config](configPath) // initialisé la config
	}
	Gam.UpdateFunc = update // fonction de mise à jour

	Gam.MapConf = ETMhelper.Jsontostruct[ETEStruct.MapConfig](Gam.Conf.MapsPath)

	Assets = assets.Assets{}                // initialisé l'assets
	Assets.SpritePath = Gam.Conf.SpritePath // indiqué le chemin des sprites
	Assets.MapsPath = Gam.MapConf.ImgMap    // indiqué le chemin des maps
	Assets.Init()                           // initialisé les assets

	Gam.Assets = &Assets

	Gam.Maps = make(map[string]ETEStruct.Map) // initialisé les maps

	Maps := ETMhelper.GetAllFilesInDirectoryToStruct[ETEStruct.Map](Gam.MapConf.JsonMap) // récupère toutes les maps

	for _, mapData := range Maps {
		Gam.Maps[mapData.Name] = mapData.Obj                                          // ajoute la map au dictionnaire
		Gam.Maps[mapData.Name] = Gam.Maps[mapData.Name].MapInit(Gam.MapConf.Elements) // initialise la map
	}

	startMapData := Gam.Maps[Gam.Conf.Map]                    // récupère la map de départ
	Gam.SetScene(Gam.Conf.Map, Assets.Maps[startMapData.Map]) // set la scene de départ
}

func GameLoop() {
	ebiten.SetWindowSize(Gam.Conf.ScreenWidth, Gam.Conf.ScreenHeight) // set la taille de la fenêtre
	ebiten.SetWindowTitle(Gam.Conf.Title)                             // set le titre de la fenêtre
	if err := ebiten.RunGame(Gam); err != nil {                       // lance le jeu
		log.Fatal(err) // log l'erreur
	}
}
