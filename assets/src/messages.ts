export class UsersList {
    Users: PublicUser[]

    constructor() {
        this.Users = []
    }
}

export class PublicUser {
    Username: string
    Ready: boolean
    IsHost: boolean
    Players: string[]

    constructor() {
        this.Username = ""
        this.Ready = false
        this.IsHost = false
        this.Players = []
    }
}

export class PublicWorld {
    Planets: Planet[]
    Players: Player[]
    CurrentPlayer: Player
    WorldSettings: WorldSettings

    constructor() {
        this.Planets = []
        this.Players = []
        this.CurrentPlayer = new Player()
        this.WorldSettings = new WorldSettings
    }
}

export class Planet {
    Location: Point
    Radius: number
    Mass: number
    ResourceID: number
    PlanetID: number

    constructor() {
        this.Location = new Point()
        this.Radius = 0
        this.Mass = 0
        this.ResourceID = 0
        this.PlanetID = 0
    }
}

export class Player {
    Name: string
    PlanetID: number
    Alpha: number
    ColorIdx: number

    constructor() {
        this.Name = ""
        this.PlanetID = 0
        this.Alpha = 0
        this.ColorIdx = 0
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

export class PollTurn {
    CurrentPlayer: string
    CurrentPlayerAlpha: number
    ShotPerformed: boolean

    constructor() {
        this.CurrentPlayer = ""
        this.CurrentPlayerAlpha = 0
        this.ShotPerformed = false
    }
}

export class Trajectory {
    ArrowStates: ArrowState[]
    CollidedWith: string

    constructor() {
        this.ArrowStates = []
        this.CollidedWith = ""
    }
}

export class ArrowState {
    Position: Point
    Orientation: number

    constructor() {
        this.Position = new Point()
        this.Orientation = 0
    }
}

export class GameStatus {
    Started: boolean
    WorldSettings: WorldSettings

    constructor() {
        this.Started = false
        this.WorldSettings = new WorldSettings()
    }
}

export class WorldSettings {
    ShootTimeout: number
    LoopedWorld: boolean

    constructor() {
        this.ShootTimeout = 0
        this.LoopedWorld = false
    }
}