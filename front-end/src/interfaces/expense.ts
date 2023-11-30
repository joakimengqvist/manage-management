/* eslint-disable @typescript-eslint/no-explicit-any */
export interface ExpenseObject {
	id: string,
	project_id: string
    expense_date: any
	expense_category: string
	vendor: string
	description: string
	amount: number
	tax: number
	currency: string
	payment_method: string
	status: string
	created_by: string
	created_at: any
	updated_by: string
	updated_at: any
}
