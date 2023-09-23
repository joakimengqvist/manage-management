/* eslint-disable @typescript-eslint/no-explicit-any */

export type State = {
    config: {
        darkMode: boolean
    }
user: {
    id: string,
    authenticated: boolean,
    firstName: string,
    lastName: string,
    email: string,
    privileges: Array<string>,
}
application: {
    privileges: Array<any>
    projects: Array<any>
    users: Array<any>
}
}