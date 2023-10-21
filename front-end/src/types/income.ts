import { PaymentStatusTypes } from "../components/tags/ExpenseAndIncomeStatus"

/* eslint-disable @typescript-eslint/no-explicit-any */
export type IncomeObject = {
	id: string
	project_id: string
    income_date: any
	income_category: string
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