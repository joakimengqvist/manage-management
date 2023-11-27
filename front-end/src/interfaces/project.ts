/* eslint-disable @typescript-eslint/no-explicit-any */
export interface Project {
    id: string
    name: string
    status: string
    sub_projects: Array<string>
    notes: Array<string>
    created_at: string
    created_by: string
    updated_at: string
    updated_by: string
    delete_project: any
}