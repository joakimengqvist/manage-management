import { ProjectStatusTypes } from "../components/tags/ProjectStatus";

/* eslint-disable @typescript-eslint/no-explicit-any */
export interface Project {
    id: string;
    name: string;
    status: ProjectStatusTypes;
    notes: Array<string>
    created_at: string
    updated_at: string
    delete_project: any
}