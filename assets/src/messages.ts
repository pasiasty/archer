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
}
