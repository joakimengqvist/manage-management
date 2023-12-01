export interface NoteAuthor {
    id: string
    name: string
    email: string
}
export interface Notes {
	id: string
	author_id: string
	author_name: string
	author_email: string
	project_id?: string
	sub_project_id?: string
    income_id?: string
    expense_id?: string
    external_company_id?: string
	product_id?: string
	invoice_id?: string
	invoice_item_id?: string
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
	project_id: string
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
	sub_project_id: string
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
	expense_id: string
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
	income_id: string
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
	external_company_id: string
	title: string
	note: string
    created_at: string
    updated_at: string
}

export interface ProductNote {
	id: string
	author_id: string
	author_name: string
	author_email: string
	product_id: string
	title: string
	note: string
    created_at: string
    updated_at: string
}

export interface InvoiceNote {
	id: string
	author_id: string
	author_name: string
	author_email: string
	invoice_id: string
	title: string
	note: string
    created_at: string
    updated_at: string
}

export interface InvoiceItemNote {
	id: string
	author_id: string
	author_name: string
	author_email: string
	invoice_item_id: string
	title: string
	note: string
    created_at: string
    updated_at: string
}