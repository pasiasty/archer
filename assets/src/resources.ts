import * as ex from "excalibur";

const Images: { [key: string]: ex.Texture } = {
    sky: new ex.Texture('/public/assets/images/sky.jpg'),
    player: new ex.Texture('/public/assets/images/player.png'),
    arrow: new ex.Texture('/public/assets/images/arrow.png'),
};

const PlanetImages: ex.Texture[] = [
    new ex.Texture('/public/assets/images/planet0.png'),
    new ex.Texture('/public/assets/images/planet1.png'),
    new ex.Texture('/public/assets/images/planet2.png'),
    new ex.Texture('/public/assets/images/planet3.png'),
    new ex.Texture('/public/assets/images/planet4.png'),
    new ex.Texture('/public/assets/images/planet5.png'),
    new ex.Texture('/public/assets/images/planet6.png'),
    new ex.Texture('/public/assets/images/planet7.png'),
];

export const Colors: ex.Color[] = [
    ex.Color.White,
    ex.Color.Blue,
    ex.Color.Green,
    ex.Color.Black,
    ex.Color.Violet,
    ex.Color.Red,
    ex.Color.Yellow,
    ex.Color.Viridian,
    ex.Color.Orange,
    ex.Color.Magenta,
    ex.Color.Gray,
    ex.Color.Cyan,
    ex.Color.Azure,
    ex.Color.Chartreuse,
    ex.Color.Vermillion,
    ex.Color.White,
]

const loader = new ex.Loader();

for (const img in Images) {
    loader.addResource(Images[img]);
}

for (const planet of PlanetImages) {
    loader.addResource(planet)
}

export function getPlanetTexture(id: number): ex.Texture {
    return PlanetImages[id % PlanetImages.length]
}

export { Images, loader };