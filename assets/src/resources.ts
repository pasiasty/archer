import * as ex from "excalibur";

const Images: { [key: string]: ex.Texture } = {
    sky: new ex.Texture('/public/assets/images/sky.jpg'),
};

const PlanetImages: ex.Texture[] = [
    new ex.Texture('/public/assets/images/planet.png'),
    new ex.Texture('/public/assets/images/planet1.png'),
    new ex.Texture('/public/assets/images/planet2.png'),
    new ex.Texture('/public/assets/images/planet3.png'),
    new ex.Texture('/public/assets/images/planet4.png'),
    new ex.Texture('/public/assets/images/planet5.png'),
    new ex.Texture('/public/assets/images/planet6.png'),
    new ex.Texture('/public/assets/images/planet7.png'),
];

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