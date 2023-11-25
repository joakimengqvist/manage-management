export type NoteAuthor = {
    id: string
    name: string
    email: string
}
export interface Notes {
	id: string
	author_id: string
	author_name: string
	author_email: string
	project?: string
	sub_project?: string
    income?: string
    expense?: string
    external_company?: string
	title: string
	note: string
    created_at: string
    updated_at: string
}

export interface ProjectNote {
	id: string
	author_id: string
	author_name: string
	author_email: string
	project: string
	title: string
	note: string
    created_at: string
    updated_at: string
}

export interface SubProjectNote {
	id: string
	author_id: string
	author_name: string
	author_email: string
	sub_project: string
	title: string
	note: string
    created_at: string
    updated_at: string
}

export interface ExpenseNote {
	id: string
	author_id: string
	author_name: string
	author_email: string
	expense: string
	title: string
	note: string
    created_at: string
    updated_at: string
}

export interface IncomeNote {
	id: string
	author_id: string
	author_name: string
	author_email: string
	income: string
	title: string
	note: string
    created_at: string
    updated_at: string
}

export interface ExternalCompanyNote {
	id: string
	author_id: string
	author_name: string
	author_email: string
	external_company: string
	title: string
	note: string
    created_at: string
    updated_at: string
}