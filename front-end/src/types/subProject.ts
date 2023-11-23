export interface SubProject {
    id: string;
    name: string;
    description: string;
    status: string;
    priority: number;
    start_date: string;
    due_date: string;
    estimated_duration: number;
    notes: Array<string>;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
    projects: Array<string>;
    invoices: Array<string>;
    incomes: Array<string>;
    expenses: Array<string>;
}