export type Notes = {
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

export type ProjectNote = {
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

export type SubProjectNote = {
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

export type ExpenseNote = {
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

export type IncomeNote = {
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

export type ExternalCompanyNote = {
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