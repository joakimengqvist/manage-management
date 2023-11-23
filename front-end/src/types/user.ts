export interface User {
    id: string;
    email: string;
    first_name: string;
    last_name: string;
    created_at: string;
    updated_at: string
    privileges: Array<string>
    projects: Array<string>
}