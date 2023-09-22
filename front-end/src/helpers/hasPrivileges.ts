export const hasPrivilege = (privileges : Array<string>, privilege : string) => {
    return privileges.includes(privilege)
}