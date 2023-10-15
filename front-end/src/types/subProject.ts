import { ProjectStatusTypes } from "../components/tags/ProjectStatus";

export interface SubProject {
    id: string;
    name: string;
    description: string;
    status: ProjectStatusTypes;
    priority: number;
    start_date: Date;
    due_date: Date;
    estimated_duration: number;
    notes: Array<string>;
    project_id: string;
    created_at: Date;
    created_by: string;
    updated_at: Date;
    updated_by: string;
    invoices: Array<string>;
    incomes: Array<string>;
    expenses: Array<string>;
}