/* eslint-disable @typescript-eslint/no-explicit-any */
export interface Income {
	id: string
	project_id: string
    income_date: any
	invoice_id: string
	income_category: string
	statistics_income: boolean
	vendor: string
	description: string
	amount: number
	tax: number
	currency: string
	status: string
	created_by: string
	created_at: any
	updated_by: any
	updated_at: any
}