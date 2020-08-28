export class UsersList {
    Users: PublicUser[]

    constructor() {
        this.Users = []
    }
}

export class PublicUser {
    Username: string | undefined
    Ready: boolean | undefined
    IsHost: boolean | undefined
    Players: string[]

    constructor() {
        this.Players = []
    }
}

export class PublicWorld {
    Planets: Planet[]

    constructor() {
        this.Planets = []
    }
}

export class Planet {
    Location: Point
    Radius: number
    Mass: number

    constructor() {
        this.Location = new Point()
        this.Radius = 0
        this.Mass = 0
    }
}

export class Point {
    X: number
    Y: number

    constructor() {
        this.X = 0
        this.Y = 0
    }
}