import * as ex from "excalibur";

const Images: { [key: string]: ex.Texture } = {
    sky: new ex.Texture('/public/assets/images/sky.jpg'),
    earth: new ex.Texture('/public/assets/images/planet.png'),
};

const loader = new ex.Loader();

for (const img in Images) {
    loader.addResource(Images[img]);
}

export { Images, loader };