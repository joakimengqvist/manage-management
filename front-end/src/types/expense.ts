import { PaymentStatusTypes } from "../components/tags/ExpenseAndIncomeStatus"

/* eslint-disable @typescript-eslint/no-explicit-any */
export type ExpenseObject = {
	id: string,
	expense_id: string
	project_id: string
    expense_date: any
	expense_category: string
	vendor: string
	description: string
	amount: number
	tax: number
	currency: string
	payment_method: string
	status: PaymentStatusTypes
	created_by: string
	created_at: any
	modified_by: any
	modified_at: any
}