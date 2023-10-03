/* eslint-disable @typescript-eslint/no-explicit-any */
export type ExpenseObject = {
	expense_id: string,
	project_id: string,
    expense_date: any,
	expense_category: string,
	vendor: string,
	description: string,
	amount: number,
	tax: number,
	currency: string,
	payment_method: string,
	created_by: string,
	created_at: any,
	modified_by: any,
	modified_at: any
}